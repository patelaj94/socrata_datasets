package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
)

const (
	DelawareDataDomain = "data.delaware.gov/resource/"
)

var DefaultRequest = Requester{
	Domain: DelawareDataDomain,
}

type Requester struct {
	Domain string
	client http.Client // Default Client
}

// Method will have a pointer receiver
// Method will take endpoint, list of query params, and return the response object
func (r *Requester) Request(endpoint string, params map[string]string, resp interface{}) error {
	var reqUrl = buildUrl(endpoint, params)
	// GET Request
	req, err := http.NewRequest("GET", reqUrl.String(), nil)
	req.Header.Set("Access-Control-Allow-Origin", "*")

	// Make Request
	response, err := r.client.Do(req) //fmt.Print(response)
	if err != nil {
		fmt.Printf("There was an error making the request", err)
		return err
	} else {
		err := json.NewDecoder(response.Body).Decode(&resp)
		if err != nil {
			fmt.Printf("Could not unmarshal response", err)
			return err
		}
		edJson, _ := json.Marshal(resp)
		ioutil.WriteFile("Output.json", edJson, os.ModePerm)

	}

	return nil
}

// Method will take query params and query host to build request url
func buildUrl(endpoint string, params map[string]string) url.URL {

	reqUrl := url.URL{}
	reqUrl.Scheme = "https"
	reqUrl.Path = DelawareDataDomain + endpoint
	q := reqUrl.Query()
	var isFirst = true
	for k, val := range params {
		if isFirst {
			q.Set(k, val)
			isFirst = false
		} else {
			q.Add(k, val)
		}
	}
	reqUrl.RawQuery = q.Encode()
	fmt.Println(reqUrl.String())

	return reqUrl
}
