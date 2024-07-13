package models

import (
	"math/rand"
	"sync"
	"time"

	"github.com/goombaio/namegenerator"
)

type Employee struct {
	ID          int
	FullName    string
	Salary      float64
	Status      string
	Insurance   float64
	Overtime    float64
	Allowance   float64
	TotalSalary float64
}

func generateName() string {
	seed := time.Now().UTC().UnixNano()
	nameGenerator := namegenerator.NewNameGenerator(seed)
	return nameGenerator.Generate()
}

var (
	statuses = []string{"Permanent", "Contract", "Trainee"}
)

func generateEmployee(id int) Employee {
	status := statuses[rand.Intn(len(statuses))]
	name := generateName()
	salary := rand.Float64()*(15000-5000) + 5000

	var insurance, overtime, allowance float64
	switch status {
	case "Permanent":
		insurance = 500000.00
	case "Contract":
		overtime = 55000.00
	case "Trainee":
		allowance = 100000.00
	}

	return Employee{
		ID:          id,
		FullName:    name,
		Salary:      salary,
		Status:      status,
		Insurance:   insurance,
		Overtime:    overtime,
		Allowance:   allowance,
		TotalSalary: 0.0,
	}
}

func CreateEmployees(n int, ch chan<- Employee) {
	for i := 1; i <= n; i++ {
		ch <- generateEmployee(i)
	}
	close(ch)
}

func CalculateSalary(emp Employee, wg *sync.WaitGroup, ch chan<- Employee) {
	defer wg.Done()
	emp.TotalSalary = emp.Salary + emp.Insurance + emp.Overtime + emp.Allowance
	ch <- emp
}

func TotalSalariesNonChannels(employees []Employee) []Employee {
	for i := range employees {
		employees[i].TotalSalary = employees[i].Salary + employees[i].Insurance + employees[i].Overtime + employees[i].Allowance
	}
	return employees
}

func TotalSalaries(employees []Employee) []Employee {
	var wg sync.WaitGroup
	ch := make(chan Employee, len(employees))

	for _, emp := range employees {
		wg.Add(1)
		go CalculateSalary(emp, &wg, ch)
	}

	wg.Wait()
	close(ch)

	var result []Employee
	for emp := range ch {
		result = append(result, emp)
	}

	return result
}
