package handler

import (
	"github.com/gin-gonic/gin"
	gin_helper "github.com/louis296/gin-helper"
)

type Handler struct {
}

func handler(c *gin.Context, h interface{}) {
	gin_helper.HandleParserHandler(c, h, nil)
}
