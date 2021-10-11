package config

import (
	"encoding/json"
	"errors"
	"github.com/kelseyhightower/envconfig"
	log "github.com/sirupsen/logrus"
	"sync"
)

type Config struct {
	Common
	DB
}

type Common struct {
	IP        string `json:"ip" envconfig:"IP" default:"0.0.0.0"`
	Port      int    `json:"port" envconfig:"PORT" default:"8867"`
	LogLevel  string `json:"log_level" envconfig:"LOG_LEVEL" default:"info"`
	LogDir    string `json:"log_dir" envconfig:"LOG_DIR"`
	RPC       string `json:"rpc" envconfig:"RPC"`
	Mode      string `json:"mode" envconfig:"MODE"`
	ProxyAddr string `json:"proxy_addr" envconfig:"PROXY_ADDR"`
}

type DB struct {
	DBUserName           string `json:"db_user_name" envconfig:"DB_USER_NAME" default:"root"`
	DBUserPassword       string `json:"db_user_password" envconfig:"DB_USER_PASSWORD" required:"true" default:"Hpb!!0988"`
	DBIP                 string `json:"db_ip" envconfig:"DB_IP" required:"true" default:"127.0.0.1"`
	DBPort               int    `json:"db_port" envconfig:"DB_PORT" default:"3306"`
	DBName               string `json:"db_name" envconfig:"DB_NAME" default:"hpbvote"`
	DBMaxAllowedPacket   int    `json:"db_max_allowed_packet" envconfig:"DB_MAX_ALLOWEDPACKET" default:"0"`
	DBSQLMaxOpenConns    int    `json:"dbsql_max_open_conns" envconfig:"DB_SQL_MAX_OPENCONNS" default:"50"`
	DBSQLMaxIdleConns    int    `json:"dbsql_max_idle_conns" envconfig:"DB_SQL_MAX_IDLECONNS" default:"25"`
	DBSQLConnMaxLifetime int64  `json:"dbsql_conn_max_lifetime" envconfig:"DB_SQL_CONNMAXLIFETIME" default:"3600"`
}

var (
	once     sync.Once
	confLock sync.RWMutex
	gConf    = &Config{}
)

func Init() error {
	var err error
	once.Do(func() {
		err = parseEnv()
	})
	return err
}

func parseEnv() error {
	if err := envconfig.Process("hpb", gConf); err != nil {
		log.WithError(err).Error("load config from env failed")
		return err
	}

	if gConf.DBUserPassword == "" {
		return errors.New("empty DBUserPassword")
	}

	if gConf.DBIP == "" {
		return errors.New("empty DBIP")
	}

	return nil
}

// GetConfig get config
func GetConfig() *Config {
	confLock.RLock()
	defer confLock.RUnlock()

	return gConf
}

func (c *Config) String() string {
	cj, _ := json.Marshal(c)
	return string(cj)
}
