package main

import (
	"fmt"
	"datasets/dataStructs"
	"datasets/endpoints"
)

const (
	StudentEnrollmentData = "6i7v-xnmf.json"
	EducatorAverageSalary = "rv4m-vy79.json"
)
func main() {

	// TODO - Input for params will come from CLI input or GraphQL
	params := dataStructs.EducatorAverageSalaryParams {
		Race: "White",
		SchoolCode: 418,
		SchoolYear: "2020",
	}

	response, err := EducatorAverageSalaryCall(params)

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

func EducatorAverageSalaryCall(params interface{}) (dataStructs.EducatorAverageSalaryData, error) {

	var resp dataStructs.EducatorAverageSalaryData

	if err := endpoints.DefaultRequest.Request(EducatorAverageSalary, params, &resp); err != nil {
		emptyResponse := resp
		return emptyResponse, err
	} else {
		responseObject := resp
		return responseObject, nil
	}
}