package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/seyfBaskara/stripe-payment-session/initializers"
	"github.com/seyfBaskara/stripe-payment-session/services"
	"github.com/stripe/stripe-go/v74"

)


var (
	server 		*gin.Engine

	productD 	services.ProductDetails
)

func init (){
	_, err := initializers.LoadConfig(".")
	if err != nil {
		log.Fatal("? Could not load environment variables", err)
	}

	services.Client = &http.Client{Timeout:20 * time.Second}

	productD = services.NewProducts()

	server = gin.Default()
}


func main () {
	config, err := initializers.LoadConfig(".")
	if err != nil {
		log.Fatal("? Could not load environment variables", err)
	}

	server.LoadHTMLGlob("templates/*")
	stripe.Key = config.StripeSecretKey

	//  product.GetPrice(config)

	router := server.Group("api")
	router.GET("/healthcheck", func(ctx *gin.Context) {
	  ctx.JSON(http.StatusOK, gin.H{
		"message": "route works well!!",
	  })
	})

	router.GET("/checkout", func (ctx *gin.Context)  {
		ctx.HTML(http.StatusOK, "checkout.html", gin.H{
			"title": "checkout",
		})
	})
	router.GET("/success", func (ctx *gin.Context)  {
		ctx.HTML(http.StatusOK, "success.html", gin.H{
			"title": "success",
		})
	})
	router.GET("/cancel", func (ctx *gin.Context)  {
		ctx.HTML(http.StatusOK, "cancel.html", gin.H{
			"title": "cancel",
		})
	})

	router.POST("/create-checkout-session", services.CreateCheckoutSession)
	
	log.Fatal(server.Run(":" + config.ServerPort))
}

