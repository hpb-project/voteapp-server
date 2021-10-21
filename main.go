package main

import (
	"flag"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/hpb-project/votedapp-server/common/logger"
	"github.com/hpb-project/votedapp-server/common/utils"
	"github.com/hpb-project/votedapp-server/config"
	"github.com/hpb-project/votedapp-server/db"
	"github.com/hpb-project/votedapp-server/router"
	"github.com/hpb-project/votedapp-server/task"
	"github.com/hpb-project/votedapp-server/version"
	log "github.com/sirupsen/logrus"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	defer func() {
		if r := recover(); r != nil {
			panicMes := fmt.Sprintf("panic: %s", r) //panic 信息
			log.WithFields(log.Fields{
				"panicMes": panicMes,
				"stack":    string(utils.PanicTrace()), //打印栈信息
			}).Error("panicStack")
		}
	}()

	initFlags()

	if err := config.Init(); err != nil {
		log.Fatalln("config init failed, err:", err)
	}
	conf := config.GetConfig()
	logger.Init(conf.LogDir, conf.LogLevel)
	err := db.Init()

	if err != nil {
		log.Fatalln("db init failed, err:", err)
	}

	err = task.Init()
	if err != nil {
		log.Fatalln("task init failed, err:", err)
	}

	err = task.Start()
	if err != nil {
		log.Fatalln("task start failed, err:", err)
	}

	go func() {
		r := gin.New()
		r.Use(Cors()) //默认跨域
		r.Use(logger.GinLoggerMiddleware(), gin.Recovery())
		router.InitRouter(r)

		log.Infof("Start HTTP Server on %s:%d", conf.IP, conf.Port)
		if err := r.Run(fmt.Sprintf("%s:%d", conf.IP, conf.Port)); err != nil {
			log.WithError(err).Fatalln("fail to run server")
		}
	}()

	exitSig := []os.Signal{os.Interrupt, syscall.SIGTERM}
	exitCh := make(chan os.Signal, 2)
	signal.Notify(exitCh, exitSig...)

	<-exitCh

	task.Stop()
	os.Exit(1)
}

var versionFlag bool

func initFlags() {
	flag.BoolVar(&versionFlag, "version", false, "print version info")
	flag.Parse()

	if versionFlag {
		version.ShowVersion()
	}
}

func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method
		origin := c.Request.Header.Get("Origin")
		if origin != "" {
			c.Header("Access-Control-Allow-Origin", "*")
			c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
			c.Header("Access-Control-Allow-Headers", "Content-Type,AccessToken,X-CSRF-Token, Authorization")
			c.Header("Access-Control-Allow-Credentials", "true")
			c.Set("content-type", "application/json")
		}
		//放行所有OPTIONS方法
		if method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
		}
		c.Next()
	}
}
