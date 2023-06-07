package services


import (
	"fmt"
	"net/http"
	"encoding/json"
	"github.com/seyfBaskara/stripe-payment-session/initializers"
)

type Products struct {
	Items []ProductFields  `json:"items"`
}
type ProductFields struct {
	Fields ProductDetails
}
type ProductDetails struct {
	Price int    `json:"price,omitempty"`
}

var (
	Client *http.Client
)

func NewProducts() Products {
	var p Products
	return p
}

func (p *Products) GetPrice (config initializers.Config){
	url := fmt.Sprintf("https://cdn.contentful.com/spaces/%v/entries?access_token=%v", config.ContentfulSpaceID, config.ContentfulAccesToken)

	err := GetJson(url, &p)
	if err != nil{
		fmt.Printf("error getting product: %v\n", err.Error())
		return
	}else {
		fmt.Println(p.Items)
	}

}

func GetJson(url string, data interface{}) error {
	response, err := Client.Get(url)
	if err != nil {
		return err
	}

	defer response.Body.Close()

	return json.NewDecoder(response.Body).Decode(data)
}	
