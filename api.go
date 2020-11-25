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
	err   error
}

func main() {

	// TODO - Input for params will come from CLI input or GraphQL
	params1 := make(map[string]string)
	params1["race"] = "White"
	params1["schoolcode"] = "418"
	params1["schoolyear"] = "2020"

	// to experiment and learn go routines and channels
	wg := sync.WaitGroup{}
	ch := make(chan Result)

	wg.Add(2)
	go educatorAverageSalaryCall(params1, ch, &wg)
	go studentEnrollmentCall(params1, ch, &wg)

	go func() {
		wg.Wait()
		close(ch)
	}()

	for resp := range ch {
		if resp.err != nil {
			fmt.Printf("There was an error with your request", resp.err)
		} else {
			fmt.Println("You did it.")
			fmt.Println(resp.value)
		}
	}

}

func studentEnrollmentCall(params map[string]string, ch chan Result, wg *sync.WaitGroup) {
	defer wg.Done()
	var resp datastructs.StudentEnrollmentData
	res := new(Result)

	if err := endpoints.DefaultRequest.Request(StudentEnrollmentData, params, &resp); err != nil {
		res.err = err
		ch <- *res
	} else {
		res.value = resp
		ch <- *res
	}
}

func educatorAverageSalaryCall(params map[string]string, ch chan Result, wg *sync.WaitGroup) {
	defer wg.Done()
	var resp datastructs.EducatorAverageSalaryData
	res := new(Result)

	if err := endpoints.DefaultRequest.Request(EducatorAverageSalary, params, &resp); err != nil {
		res.err = err
		ch <- *res
	} else {
		res.value = resp
		ch <- *res
	}
}
