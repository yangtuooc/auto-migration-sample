package routes

import (
	"auto-migration-sample/application"
	"auto-migration-sample/common"
	"fmt"
	"github.com/gin-gonic/gin"
)

func Bootstrap() {
	router := gin.Default()
	setupRoute(router)
	if err := router.Run(common.Cfg.Server.UsePort()); err != nil {
		panic(fmt.Sprintf("start server err: %v", err))
	}
}

func setupRoute(router *gin.Engine) {
	apiGroup := router.Group(common.Cfg.Server.UseContext())
	apiGroup.GET("/ping", application.Ping)
	apiGroup.GET("/user", application.GetUser)
	apiGroup.POST("/user", application.RegisterUser)
}
