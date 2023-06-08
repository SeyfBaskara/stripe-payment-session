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
	Price int    `json:"price"`
	Id int		`json:"id"`
}

var (
	Client *http.Client
)

func NewProducts() Products {
	var p Products
	return p
}

func (p *Products) GetPrice (config initializers.Config){

	url := fmt.Sprintf("https://cdn.contentful.com/spaces/%s/environments/%s/entries?content_type=%s", config.ContentfulSpaceID, config.EnvironmentID, config.ContentTypes)

	err := GetJson(url, config.ContentfulAccesToken ,&p)
	if err != nil{
		fmt.Printf("error getting product: %v\n", err.Error())
		return
	}else {
		fmt.Println(p.Items)
	}

}

func GetJson(url string, token string, data interface{}) error {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return err
	}

	req.Header.Set("Authorization", "Bearer "+token)
	
	response, err := Client.Do(req)
	if err != nil {
		return err
	}

	defer response.Body.Close()

	return json.NewDecoder(response.Body).Decode(data)
}	
