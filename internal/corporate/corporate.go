package corparate

import (
	"context"
	"errors"

	"github.com/tebrizetayi/optiopay/internal/storage"
)

var ErrEmployeeNotFound = errors.New("Employee not found")

type Employee struct {
	Id            int
	Name          string
	DirectReports []*Employee
}

type Corparate struct {
	Directory *Employee
}

type Directory interface {
	FindClosestCommonManager(emp1, emp2 int) (Employee, error)
}

func NewCorporate(ctx context.Context, storage storage.Storage) Corparate {
	employees, _ := storage.GetEmployees(ctx)
	ceo := createDirectory(employees)
	return Corparate{ceo}
}

func createDirectory(employees []storage.Employee) *Employee {
	employeeMap := make(map[int]*Employee, 0)

	ceoId := 0
	for _, emp := range employees {
		if _, ok := employeeMap[emp.Id]; !ok {
			employeeMap[emp.Id] = &Employee{Id: emp.Id, Name: emp.Name, DirectReports: []*Employee{}}
		}
		if _, ok := employeeMap[emp.ManagerId]; !ok {
			employeeMap[emp.ManagerId] = &Employee{}
		}
		employeeMap[emp.ManagerId].DirectReports = append(employeeMap[emp.ManagerId].DirectReports, employeeMap[emp.Id])

		// Finding CEO. CEO is employee that has no manager.
		if emp.ManagerId == 0 {
			ceoId = emp.Id
		}
	}

	return employeeMap[ceoId]
}

// Returns the closest common manager for two employees
func (c Corparate) FindClosestCommonManager(empId1, empId2 int) (Employee, error) {
	// Get the path from the CEO to each employee
	path1 := c.getPathToEmployee(&Employee{Id: empId1})
	if path1 == nil {
		return Employee{}, ErrEmployeeNotFound
	}
	path2 := c.getPathToEmployee(&Employee{Id: empId2})
	if path2 == nil {
		return Employee{}, ErrEmployeeNotFound
	}

	// Traverse the paths until they diverge
	var i int
	for i < len(path1) && i < len(path2) && path1[i] == path2[i] {
		i++
	}

	// Return the closest common manager (i-1) from the CEO
	if i > 0 {

		return Employee{
			Id:   path1[i-1].Id,
			Name: path1[i-1].Name,
		}, nil
	}
	return Employee{}, nil
}

// Returns the path from the CEO to an employee
func (c Corparate) getPathToEmployee(employee *Employee) []*Employee {
	// Check if the given employee is the CEO
	if employee == c.Directory {
		return []*Employee{}
	}

	// Traverse the hierarchy to find the given employee and store the path of managers from the CEO
	var path []*Employee
	queue := []*Employee{c.Directory}
	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]
		for _, sub := range current.DirectReports {
			if sub.Id == employee.Id {
				// Found the employee, add all managers from the CEO to the path and return it
				path = append(path, current)
				return append(c.getPathToEmployee(current), path...)
			}
			queue = append(queue, sub)
		}
	}

	// Employee not found, return nil
	return nil
}
