package main

import (
	"datasets/datastructs"
	"datasets/endpoints"
	"fmt"
	"sync"
)

const (
	StudentEnrollmentData = "6i7v-xnmf.json"
	EducatorAverageSalary = "rv4m-vy79.json"
)

type Result struct {
	value interface{}
	error error
}

func main() {

	// TODO - Input for params will come from CLI input or GraphQL
	params1 := datastructs.EducatorAverageSalaryParams{
		Race:       "White",
		SchoolCode: 418,
		SchoolYear: "2020",
	}

	params2 := datastructs.StudentEnrollmentParams{
		Race:       "White",
		SchoolCode: 418,
		SchoolYear: "2020",
	}

	// to experiment and learn go routines and channels
	wg := sync.WaitGroup{}
	ch := make(chan Result)

	wg.Add(2)
	go EducatorAverageSalaryCall(params1, ch, &wg)
	go StudentEnrollmentCall(params2, ch, &wg)

	go func() {
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

	var resp datastructs.StudentEnrollmentData
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

	var resp datastructs.EducatorAverageSalaryData
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
