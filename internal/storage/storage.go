package storage

import "context"

type Employee struct {
	Id        int
	Name      string
	ManagerId int
}

type Storage interface {
	CreateEmployee(ctx context.Context, emp Employee) error
	GetEmployees(ctx context.Context) ([]Employee, error)
}

type InMemoryStorage struct {
	data []Employee
}

func NewStorage() *InMemoryStorage {
	return &InMemoryStorage{[]Employee{}}
}

func (s *InMemoryStorage) CreateEmployee(ctx context.Context, emp Employee) error {
	s.data = append(s.data, emp)
	return nil
}

func (s *InMemoryStorage) GetEmployees(ctx context.Context) ([]Employee, error) {
	return s.data, nil
}
