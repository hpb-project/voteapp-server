package main

import (
	"flag"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/hpb-project/votedapp-server/common/logger"
	"github.com/hpb-project/votedapp-server/common/utils"
	"github.com/hpb-project/votedapp-server/config"
	"github.com/hpb-project/votedapp-server/router"
	"github.com/hpb-project/votedapp-server/version"
	log "github.com/sirupsen/logrus"
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

	go func() {
		r := gin.New()
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
