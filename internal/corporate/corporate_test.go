package corparate

import (
	"context"
	"testing"

	"github.com/tebrizetayi/optiopay/internal/storage"
)

func TestClosestCommonManager(t *testing.T) {
	// Create some employees and a corporate directory

	storageService := storage.NewStorage()
	ctx := context.Background()
	storageService.CreateEmployee(ctx, storage.Employee{Id: 1, Name: "Claire", ManagerId: 0})
	storageService.CreateEmployee(ctx, storage.Employee{Id: 2, Name: "John", ManagerId: 1})
	storageService.CreateEmployee(ctx, storage.Employee{Id: 3, Name: "Mary", ManagerId: 1})
	storageService.CreateEmployee(ctx, storage.Employee{Id: 4, Name: "Alice", ManagerId: 2})
	storageService.CreateEmployee(ctx, storage.Employee{Id: 5, Name: "Bob", ManagerId: 2})
	storageService.CreateEmployee(ctx, storage.Employee{Id: 6, Name: "Charlie", ManagerId: 3})

	corporate := NewCorporate(ctx, storageService)

	// Test the closest common manager between two employees
	testCases := []struct {
		name     string
		empId1   int
		empId2   int
		expected Employee
		err      error
	}{
		{"Both employees are the same", 2, 2, Employee{Id: 1, Name: "Claire"}, nil},
		{"Employees are not in the same branch of the hierarchy", 4, 6, Employee{Id: 1, Name: "Claire"}, nil},
		{"Employees are in the same branch of the hierarchy", 4, 5, Employee{Id: 2, Name: "John"}, nil},
		{"Employees are not in the same branch of the hierarchy but share a common ancestor", 4, 4, Employee{Id: 1, Name: "John"}, nil},
		{"Employees are in the same branch of the hierarchy but one is a manager of the other", 4, 2, Employee{Id: 1, Name: "Claire"}, nil},
		{"Ceo", 1, 1, Employee{}, nil},
		{"One of the employee is ceo", 1, 8, Employee{}, nil},
		{"One of the employees is not employed", -1, 3, Employee{}, ErrEmployeeNotFound},
		{"Both employees are not employed", -9, -8, Employee{}, ErrEmployeeNotFound},
	}

	for _, tc := range testCases {
		actual, err := corporate.FindClosestCommonManager(tc.empId1, tc.empId2)
		if actual.Name != tc.expected.Name && err == tc.err {
			t.Errorf("Expected closest common manager of %d and %d to be %s, but got %s", tc.empId1, tc.empId2, tc.expected.Name, actual.Name)
		}
	}
}
