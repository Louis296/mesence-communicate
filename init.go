package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	gin_helper "github.com/louis296/gin-helper"
	"github.com/louis296/mesence-communicate/conf"
	"github.com/louis296/mesence-communicate/dao"
	"github.com/louis296/mesence-communicate/handler"
	"github.com/louis296/mesence-communicate/middleware"
	"github.com/louis296/mesence-communicate/pkg/jwt"
	"github.com/louis296/mesence-communicate/pkg/mongodb"
)

func Init(r *gin.Engine) {

	// init websocket handler
	r.GET("/ws", handler.WebSocketHandler)

	// init router
	r.Use(middleware.Cors())
	r.POST("/login", handler.UserLogin)
	r.Use(middleware.JWT())
	r.Any("/v1", gin_helper.MainHandler(nil, handler.Handler{}))

	// init configure
	configure, err := conf.GetConf()
	if err != nil {
		panic(err.Error())
	}

	// init database
	err = dao.InitDB(
		configure.Database.URL,
		configure.Database.DatabaseName,
		configure.Database.UserName,
		configure.Database.Password,
		configure.Database.MaxConn,
		configure.Database.MaxOpen,
	)
	if err != nil {
		panic(err.Error())
	}

	// set jwt secret
	jwt.Secret = configure.Jwt.Secret

	// init mongodb
	err = mongodb.InitClient(configure.MongoDB.Url)
	if err != nil {
		panic(err.Error())
	}

	err = r.Run(fmt.Sprintf(":%v", configure.Server.Port))
	if err != nil {
		panic("Start gin engine error")
	}
}
