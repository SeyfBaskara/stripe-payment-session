package main

import (
	"log"
	"net/http"
	"github.com/gin-gonic/gin"
)

var (
	server 		*gin.Engine
)

func init (){
	server = gin.Default()
}


func main () {

	router := server.Group("api")
	router.GET("/ping", func(c *gin.Context) {
	  c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	  })
	})
	
	log.Fatal(server.Run())
}