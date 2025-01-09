package main

import (
    "fmt"
    "go-api-example/src/config"
    "go-api-example/src/routes"

    "github.com/gin-gonic/gin"
)

func main() {
    config, err := config.LoadEnvConfig()
    if err != nil {
        fmt.Println(err)
        return
    }

    engine := gin.Default()
    routes.SetupRoutes(engine)

    engine.SetTrustedProxies(nil)
    engine.Run(fmt.Sprintf(":%s", config["port"]))
}

func getEcho(ctx *gin.Context) {
    ctx.Data(200, "text/plain", []byte("Pong!"))
}
