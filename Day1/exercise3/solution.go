package main

import "fmt"

type totalSalary interface {
	Salary() int
}

type fullTime struct {
	numOfDays int
}

type contractor struct {
	numOfDays int
}

type freelancer struct {
	numOfDays int
	numOfHours int
}

func (fullTime fullTime) Salary() int{
	return 500*(fullTime.numOfDays)
}

func (contractor contractor) Salary() int{
	return 100*(contractor.numOfDays)
}

func (freelancer freelancer) Salary() int{
	return 10*(freelancer.numOfDays)*(freelancer.numOfHours)
}

func main(){
	var numOfDays, numOfHours int
	fmt.Println("Enter the number of days in the current month: ")
	fmt.Scan(&numOfDays)
	fmt.Println("Enter the number of hours the freelancer worked: ")
	fmt.Scan(&numOfHours)
	fullTime := fullTime{numOfDays: numOfDays}
	contractor := contractor{numOfDays: numOfDays}
	freelancer := freelancer{numOfDays: numOfDays, numOfHours: numOfHours}

	totalSalary := []totalSalary{fullTime, contractor, freelancer}
	fmt.Println("Salary of fullTime employee is: ",totalSalary[0].Salary())
	fmt.Println("Salary of contract employee is: ",totalSalary[1].Salary())
	fmt.Println("Salary of freelance employee is: ",totalSalary[2].Salary())


}