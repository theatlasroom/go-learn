package main

import (
	"fmt"
	"reflect"
)

type TestStruct struct {
	a, b, C, D string
}

func main() {
	arr := make([]string, 0)
	str := "test"
	fmt.Println(reflect.TypeOf(arr))
	fmt.Println(reflect.TypeOf(str))

	t := TestStruct{"this", "is", "a", "test"}
	fmt.Println(reflect.TypeOf(t))

	// TypeOf returns the reflection Type that represents the dynamic type of i.
	// If i is a nil interface value, TypeOf returns nil.
	tType := reflect.TypeOf(t)
	fmt.Println(tType, "\n")

	// ValueOf returns a new Value initialized to the concrete value
	// stored in the interface i. ValueOf(nil) returns the zero Value.
	tVal := reflect.ValueOf(t)
	fmt.Println(tVal, "\n")
	//tValue := tPtr.Elem()

	for i := 0; i < tType.NumField(); i++ {
		// returns an object of type StructField
		// https://golang.org/pkg/reflect/#StructField
		k := tType.Field(i)
		v := tVal.Field(i)

		// print some stuff
		fmt.Println(k.Name)  // struct field at index i name
		fmt.Println(k.Type)  // struct field at index i type
		fmt.Println(v, "\n") // struct field at index i's value'
	}
}
