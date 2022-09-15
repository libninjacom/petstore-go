package petstore

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"petstore/request/createpets"
	"petstore/request/listpets"
	"petstore/request/showpetbyid"
)

type Client struct {
	baseUrl string
}

func (client Client) ListPets(opts ...listpets.RequestOption) (*listpets.Response, error) {
	builder := &listpets.Request{}
	for _, o := range opts {
		o(builder)
	}
	url := fmt.Sprintf("%s/pets", client.baseUrl)
	body := bytes.NewBuffer(nil)
	req, err := http.NewRequest("GET", url, body)
	if err != nil {
		return nil, err
	}
	q := req.URL.Query()
	if builder.Limit != nil {
		q.Add("limit", fmt.Sprint(*builder.Limit))
	}
	req.URL.RawQuery = q.Encode()
	httpclient := http.DefaultClient
	res, err := httpclient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	var result listpets.Response
	err = json.NewDecoder(res.Body).Decode(&result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (client Client) CreatePets() (*createpets.Response, error) {
	url := fmt.Sprintf("%s/pets", client.baseUrl)
	body := bytes.NewBuffer(nil)
	req, err := http.NewRequest("POST", url, body)
	if err != nil {
		return nil, err
	}
	httpclient := http.DefaultClient
	res, err := httpclient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	var result createpets.Response
	return &result, nil
}

func (client Client) ShowPetById(petId int) (*showpetbyid.Response, error) {
	builder := &showpetbyid.Request{PetId: petId}
	url := fmt.Sprintf("%s/pets/%d", client.baseUrl, builder.PetId)
	body := bytes.NewBuffer(nil)
	req, err := http.NewRequest("GET", url, body)
	if err != nil {
		return nil, err
	}
	httpclient := http.DefaultClient
	res, err := httpclient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	var result showpetbyid.Response
	err = json.NewDecoder(res.Body).Decode(&result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func NewClientFromEnv() Client {
	baseUrl, exists := os.LookupEnv("PET_STORE_BASE_URL")
	if !exists {
		fmt.Fprintln(os.Stderr, "Environment variable PET_STORE_BASE_URL is not set.")
		os.Exit(1)
	}
	return Client{baseUrl: baseUrl}
}

func NewClient(url string) Client {
	return Client{baseUrl: url}
}
