package main

import (
	"log"
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/seyfBaskara/stripe-payment-session/initializers"

)

var (
	server 		*gin.Engine
)

func init (){
	_, err := initializers.LoadConfig(".")
	if err != nil {
		log.Fatal("? Could not load environment variables", err)
	}

	server = gin.Default()
}


func main () {
	config, err := initializers.LoadConfig(".")
	if err != nil {
		log.Fatal("? Could not load environment variables", err)
	}


	router := server.Group("api")
	router.GET("/ping", func(c *gin.Context) {
	  c.JSON(http.StatusOK, gin.H{
		"message": "ponggg",
	  })
	})
	
	log.Fatal(server.Run(":" + config.ServerPort))
}