package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/seyfBaskara/stripe-payment-session/initializers"
	"github.com/seyfBaskara/stripe-payment-session/services"
	"github.com/stripe/stripe-go/v74"
	"github.com/stripe/stripe-go/v74/checkout/session"
	"github.com/stripe/stripe-go/v74/price"
	"github.com/stripe/stripe-go/v74/product"

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

	router.POST("/create-checkout-session", createCheckoutSession)

	
	log.Fatal(server.Run(":" + config.ServerPort))
}

func createCheckoutSession(ctx *gin.Context) {
	domain := "http://localhost:8080"

	tempData := make(map[string]string)
	tempID := "temp_12345"

	priceID, err := createPrice()
	if err != nil {
		log.Printf("createPrice: %v", err)
	}

	tempData[tempID] = priceID


	params := &stripe.CheckoutSessionParams{
	  LineItems: []*stripe.CheckoutSessionLineItemParams{
		{
		  Price: stripe.String(tempData[tempID]),
		  Quantity: stripe.Int64(1),
		},
	  },
	  Mode: stripe.String(string(stripe.CheckoutSessionModePayment)),
	  SuccessURL: stripe.String(domain + "/api/success"),
	  CancelURL: stripe.String(domain + "/api/cancel"),
	}
  
	s, err := session.New(params)
  
	if err != nil {
	  log.Printf("session.New: %v", err)
	}
  
	ctx.Redirect(http.StatusSeeOther, s.URL)
  }

func createPrice() (string, error) {
	config, err := initializers.LoadConfig(".")
	if err != nil {
		log.Fatal("? Could not load environment variables", err)
	}
	
	stripe.Key = config.StripeSecretKey

	productParams := &stripe.ProductParams{
		Name: stripe.String("My Product"),
	}

	p, err := product.New(productParams)
	if err != nil {
		return "", err
	}

	params := &stripe.PriceParams{
		Product:    stripe.String(p.ID), 
		UnitAmount: stripe.Int64(1000),        
		Currency:   stripe.String("usd"),      
	}

	pr, err := price.New(params)
	if err != nil {
		return "", err
	}

	return pr.ID, nil
}
