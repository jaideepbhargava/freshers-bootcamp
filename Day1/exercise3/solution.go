
package main

import "fmt"

type finalSalary interface {
	Salary() int
}
type fullTime int
type freeLancer struct {
	numofDays int
	numofHours int
}
type contractor int

func (numOfDays fullTime) Salary() int{
	return 500 * int(numOfDays)
}

func (numOfDays contractor) Salary() int{
	return 100*int(numOfDays)
}

func (salary freeLancer) Salary() int{

	return salary.numofHours*salary.numofDays*10
}

func main(){
	fmt.Println("Enter the number of days in the current month")
	var numOfDays, numOfHours int
	fmt.Scan(&numOfDays)
	fmt.Println("Enter the number of hours the freelancer worked")
	fmt.Scan(&numOfHours)

	salaries := make([]finalSalary, 3)

	salaries[0] = fullTime(numOfDays)
	salaries[1] = contractor(numOfDays)
	salaries[2] = freeLancer{numofDays: numOfDays, numofHours: numOfHours}

	fmt.Println("Salary for full time employee is:", salaries[0].Salary())
	fmt.Println("Salary for full contract employee is:", salaries[1].Salary())
	fmt.Println("Salary for freelancer is:", salaries[2].Salary())

}