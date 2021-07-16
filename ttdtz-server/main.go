package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
	"ttdtz-server/global"
	"ttdtz-server/internal/models"
	"ttdtz-server/internal/rmodels"
	"ttdtz-server/internal/routers"
	"ttdtz-server/pkg/logger"
	"ttdtz-server/pkg/setting"

	"github.com/gin-gonic/gin"
	lumberjack "gopkg.in/natefinch/lumberjack.v2"
)

func init() {
	err := setupSetting()
	if err != nil {
		log.Fatalf("init,setupSetting err: %v", err)
	}
	err = setupLogger()
	if err != nil {
		log.Fatalf("init.setupLogger err: %v", err)
	}
	err = setupDBEngine()
	if err != nil {
		log.Fatalf("init,setupDBEngine err: %v", err)
	}
}

// @title 突突大挑战
// @version 1.0
// @description Golang gin
// @termsOfService
func main() {
	gin.SetMode(global.GlobalConfig.Server.RunMode)
	router := routers.NewRouter()
	s := &http.Server{
		Addr:           ":" + global.GlobalConfig.Server.HttpPort,
		Handler:        router,
		ReadTimeout:    global.GlobalConfig.Server.ReadTimeout,
		WriteTimeout:   global.GlobalConfig.Server.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}

	global.Logger.Infof("%s: ttdtz-server/%s", "eddycjy", "ttdtz-server")

	s.ListenAndServe()

	//
	// r := gin.Default()
	// r.GET("/ping", func(c *gin.Context) {
	// 	c.JSON(200, gin.H{
	// 		"message": "pong",
	// 	})
	// })
	// r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

func setupSetting() error {
	setting, err := setting.NewSetting()
	if err != nil {
		return err
	}
	err = setting.ReadSection(&global.GlobalConfig)
	if err != nil {
		return err
	}
	//log.Fatalf("config = %+v, \n", global.GlobalConfig)
	global.GlobalConfig.Server.ReadTimeout *= time.Second
	global.GlobalConfig.Server.WriteTimeout *= time.Second
	return nil
}

func setupLogger() error {
	fileName := global.GlobalConfig.App.LogSavePath + "/" +
		global.GlobalConfig.App.LogFileName + global.GlobalConfig.App.LogFileExt
	fmt.Printf("log path = %+v", fileName)
	global.Logger = logger.NewLogger(&lumberjack.Logger{
		Filename:  fileName,
		MaxSize:   600,
		MaxAge:    30,
		LocalTime: true,
	}, "", log.LstdFlags).WithCaller(2)
	return nil
}

func setupDBEngine() error {
	var err error
	global.DBEngine, err = models.NewDBEngine(&global.GlobalConfig.Database)
	if err != nil {
		return err
	}
	global.CacheConnStrategy = rmodels.NewCache(&global.GlobalConfig.Redis)
	return nil
}
