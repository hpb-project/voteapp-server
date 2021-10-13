package db

import (
	"context"
	"errors"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/hpb-project/votedapp-server/common/logger"
	"github.com/hpb-project/votedapp-server/config"
	"github.com/jinzhu/gorm"
	log "github.com/sirupsen/logrus"
	"reflect"
	"strings"
	"sync"
	"time"
)

const (
	DBHealthCheckInterval       = 10 * time.Second
	DefaultDBAllowedPacket      = 16 * 1024 * 1024
	DBConnHealthCheckTimeout    = 10 * time.Second
	DBConnHealthCheckRetryTimes = 3
)

var (
	errDBConnHealthCheckTimeout = errors.New("db connection health check timeout")
	errOrmNotInited             = errors.New("orm  not inited")
)

var (
	omOnce sync.Once
	gOrm   *gorm.DB
)

// Init init db
func Init() error {
	var err error
	omOnce.Do(func() {
		err = initDB()
		if err != nil {
			return
		}

		go dbConnectionHealthCheck()
	})

	return err
}

func initDB() error {
	gConfig := config.GetConfig()
	dbSourceName := getDBSourceName(gConfig)
	log.WithField("dbSourceName", dbSourceName).Info("init db dbSourceName")

	db, err := gorm.Open("mysql", dbSourceName)
	if err != nil {
		log.WithError(err).Error("init db engine error")
		return err
	}

	db.DB().SetMaxOpenConns(gConfig.DBSQLMaxOpenConns)
	db.DB().SetMaxIdleConns(gConfig.DBSQLMaxIdleConns)
	db.DB().SetConnMaxLifetime(time.Duration(gConfig.DBSQLConnMaxLifetime) * time.Second)
	// 设置数据库表不为复数
	db.SingularTable(true)

	gOrm = db

	log.Info("init db connection success")
	return nil
}

// getBSourceName get dbSourceName
func getDBSourceName(config *config.Config) string {
	dbMaxAllowedPacket := config.DBMaxAllowedPacket
	if dbMaxAllowedPacket < DefaultDBAllowedPacket {
		dbMaxAllowedPacket = DefaultDBAllowedPacket
	}
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?maxAllowedPacket=%d&parseTime=true", config.DBUserName, config.DBUserPassword,
		config.DBIP, config.DBPort, config.DBName, dbMaxAllowedPacket)
}

// connHealthChecker check db connection health
func connHealthChecker() error {
	ctx, cancel := context.WithTimeout(context.TODO(), DBConnHealthCheckTimeout)
	defer cancel()

	if gOrm == nil {
		log.Error("orm is nil, maybe not inited")
		return errOrmNotInited
	}

	errChan := make(chan error)
	go func() {
		if err := gOrm.Exec("select 1").Error; err != nil {
			log.WithError(err).Error("db connection health check error")
		}
		errChan <- nil
	}()

	select {
	case err := <-errChan:
		if err != nil {
			return err
		}
	case <-ctx.Done():
		log.Error("db connection health check timeout")
		return errDBConnHealthCheckTimeout
	}
	return nil
}

// dbConnectionHealthCheck check db connection health with retry
func dbConnectionHealthCheck() {
	ticker := time.NewTicker(DBHealthCheckInterval)
	log.Debug("begin db health check")
	for range ticker.C {
		var err error
		for i := 0; i < DBConnHealthCheckRetryTimes; i++ {
			err = connHealthChecker()
			if err != nil {
				log.WithError(err).Error("db health check error ,begin retry")
				continue
			}
			break
		}
		if err != nil {
			log.WithError(err).Error("db health check error after retry")
		}
	}
}

func GetORM() *gorm.DB {
	if gOrm == nil {
		return nil
	}

	gConfig := config.GetConfig()
	// set db log mode
	if gConfig.LogLevel == logger.LogLevelDebug {
		gOrm.SetLogger(log.StandardLogger())
		return gOrm.Debug()
	}

	return gOrm
}

// GetBranchInsertSql 获取批量添加数据sql语句
func GetBatchInsertSql(objs []interface{}, tableName string) string {
	if len(objs) == 0 {
		return ""
	}
	fieldName := ""
	var valueTypeList []string
	fieldNum := reflect.TypeOf(objs[0]).NumField()
	fieldT := reflect.TypeOf(objs[0])
	for a := 0; a < fieldNum; a++ {
		name := GetColumnName(fieldT.Field(a).Tag.Get("gorm"))
		// 添加字段名
		if a == fieldNum-1 {
			fieldName += fmt.Sprintf("`%s`", name)
		} else {
			fieldName += fmt.Sprintf("`%s`,", name)
		}
		// 获取字段类型
		if fieldT.Field(a).Type.Name() == "string" {
			valueTypeList = append(valueTypeList, "string")
		} else if strings.Index(fieldT.Field(a).Type.Name(), "uint") != -1 {
			valueTypeList = append(valueTypeList, "uint")
		} else if strings.Index(fieldT.Field(a).Type.Name(), "int") != -1 {
			valueTypeList = append(valueTypeList, "int")
		}
	}
	var valueList []string
	for _, obj := range objs {
		objV := reflect.ValueOf(obj)
		v := "("
		for index, i := range valueTypeList {
			if index == fieldNum-1 {
				v += GetFormatField(objV, index, i, "")
			} else {
				v += GetFormatField(objV, index, i, ",")
			}
		}
		v += ")"
		valueList = append(valueList, v)
	}
	insertSql := fmt.Sprintf("insert into `%s` (%s) values %s", tableName, fieldName, strings.Join(valueList, ",")+";")
	return insertSql
}

// GetFormatField 获取字段类型值转为字符串
func GetFormatField(objV reflect.Value, index int, t string, sep string) string {
	v := ""
	if t == "string" {
		v += fmt.Sprintf("'%s'%s", objV.Field(index).String(), sep)
	} else if t == "uint" {
		v += fmt.Sprintf("%d%s", objV.Field(index).Uint(), sep)
	} else if t == "int" {
		v += fmt.Sprintf("%d%s", objV.Field(index).Int(), sep)
	}
	return v

}

// GetColumnName 获取字段名
func GetColumnName(jsonName string) string {
	for _, name := range strings.Split(jsonName, ";") {
		if strings.Index(name, "column") == -1 {
			continue
		}
		return strings.Replace(name, "column:", "", 1)
	}
	return ""
}

// BatchCreateModelsByPage 分页批量插入
func BatchCreateModelsByPage(tx *gorm.DB, dataList []interface{}, tableName string) (err error) {
	if len(dataList) == 0 {
		return
	}
	// 如果超过一百条, 则分批插入
	size := 100
	page := len(dataList) / size
	if len(dataList)%size != 0 {
		page += 1
	}
	for a := 1; a <= page; a++ {
		var bills = make([]interface{}, 0)
		if a == page {
			bills = dataList[(a-1)*size:]
		} else {
			bills = dataList[(a-1)*size : a*size]
		}
		sql := GetBatchInsertSql(bills, tableName)
		if err = tx.Exec(sql).Error; err != nil {
			fmt.Println(fmt.Sprintf("batch create data error: %v, sql: %s, tableName: %s", err, sql, tableName))
			return
		}
	}
	return
}
