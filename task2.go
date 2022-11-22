package main

import "fmt"

type employee struct {
	name string

	salary int

	position string
}

type company struct {
	companyName string

	employees []employee
}

func main() {
	var employee1 employee
	var employee2 employee
	var employee3 employee

	employee1.name = "Talha"
	employee1.salary = 300000
	employee1.position = "Head Cyber Security"
	employee2.name = "Ahmed"
	employee2.salary = 33000
	employee2.position = "Team lead Management"
	employee3.name = "Aslam"
	employee3.salary = 79000
	employee3.position = "Team lead Finance"

	fmt.Println("*** First Employee Details ***")
	fmt.Println("Name: ", employee1.name)
	fmt.Println("Salary: ", employee1.salary)
	fmt.Println("Position: ", employee1.position)
	fmt.Println("*** Second Employee Details ***")
	fmt.Println("Name: ", employee2.name)
	fmt.Println("Salary: ", employee2.salary)
	fmt.Println("Position: ", employee2.position)
	fmt.Println("*** Third Employee Details ***")
	fmt.Println("Name: ", employee3.name)
	fmt.Println("Salary: ", employee3.salary)
	fmt.Println("Position: ", employee3.position)

	emplys := []employee{employee1, employee2, employee3} // Array of Employees
	for i := 0; i < len(emplys); i++ {
		fmt.Println(emplys[i])
	}

	var comp company
	comp.companyName = "Tetra"
	comp.employees = emplys
	fmt.Println("Company Name: ", comp.companyName)
	fmt.Println("Employee in Company: ", comp.employees)

}
