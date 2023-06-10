package services


import (
	"fmt"
	"net/http"
	"encoding/json"
	"github.com/seyfBaskara/stripe-payment-session/initializers"
	"github.com/seyfBaskara/stripe-payment-session/models"

)

type ProductDetails struct {
	Products *models.Products
}

var (
	Client *http.Client
)

func NewProducts() ProductDetails {
	var pd ProductDetails
	return pd
}

func (pd *ProductDetails) GetPrice (config initializers.Config){

	url := fmt.Sprintf("https://cdn.contentful.com/spaces/%s/environments/%s/entries?content_type=%s", config.ContentfulSpaceID, config.EnvironmentID, config.ContentTypes)

	err := GetJson(url, config.ContentfulAccesToken ,&pd.Products)
	if err != nil{
		fmt.Printf("error getting product: %v\n", err.Error())
		return
	}else {
		fmt.Println(pd.Products.Items)
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
