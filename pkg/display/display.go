package display

import "fmt"

type Displayable interface {
	DisplayEmployee()
}

// Employee хранит данные о сотруднке
type Employee struct {
	ID       int
	Name     string
	LastName string
	Age      int
	Function string
	Salary   int
}

// Display выводит информацию о сотруднике типа Employee
func (e Employee) DisplayEmployee() {
	fmt.Printf("ID: %d, Имя: %s %s, Возраст: %d, Должность: %s, Зарплата: %d\n",
		e.ID, e.Name, e.LastName, e.Age, e.Function, e.Salary)
}

func Display(d Displayable) {
	d.DisplayEmployee()
}

// DisplayAll выводит на консоль всех сотрудников из среза
func DisplayAll(e []Employee) {
	for _, employee := range e {
		employee.DisplayEmployee()
	}
}

// AgeSort Выводит сотрудников старше указоного возраста
func AgeSort(e []Employee, age int) {

	for _, emp := range e {
		if emp.Age >= age {
			Display(emp)
		}
	}
}

// SalarySort Выводит сотрудников C зарплатой
func SalarySort(e []Employee, Salary int) {

	for _, emp := range e {
		if emp.Salary >= Salary {
			Display(emp)
		}
	}
}
