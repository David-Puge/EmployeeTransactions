package transactions

import (
	"fmt"
	"interface/pkg/display"
)

// RemoveEmployee Удаляет сотрудника из списка
func RemoveEmployee(emp *[]display.Employee, ID int) error {
	RemoveIndex := -1
	for i, Employee := range *emp {
		if Employee.ID == ID {
			RemoveIndex = i
			fmt.Printf("%s %s удален! \n", Employee.Name, Employee.LastName)
			break
		}
	}

	if RemoveIndex == -1 {
		return fmt.Errorf("Ошибка удаления сотрудника!(Неверный ID): \n ")
	}
	*emp = append((*emp)[:RemoveIndex], (*emp)[RemoveIndex+1:]...)

	return nil
}

// AddEmployee добавляет нового сотрудника в срез
func AddEmployee(e *[]display.Employee) error {
	var newEmployee display.Employee

	newEmployee.ID = 1
	if len(*e) > 0 {
		newEmployee.ID = (*e)[len(*e)-1].ID + 1
	}

	fmt.Println("Введите Имя сотрудника: ")
	_, err := fmt.Scan(&newEmployee.Name)
	if err != nil {
		return fmt.Errorf("Ошибка Ввода имени: %v", err)
	}

	fmt.Println("Введите Фамилию сотрудника: ")
	_, err = fmt.Scan(&newEmployee.LastName)
	if err != nil {
		return fmt.Errorf("Ошибка Ввода Фамилии: %v", err)
	}

	fmt.Println("Введите возраст сотрудника: ")
	_, err = fmt.Scan(&newEmployee.Age)
	if err != nil {
		return fmt.Errorf("Ошибка Ввода возраста: %v", err)
	}

	fmt.Println("Введите должность сотрудника: ")
	_, err = fmt.Scan(&newEmployee.Function)
	if err != nil {
		return fmt.Errorf("Ошибка Ввода должности: %v", err)
	}
	fmt.Println("Введите заработную плату сотрудника: ")
	_, err = fmt.Scan(&newEmployee.Salary)
	if err != nil {
		return fmt.Errorf("Ошибка Ввода заработной платы: %v", err)

	}

	*e = append(*e, newEmployee)
	return nil
}
