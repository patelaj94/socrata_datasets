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

	response, err := StudentEnrollmentCall("White", 418, "2020")
	if err != nil {
		fmt.Printf("There was an error with your request", err)
	} else {
		fmt.Printf("You did it.")
		fmt.Println(response)
	}


}


func StudentEnrollmentCall(race string, schoolCode int, schoolYear string) (dataStructs.StudentEnrollmentData, error) {

	params := dataStructs.StudentEnrollmentParams {
		Race: race,
		SchoolCode: schoolCode,
		SchoolYear: schoolYear,
	}

	var resp dataStructs.StudentEnrollmentData

	if err := endpoints.DefaultRequest.Request(StudentEnrollmentData, params, &resp); err != nil {
		emptyResponse := resp
		return emptyResponse, err
	} else {
		responseObject := resp
		return responseObject, nil
	}
}