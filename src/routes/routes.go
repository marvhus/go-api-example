package routes

import (
    "go-api-example/src/controllers"

    "github.com/gin-gonic/gin"
)

func SetupRoutes(engine *gin.Engine) {
    group := engine.Group("/api")

    group.GET("/ping", controllers.GetPing)

    setupUserRoutes(group)
}

func setupUserRoutes(router_group *gin.RouterGroup) {
    group := router_group.Group("/users")

    group.GET("", controllers.GetUsers)
    group.GET("/:id", controllers.GetUser)

    group.POST("", controllers.CreateUser)
}
