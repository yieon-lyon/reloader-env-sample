package main

import (
  "net/http"
  "os"
  "github.com/gin-gonic/gin"
)

func main() {
  r := gin.Default()
  r.GET("/ping", func(c *gin.Context) {
    c.JSON(http.StatusOK, gin.H{
      "message": "pong",
    })
  })
  r.GET("/", func(c *gin.Context) {
    env := os.Getenv("NODE_ENV")
    c.JSON(http.StatusOK, gin.H{
      "message": "test",
      "env": env,
    })
  })
  r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
