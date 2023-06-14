package models

type CheckoutItem struct {
	Id int 				  `json:"id"`
	ProductName string   `json:"productName"` 
	Quantity int64 		 `json:"quantity"`
}