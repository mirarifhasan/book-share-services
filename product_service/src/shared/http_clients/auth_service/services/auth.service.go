package auth_service

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"product_service/configs"
)

type UserIdsDto struct {
	Ids string `json:"ids"`
}

func GetUsersByIds(ids []int) ([]byte, error) {
	url := configs.GetEnv().AuthServiceBaseUrl + "/api/v1/users"

	reqBody := &UserIdsDto{
		Ids: strings.Trim(strings.Join(strings.Fields(fmt.Sprint(ids)), ","), "[]"),
	}
	data, _ := json.Marshal(reqBody)
	request, err := http.NewRequest("POST", url, bytes.NewBuffer(data))
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
	return respData, nil
}
