package main

import (
    "fmt"
    "net/http"
    "os"
    "github.com/gin-gonic/gin"
)

func main() {
    port := os.Getenv("PORT")
    if port == "" {
        port = "8080"
    }

    router := gin.Default()
    router.GET("/", func(c *gin.Context) {
        c.String(http.StatusOK, "Hello, Dude!")
    })

    fmt.Printf("Server listening on port %s...\n", port)
    err := router.Run(":" + port)
    if err != nil {
        panic(err)
    }
}
