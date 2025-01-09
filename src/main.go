package main

import (
    "fmt"
    "io"
    "net/http"
    "time"
)

func MainHandler(writer http.ResponseWriter, request *http.Request) {
    io.WriteString(writer, time.Now().Format("2006-01-02 15:04:05"))
}

func main() {
    http.HandleFunc("/", MainHandler)

    fmt.Println("Listening on port 5050...")

    http.ListenAndServe(":5050", nil)
}
