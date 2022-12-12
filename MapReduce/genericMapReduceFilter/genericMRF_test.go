package genericMapReduceFilter

import (
	"fmt"
	"testing"
)

func sliceString() {

	list := []string{"1", "2", "3", "4", "5", "6"}
	result := Transform(list, func(a string) string {
		return a + a + a
	})

	fmt.Println(result)

	//{"111","222","333","444","555","666"}
}

func sliceInt() {

	list := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	r := TransformInPlace(list, func(a int) int {
		return a * 3
	})
	fmt.Println(r)
	//{3, 6, 9, 12, 15, 18, 21, 24, 27}
}

func sliceStruct() {

	var list = []Employee{
		{"Hao", 44, 0, 8000},
		{"Bob", 34, 10, 5000},
		{"Alice", 23, 5, 9000},
		{"Jack", 26, 0, 4000},
		{"Tom", 48, 9, 7500},
	}

	result := TransformInPlace(list, func(e Employee) Employee {
		e.Salary += 1000
		e.Age += 1
		return e
	})

	fmt.Println(result)
}

func Test_transform(t *testing.T) {
	sliceString()
	sliceInt()
	sliceStruct()
}
