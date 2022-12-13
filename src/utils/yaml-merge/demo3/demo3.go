package main

import (
	"fmt"
	"reflect"
)

type Addr struct {
	Province  string
	City      string
	Telephone string
}

type Student struct {
	Name    string
	Age     int
	Sex     string
	Address Addr
}

func testTypeOfAndValueOf() {
	fmt.Println("--------------start testTypeOfAndValueOf----------------")

	student := new(Student)
	student.Name = "python"
	fmt.Println(reflect.TypeOf(student))

	var student2 Student
	student2.Name = "golang"
	fmt.Println(reflect.TypeOf(student2))

	fmt.Println(reflect.ValueOf(student2).FieldByName("Name"))

	fmt.Println("--------------end testTypeOfAndValueOf----------------")
}

func main() {
	testTypeOfAndValueOf()
}
