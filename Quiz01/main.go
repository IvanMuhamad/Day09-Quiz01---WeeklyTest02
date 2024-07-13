package main

import (
	"Day09/Quiz01/models"
	"fmt"
	"time"
)

func main() {
	n := 100
	empCh := make(chan models.Employee, n)
	empCh2 := make(chan models.Employee, n)

	go models.CreateEmployees(n, empCh)
	go models.CreateEmployees(n, empCh2)

	var employees []models.Employee
	for emp := range empCh {
		employees = append(employees, emp)
	}

	var employees2 []models.Employee
	for emp := range empCh2 {
		employees2 = append(employees2, emp)
	}

	fmt.Println("Benchmark Test")
	fmt.Println("=======================================")

	start := time.Now()
	employees = models.TotalSalaries(employees)
	end := time.Since(start)
	fmt.Printf("Channels: %s\n", end)

	start = time.Now()
	employees2 = models.TotalSalariesNonChannels(employees2)
	end = time.Since(start)
	fmt.Printf("Tanpa Channels: %s\n", end)

	for _, emp := range employees {
		fmt.Printf("%+v\n", emp)
	}

	for _, emp := range employees2 {
		fmt.Printf("%+v\n", emp)
	}
}
