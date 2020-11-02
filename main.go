package main

import (
	"net/http"
	"fmt"
	"io/ioutil"
	"os"
	"encoding/json"
	"datasets/dataStructs"
	"reflect"
	"strings"
)

const (
	DelawareDataDomain = "https://data.delaware.gov/resource/"
)

var DefaultRequest = Requester {
	Domain: DelawareDataDomain,

}

type Requester struct {
	Domain string
	client http.Client	// Default Client
}

// TODO: Move the logic below to a Generic "api.go" class which can be used to make multiple types of requests from this class.
func main() {

	params := dataStructs.StudentEnrollmentParams {
		Race: "White",
		SchoolCode: 418,
		SchoolYear: "2020",
	}

	var resp dataStructs.StudentEnrollmentData

	// TODO: Extract endpoint to grab from dataStruct Class
	err := DefaultRequest.Request("6i7v-xnmf.json", params , &resp)
	if err != nil {
		fmt.Printf("There was an error with your request", err)
	} else {
		fmt.Printf("You did it.")
		fmt.Println(resp)
	}

}

// Method for Requester Struct
// Method will have a pointer receiver
// Method will take endpoint, list of query params, and return the response object
func (r *Requester) Request(endpoint string, params interface{}, resp interface{}) error {

	// Make Params
	urlParams := r.makeParams(params)
	// Build Query Endpoint
	var reqUrl = DelawareDataDomain + endpoint + urlParams
	fmt.Println(reqUrl)
	// GET Request
	req, err := http.NewRequest("GET", reqUrl, nil)
	req.Header.Set("Access-Control-Allow-Origin", "*")

	// Make Request
	response, err := r.client.Do(req)

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


// Method for Requester Struct
// Method will take params interface and build http query params
func (r *Requester) makeParams(params interface{}) string {

	var output strings.Builder
	// Using reflection - idk what it's doing
	key := reflect.TypeOf(params)
	value := reflect.ValueOf(params)

	num := key.NumField()
	var firstParam = true
	for i:=0;i<num;i++ {
		tmp := value.Field(i).Interface()
		if !(tmp == 0 || tmp == "") {
			if (firstParam) {
				output.WriteString("?")
				firstParam = false
			} else {
				output.WriteString("&")
			}
			output.WriteString(strings.ToLower(key.Field(i).Name))
			output.WriteString("=")
			val := fmt.Sprintf("%v", tmp)
			output.WriteString(val)
		}
	}
	return output.String()
}