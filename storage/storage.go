package storage

import (
	"errors"
	"sync"
)

type Employee struct {
	Id     int    `json:"id"`
	Name   string `json:"name"`
	Age    int    `json:"age"`
	Salary int    `json:"salary"`
	Sex    bool   `json:"sex"`
}

type Storage interface {
	Insert(e *Employee) error
	Get(id int) (Employee, error)
	Delete(id int) error
	Update(id int, e Employee)
	GetAll() ([]Employee, error)
}

type MemoryStorage struct {
	counter int
	data    map[int]Employee
	sync.Mutex
}

func NewMemoryStorage() *MemoryStorage {
	return &MemoryStorage{
		data:    make(map[int]Employee),
		counter: 1,
	}
}

func (s *MemoryStorage) Insert(e *Employee) error {
	s.Lock()
	defer s.Unlock()

	e.Id = s.counter
	s.data[e.Id] = *e

	s.counter++

	return nil
}

func (s *MemoryStorage) Get(id int) (Employee, error) {
	s.Lock()
	defer s.Unlock()

	e, exists := s.data[id]
	if !exists {
		return Employee{}, errors.New("employee with such id doesn't exist")
	}

	return e, nil
}

func (s *MemoryStorage) Delete(id int) error {
	s.Lock()
	defer s.Unlock()
	delete(s.data, id)
	return nil
}

func (s *MemoryStorage) Update(id int, e Employee) {
	s.Lock()
	defer s.Unlock()

	s.data[id] = e
}

func (s *MemoryStorage) GetAll() ([]Employee, error) {
	s.Lock()
	defer s.Unlock()

	employees := make([]Employee, 0)

	for _, v := range s.data {
		employees = append(employees, v)
	}

	return employees, nil
}
