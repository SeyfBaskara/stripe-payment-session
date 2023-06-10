package main

import (
	"log"
	"net/http"
	"time"
	
	"github.com/gin-gonic/gin"
	"github.com/seyfBaskara/stripe-payment-session/initializers"
	"github.com/seyfBaskara/stripe-payment-session/services"

)


var (
	server 		*gin.Engine

	product 	services.ProductDetails
)

func init (){
	_, err := initializers.LoadConfig(".")
	if err != nil {
		log.Fatal("? Could not load environment variables", err)
	}

	services.Client = &http.Client{Timeout:20 * time.Second}

	product = services.NewProducts()

	server = gin.Default()
}


func main () {
	config, err := initializers.LoadConfig(".")
	if err != nil {
		log.Fatal("? Could not load environment variables", err)
	}

	 product.GetPrice(config)



	// router := server.Group("api")
	// router.GET("/ping", func(c *gin.Context) {
	//   c.JSON(http.StatusOK, gin.H{
	// 	"message": "ponggg",
	//   })
	// })
	
	// log.Fatal(server.Run(":" + config.ServerPort))
}