package product_service

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"product_service/configs"
)

func GetAllProducts() ([]byte, error) {
	url := configs.GetEnv().ProductServiceBaseUrl + "/api/v1/products"

	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	request.Header.Set("content-type", "application/json")
	resp, err := http.DefaultClient.Do(request)
	//Handle Error
	if err != nil {
		return nil, fmt.Errorf("an error occurred %v", err)
	}
	defer resp.Body.Close()
	//Read the response body
	respData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(string(respData))
	return respData, nil
}
