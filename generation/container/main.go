package main

import (
	"fmt"
	"github.com/a11en4sec/program-pattern/generation/container/gen"
)

//go:generate ./gen.sh ./template/container.tmp.go gen uint32 container
func generateUint32Example() {
	var u uint32 = 42
	c := gen.NewUint32Container()
	c.Put(u)
	v := c.Get()
	fmt.Printf("generateExample: %d (%T)\n", v, v)
}

//go:generate ./gen.sh ./template/container.tmp.go gen string container
func generateStringExample() {
	var s string = "Hello"
	c := gen.NewStringContainer()
	c.Put(s)
	v := c.Get()
	fmt.Printf("generateExample: %s (%T)\n", v, v)
}

type Employee struct {
	Name     string
	Age      int
	Vacation int
	Salary   int
}

//go:generate ./gen.sh ./template/filter.tmp.go gen Employee filter
func filterEmployeeExample() {

	var list = EmployeeList{
		{"Hao", 44, 0, 8000},
		{"Bob", 34, 10, 5000},
		{"Alice", 23, 5, 9000},
		{"Jack", 26, 0, 4000},
		{"Tom", 48, 9, 7500},
	}

	var filter EmployeeList
	filter = list.Filter(func(e *Employee) bool {
		return e.Age > 40
	})

	fmt.Println("----- Employee.Age > 40 ------")
	for _, e := range filter {
		fmt.Println(e)
	}

	filter = list.Filter(func(e *Employee) bool {
		return e.Salary <= 5000
	})

	fmt.Println("----- Employee.Salary <= 5000 ------")
	for _, e := range filter {
		fmt.Println(e)
	}
}

func main() {
	generateUint32Example()
	generateStringExample()

	filterEmployeeExample()
}
