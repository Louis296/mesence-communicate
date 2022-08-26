package main

import (
	"github.com/gin-gonic/gin"
	gin_helper "github.com/louis296/gin-helper"
	"github.com/louis296/mesence-communicate/conf"
	"github.com/louis296/mesence-communicate/dao"
	"github.com/louis296/mesence-communicate/handler"
	"github.com/louis296/mesence-communicate/pkg/jwt"
)

func Init(r *gin.Engine) {

	// init router
	r.Any("/v1", gin_helper.MainHandler(nil, handler.Handler{}))

	// init websocket handler
	r.GET("/ws", handler.WebSocketHandler)

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

	err = r.Run(":8081")
	if err != nil {
		panic("Start gin engine error")
	}
}
