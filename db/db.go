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
