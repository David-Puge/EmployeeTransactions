package main

import (
	"fmt"
	"interface/pkg/display"
	"interface/pkg/transactions"
)

func main() {
	employees := make([]display.Employee, 0, 5)
	Alise := display.Employee{
		ID:       1,
		Name:     "Алиса",
		LastName: "Фролова",
		Age:      24,
		Function: "Разработчик",
		Salary:   100000,
	}
	Dima := display.Employee{
		ID:       2,
		Name:     "Дима",
		LastName: "Пентичев",
		Age:      21,
		Function: "Разработчик",
		Salary:   60000,
	}
	Grisha := display.Employee{
		ID:       3,
		Name:     "Grisha",
		LastName: "Lucang",
		Age:      28,
		Function: "Разработчик",
		Salary:   80000,
	}
	employees = append(employees, Alise, Dima, Grisha)

	fmt.Println("Добавление нового сотрудника: ")
	err := transactions.AddEmployee(&employees)
	if err != nil {
		fmt.Println("err:", err)
	}

	fmt.Println("Вывод всех сотрудников на консоль: ")
	display.DisplayAll(employees)

	fmt.Println("Удаление сотрудника: ")
	err = transactions.RemoveEmployee(&employees, 1)
	if err != nil {
		fmt.Print(err)
	}

	fmt.Println("Вывод всех сотрудников на консоль: ")
	display.DisplayAll(employees)

	fmt.Println("Вывод отсортированых по возрасту сотрудников: ")
	display.AgeSort(employees, 22)
	fmt.Println("Вывод отсортированых по зарплате сотрудников: ")
	display.SalarySort(employees, 70000)
}
