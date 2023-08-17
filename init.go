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
	"github.com/louis296/mesence-communicate/pkg/kafka"
	"github.com/louis296/mesence-communicate/pkg/mongodb"
	"github.com/louis296/mesence-communicate/pkg/redis_client"
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
		configure.MySQL.URL,
		configure.MySQL.DatabaseName,
		configure.MySQL.UserName,
		configure.MySQL.Password,
		configure.MySQL.MaxConn,
		configure.MySQL.MaxOpen,
	)
	if err != nil {
		panic(err.Error())
	}

	// set jwt secret
	jwt.Secret = configure.Jwt.Secret

	// init mongodb
	err = mongodb.InitClient(configure.MongoDB.Url, configure.MongoDB.Database)
	if err != nil {
		panic(err.Error())
	}

	// init kafka
	err = kafka.InitProducer(configure.Kafka.Url, configure.Kafka.Topic)
	if err != nil {
		panic(err)
	}

	// init redis
	redis_client.InitClient(configure.Redis.Url, configure.Redis.Password, configure.Redis.DB)

	err = r.Run(fmt.Sprintf(":%v", configure.Server.Port))
	if err != nil {
		panic("Start gin engine error")
	}
}
