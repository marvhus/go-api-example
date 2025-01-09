package controllers

import (
    "fmt"
    "net/http"

    "github.com/gin-gonic/gin"
)

type User struct {
    ID      string `json:"id"`
    Name    string `json:"name" binding:"required"`
    Email   string `json:"email" binding:"required"`
}

var users = []User{}

func GetUsers(ctx *gin.Context) {
    ctx.JSON(http.StatusOK, users)
}

func GetUser(ctx *gin.Context) {
    id := ctx.Param("id")

    for _, user := range users {
        if user.ID != id {
            continue
        }

        ctx.JSON(http.StatusOK, user)
        return
    }

    ctx.JSON(http.StatusNotFound, gin.H{
        "message": fmt.Sprint("Unable to find user with ID '", id, "'."),
    })
}

func CreateUser(ctx *gin.Context) {
    if ctx.ContentType() != "application/json" {
        ctx.JSON(http.StatusBadRequest, gin.H{
            "message": fmt.Sprint("Invalid content type.  Expected 'application/json', but got '", ctx.ContentType(), "'."),
        })
        return
    }

    var received_user *User
    if err := ctx.Bind(&received_user); err != nil {
        ctx.Error(err)
        return
    }

    new_user := User{
        ID: fmt.Sprint(len(users)),
        Name: received_user.Name,
        Email: received_user.Email,
    }
    users = append(users, new_user)

    ctx.JSON(http.StatusOK, gin.H{
        "message": "Success.",
        "user": gin.H{
            "id": new_user.ID,
        },
    })
}
