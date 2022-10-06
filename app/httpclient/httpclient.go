package httpclient

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/go-resty/resty/v2"
)

func Get(uri string, queries map[string]string, headers map[string]string) (interface{}, error) {

	req := resty.New().R()
	req.SetHeader("Content-Type", "application/json")

	if queries != nil {
		req.SetQueryParams(queries)
	}

	if headers != nil {
		req.SetHeaders(headers)
	}
	resp, err := req.Get(uri)

	if err != nil {
		fmt.Println("Request failed:", err)
		return nil, err
	}

	fmt.Println("GET", uri, resp.StatusCode())

	if resp.StatusCode() != 200 {
		return nil, errors.New("StatusCode Not 200")
	}

	var res interface{}
	json.Unmarshal(resp.Body(), &res)

	// fmt.Println("RES", string(resp.Body()))

	return res, nil
}

func Post(uri string, queries map[string]string, headers map[string]string, body interface{}) (interface{}, error) {

	req := resty.New().R()
	req.SetHeader("Content-Type", "application/json")

	if queries != nil {
		req.SetQueryParams(queries)
	}

	if headers != nil {
		req.SetHeaders(headers)
	}

	if body != nil {
		req.SetBody(body)
	}

	resp, err := req.Post(uri)

	if err != nil {
		fmt.Println("Request failed:", err)
		return nil, err
	}

	fmt.Println("POST", uri, resp.StatusCode())

	if resp.StatusCode() != 200 {
		return nil, errors.New("StatusCode Not 200")
	}

	var res interface{}
	json.Unmarshal(resp.Body(), &res)

	// fmt.Println("RES", string(resp.Body()))

	return res, nil
}
