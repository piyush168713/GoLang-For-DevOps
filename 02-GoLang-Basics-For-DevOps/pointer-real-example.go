package main
import "fmt"

type Employee struct{
	Name string
	Position string
	Salary int64
}

func updateSalary(emp *Employee, newSalary int64){
	fmt.Println(emp);
	emp.Salary = newSalary;
}

func main(){
	// create a Employee Instance
	emp := Employee{
		Name: "John Doe",
		Position: "Software Developer",
		Salary: 50000,
	}

	fmt.Println("Before Update:", emp)

	updateSalary(&emp, 60000)

	fmt.Println("After Update:", emp)
}