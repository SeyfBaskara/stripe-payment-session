package services

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/seyfBaskara/stripe-payment-session/initializers"
	"github.com/seyfBaskara/stripe-payment-session/models"
	"github.com/stripe/stripe-go/v74"
	"github.com/stripe/stripe-go/v74/checkout/session"
)

var (
	productD 	ProductDetails
)


func CreateSessionTest (ctx *gin.Context){
	config, err := initializers.LoadConfig(".")
	if err != nil {
		log.Fatal("? Could not load environment variables", err)
	}

	Client = &http.Client{Timeout:20 * time.Second}

	productD = NewProducts()

	prices, err := productD.GetPrice(config)
	if err != nil {
		log.Fatal("Could not get prices", err)
	}

	fmt.Println(prices)
	priceMap := make(map[int]int)
	for _, price := range prices {
		priceMap[price.Fields.Id] = price.Fields.Price
	}

	var payload []*models.CheckoutItem

	if err := ctx.ShouldBindJSON(&payload); err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	var items []models.CheckoutItem
	for _, item := range payload {
		if price, ok := priceMap[item.Id]; ok {
			newItem := models.CheckoutItem{
				Id:          item.Id,
				ProductName: item.ProductName,
				Quantity:    item.Quantity,
				Price: price,
			}
			items = append(items, newItem)
		}
	}


	ctx.JSON(http.StatusCreated, gin.H{"status": "success", "data": items})

}


func CreateCheckoutSession(ctx *gin.Context) {
	domain := "http://localhost:8080"

	var payload []*models.CheckoutItem

	if err := ctx.ShouldBindJSON(&payload); err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}


	var lineItems []*stripe.CheckoutSessionLineItemParams
	for _, item := range payload {
		lineItem := &stripe.CheckoutSessionLineItemParams{
			PriceData: &stripe.CheckoutSessionLineItemPriceDataParams{
				Currency: stripe.String("usd"),
				ProductData: &stripe.CheckoutSessionLineItemPriceDataProductDataParams{
					Name: stripe.String(item.ProductName),
				},
				UnitAmount: stripe.Int64(1000), // Replace with the actual price
			},
			Quantity: stripe.Int64(item.Quantity),
		}
		lineItems = append(lineItems, lineItem)
	}

	// Create a new Checkout Session
	params := &stripe.CheckoutSessionParams{
		LineItems:          lineItems,
		Mode: stripe.String(string(stripe.CheckoutSessionModePayment)),
		 SuccessURL: stripe.String(domain + "/api/success"),
	 	 CancelURL: stripe.String(domain + "/api/cancel"),
	}

	s, err := session.New(params)
  
	if err != nil {
	  log.Printf("session.New: %v", err)
	}
  
	ctx.JSON(http.StatusSeeOther, gin.H{"url": s.URL})
  }

  /*
  - get product price from prices list with id
  - create new payload with matched prices
  - append new payload array to items variable 
  */