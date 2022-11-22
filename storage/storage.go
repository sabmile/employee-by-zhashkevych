package storage

import (
	"errors"
	"fmt"
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

type DumbStorage struct{}

func NewDumbStorage() *DumbStorage {
	return &DumbStorage{}
}

func (s *DumbStorage) Insert(e Employee) error {
	fmt.Printf("succesfully insert with such id: %d\n", e.Id)
	return nil
}

func (s *DumbStorage) Get(id int) (Employee, error) {
	e := Employee{
		Id: id,
	}
	return e, nil
}

func (s *DumbStorage) Delete(id int) error {
	fmt.Printf("succesfully delete with such id: %d", id)
	return nil
}
