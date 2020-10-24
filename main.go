package main

import (
	"net/http"
	"time"
	"fmt"
	"io/ioutil"
	"os"
	"encoding/json"
	"datasets/dataStructs"
)

func main() {

	// GET Request
	req, err := http.NewRequest("GET", "https://data.delaware.gov/resource/6i7v-xnmf.json?Race=White&" +
		"schoolcode=418&schoolyear=2020", nil)
	req.Header.Set("Access-Control-Allow-Origin", "*")
	client := &http.Client{Timeout: time.Second * 30}
	response, err := client.Do(req)

	// Log Response
	if err != nil {
		fmt.Printf("There was an error making the request", err)
	} else {
		var ed []dataStructs.EnrollmentData
		err := json.NewDecoder(response.Body).Decode(&ed)
		if err != nil {
			fmt.Printf("Could not unmarshal response", err)
		}
		edJson, _ := json.Marshal(ed)
		ioutil.WriteFile("Output.json", edJson, os.ModePerm)
		fmt.Println(ed)
		response.Body.Close()
	}
}

