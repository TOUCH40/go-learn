package main

import (
	"fmt"

	. "gopkg.in/ahmetb/go-linq.v3"
)

type Employee struct {
	Name     string
	Age      int
	Sex      int // 0 男 1 女
	WorkYear int // 工龄
}

func initEmployeeData() []Employee {
	list := make([]Employee, 0)
	for i := 0; i < 10; i++ {
		list = append(list, Employee{
			Name:     "张一",
			Age:      10,
			Sex:      i % 2,
			WorkYear: 1,
		})
	}
	return list
}

func main() {
	var manEmpRows []Employee
	rows := initEmployeeData()
	fmt.Println("===性别是男的所有员工列表去重===")
	From(rows).Distinct().ToSlice(&manEmpRows)
	fmt.Println(manEmpRows)
}
