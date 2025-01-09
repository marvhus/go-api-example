package controllers

import (
    "github.com/gin-gonic/gin"
)

func GetPing(ctx *gin.Context) {
    ctx.Data(200, "text/plain", []byte("Pong!"))
}
