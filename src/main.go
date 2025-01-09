package main

import (
    "fmt"
    "os"

    "go-api-example/src/routes"

    "github.com/gin-gonic/gin"
    "github.com/spf13/viper"
)

func main() {
    loadConfig(os.Getenv("API_ENV"))

    engine := gin.Default()
    engine.SetTrustedProxies(nil)

    routes.SetupRoutes(engine)

    port := viper.GetString("api.port")
    if port == "" {
        fmt.Println("Unable to find port in config file.")
        return
    }
    engine.Run(fmt.Sprintf(":%s", port))
}

func getEcho(ctx *gin.Context) {
    ctx.Data(200, "text/plain", []byte("Pong!"))
}

func loadConfig(name string) {
    if name != "test" && name != "prod" {
        panic(fmt.Errorf("Invalid or missing API_ENV environment variable. Expected 'test' or 'prod'."))
    }
    fmt.Printf("Loading config file '%s'\n", name)

    viper.SetConfigName(name)
    viper.SetConfigType("yaml")
    viper.AddConfigPath("./config")
    err := viper.ReadInConfig()
    if err != nil {
        panic(fmt.Errorf("Fatal error config file: %w", err))
    }
}
