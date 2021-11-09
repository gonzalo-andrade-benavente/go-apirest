package main

import (
	"github.com/gin-gonic/gin"
)

func NewServer() *gin.Engine {
	s := gin.New()
	return s
}

func main() {
	srv := NewServer()
	srv.Run(":8080")
}
