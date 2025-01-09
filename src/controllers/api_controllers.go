package controllers

import (
    "net/http"

    "github.com/gin-gonic/gin"
)

func GetPing(ctx *gin.Context) {
    ctx.Data(http.StatusOK, "text/plain", []byte("Pong!"))
}
