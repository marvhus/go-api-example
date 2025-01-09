package main

import (
    "go-api-example/src/routes"

    "github.com/gin-gonic/gin"
)

func main() {
    engine := gin.Default()
    routes.SetupRoutes(engine)

    engine.SetTrustedProxies(nil)
    engine.Run(":8080")
}

func getEcho(ctx *gin.Context) {
    ctx.Data(200, "text/plain", []byte("Pong!"))
}
