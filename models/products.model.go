package models

type Products struct {
	Items []ProductFields  `json:"items"`
}
type ProductFields struct {
	Fields ProductDetails
}
type ProductDetails struct {
	Price int    `json:"price"`
	Id int		`json:"id"`
}