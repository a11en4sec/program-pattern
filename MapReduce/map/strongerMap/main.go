package main

import (
	"fmt"
	"reflect"
)

type Employee struct {
	Name     string
	Age      int
	Vacation int
	Salary   int
}

func Transform(slice, function interface{}) interface{} {
	return transform(slice, function, false)
}

func TransformInPlace(slice, function interface{}) interface{} {
	return transform(slice, function, true)
}

func transform(slice, function interface{}, inPlace bool) interface{} {

	//check the `slice` type is Slice
	sliceInType := reflect.ValueOf(slice)
	if sliceInType.Kind() != reflect.Slice {
		panic("transform: not slice")
	}

	//check the function signature
	fn := reflect.ValueOf(function)
	elemType := sliceInType.Type().Elem()
	if !verifyFuncSignature(fn, elemType, nil) {
		panic("trasform: function must be of type func(" + sliceInType.Type().Elem().String() + ") outputElemType")
	}

	sliceOutType := sliceInType
	if !inPlace {
		sliceOutType = reflect.MakeSlice(reflect.SliceOf(fn.Type().Out(0)), sliceInType.Len(), sliceInType.Len())
	}
	for i := 0; i < sliceInType.Len(); i++ {
		sliceOutType.Index(i).Set(fn.Call([]reflect.Value{sliceInType.Index(i)})[0])
	}
	return sliceOutType.Interface()

}

func verifyFuncSignature(fn reflect.Value, types ...reflect.Type) bool {

	//Check it is a funciton
	if fn.Kind() != reflect.Func {
		return false
	}
	// NumIn() - returns a function type's input parameter count.
	// NumOut() - returns a function type's output parameter count.
	if (fn.Type().NumIn() != len(types)-1) || (fn.Type().NumOut() != 1) {
		return false
	}
	// In() - returns the type of a function type's i'th input parameter.
	for i := 0; i < len(types)-1; i++ {
		if fn.Type().In(i) != types[i] {
			return false
		}
	}
	// Out() - returns the type of a function type's i'th output parameter.
	outType := types[len(types)-1]
	if outType != nil && fn.Type().Out(0) != outType {
		return false
	}
	return true
}

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

func main() {
	sliceString()
	sliceInt()
	sliceStruct()

}
