package services

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	// "github.com/seyfBaskara/stripe-payment-session/initializers"
	"github.com/seyfBaskara/stripe-payment-session/models"
	"github.com/stripe/stripe-go/v74"
	"github.com/stripe/stripe-go/v74/checkout/session"
	// "github.com/stripe/stripe-go/v74/price"
	// "github.com/stripe/stripe-go/v74/product"

)


func CreateSessionTest (ctx *gin.Context){
	var payload []*models.CheckoutItem

	if err := ctx.ShouldBindJSON(&payload); err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	var items []models.CheckoutItem
	for _, item := range payload {
		newItem := models.CheckoutItem{
			Id:          item.Id,
			ProductName: item.ProductName,
			Quantity:    item.Quantity,
		}
		items = append(items, newItem)
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

	// var items []models.CheckoutItem
	// for _, item := range payload {
	// 	newItem := models.CheckoutItem{
	// 		Id:          item.Id,
	// 		ProductName: item.ProductName,
	// 		Quantity:    item.Quantity,
	// 	}
	// 	items = append(items, newItem)
	// }

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

	// priceID, err := CreatePrice(newItem)
	// if err != nil {
	// 	log.Printf("createPrice: %v", err)
	// }
	

	// params := &stripe.CheckoutSessionParams{
	//   LineItems: []*stripe.CheckoutSessionLineItemParams{
	// 	{
	// 		Price: stripe.String(priceID),
	// 		Quantity: stripe.Int64(1),
	// 	},
	//   },
	//   Mode: stripe.String(string(stripe.CheckoutSessionModePayment)),
	//   SuccessURL: stripe.String(domain + "/api/success"),
	//   CancelURL: stripe.String(domain + "/api/cancel"),
	// }
  
	s, err := session.New(params)
  
	if err != nil {
	  log.Printf("session.New: %v", err)
	}

	//NOTE might need to redirect from backend directly 
  
	ctx.JSON(http.StatusSeeOther, gin.H{"url": s.URL})
  }

// func CreatePrice(newItem models.CheckoutItem) (string, error) {
// 	config, err := initializers.LoadConfig(".")
// 	if err != nil {
// 		log.Fatal("? Could not load environment variables", err)
// 	}
	
// 	stripe.Key = config.StripeSecretKey

// 	productParams := &stripe.ProductParams{
// 		Name: stripe.String(newItem.ProductName),
// 	}

// 	p, err := product.New(productParams)
// 	if err != nil {
// 		return "", err
// 	}

// 	params := &stripe.PriceParams{
// 		Product:    stripe.String(p.ID), 
// 		UnitAmount: stripe.Int64(1000),        
// 		Currency:   stripe.String("usd"),      
// 	}

// 	pr, err := price.New(params)
// 	if err != nil {
// 		return "", err
// 	}

// 	return pr.ID, nil
// }
