package main

import (
	"fmt"
	"datasets/dataStructs"
	"datasets/endpoints"
)

const (
	StudentEnrollmentData = "6i7v-xnmf.json"
)
func main() {

	// TODO - Input for params will come from CLI input or GraphQL
	params := dataStructs.StudentEnrollmentParams {
		Race: "White",
		SchoolCode: 418,
		SchoolYear: "2020",
	}

	response, err := StudentEnrollmentCall(params)

	if err != nil {
		fmt.Printf("There was an error with your request", err)
	} else {
		fmt.Printf("You did it.")
		fmt.Println(response)
	}


}


func StudentEnrollmentCall(params interface{}) (dataStructs.StudentEnrollmentData, error) {

	var resp dataStructs.StudentEnrollmentData

	if err := endpoints.DefaultRequest.Request(StudentEnrollmentData, params, &resp); err != nil {
		emptyResponse := resp
		return emptyResponse, err
	} else {
		responseObject := resp
		return responseObject, nil
	}
}