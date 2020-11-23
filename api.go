package main

import (
	"fmt"
	"datasets/dataStructs"
	"datasets/endpoints"
	"sync"
)

const (
	StudentEnrollmentData = "6i7v-xnmf.json"
	EducatorAverageSalary = "rv4m-vy79.json"
)

type Result struct {
	value 	interface{}
	error  	error
}

func main() {

	// TODO - Input for params will come from CLI input or GraphQL
	params1 := dataStructs.EducatorAverageSalaryParams {
		Race: "White",
		SchoolCode: 418,
		SchoolYear: "2020",
	}

	params2 := dataStructs.StudentEnrollmentParams {
		Race: "White",
		SchoolCode: 418,
		SchoolYear: "2020",
	}

	// to experiment and learn go routines and channels
	wg := sync.WaitGroup{}
	ch := make(chan Result)

	wg.Add(2) // tentative
	go EducatorAverageSalaryCall(params1, ch, &wg)
	go StudentEnrollmentCall(params2, ch, &wg)


	go func(){
		wg.Wait()
		close(ch)
	}()

	for resp := range ch {
		if resp.error != nil {
			fmt.Printf("There was an error with your request", resp.error)
		} else {
			fmt.Println("You did it.")
			fmt.Println(resp.value)
		}
	}

}


func StudentEnrollmentCall(params interface{}, ch chan Result, wg *sync.WaitGroup) {

	var resp dataStructs.StudentEnrollmentData
	res := new(Result)

	if err := endpoints.DefaultRequest.Request(StudentEnrollmentData, params, &resp); err != nil {
		res.error = err
		ch <- *res
	} else {
		res.value = resp
		ch <- *res
	}
	wg.Done()
}

func EducatorAverageSalaryCall(params interface{}, ch chan Result, wg *sync.WaitGroup) {

	var resp dataStructs.EducatorAverageSalaryData
	res := new(Result)

	if err := endpoints.DefaultRequest.Request(EducatorAverageSalary, params, &resp); err != nil {
		res.error = err
		ch <- *res
	} else {
		res.value = resp
		ch <- *res
	}
	wg.Done()
}