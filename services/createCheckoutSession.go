package services

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/seyfBaskara/stripe-payment-session/initializers"
	"github.com/stripe/stripe-go/v74"
	"github.com/stripe/stripe-go/v74/checkout/session"
	"github.com/stripe/stripe-go/v74/price"
	"github.com/stripe/stripe-go/v74/product"

)


func CreateCheckoutSession(ctx *gin.Context) {
	domain := "http://localhost:8080"

	tempData := make(map[string]string)
	tempID := "temp_12345"

	priceID, err := CreatePrice()
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

func CreatePrice() (string, error) {
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
