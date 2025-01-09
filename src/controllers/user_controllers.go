package controllers

import (
    "fmt"

    "github.com/gin-gonic/gin"
)

type User struct {
    ID      int    `json:"id"`
    Name    string `json:"name" binding:"required"`
    Email   string `json:"email" binding:"required"`
}

var users []User = []User{}

func GetUsers(ctx *gin.Context) {
    ctx.JSON(200, users)
}

func CreateUser(ctx *gin.Context) {
    if ctx.ContentType() != "application/json" {
        ctx.JSON(400, gin.H{
            "message": fmt.Sprintf("Invalid content type.  Expected 'application/json', but got '%s'.", ctx.ContentType()),
        })
        return
    }

    var received_user *User
    if err := ctx.Bind(&received_user); err != nil {
        ctx.Error(err)
        return
    }

    new_user := User{
        ID: len(users),
        Name: received_user.Name,
        Email: received_user.Email,
    }
    users = append(users, new_user)

    ctx.JSON(200, gin.H{
        "message": "Success.",
    })
}
