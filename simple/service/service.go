package service

import (
	"fmt"

	"github.com/mohammadanang/di_go_sample/simple/repository"
)

type SimpleService interface {
	Find(...interface{}) (interface{}, error)
	FindAll(...interface{}) ([]interface{}, error)
	Insert(...interface{}) (interface{}, error)
	Edit(...interface{}) (interface{}, error)
	Remove(...interface{}) (interface{}, error)
}

type Service struct {
	repo repository.SimpleRepository
}

func NewService(r repository.SimpleRepository) *Service {
	return &Service{
		repo: r,
	}
}

func (s *Service) Find(...interface{}) (interface{}, error) {
	s.repo.Read()
	fmt.Println("find from service...")
	return nil, nil
}

func (s *Service) FindAll(...interface{}) ([]interface{}, error) {

	fmt.Println("find all from service...")
	return nil, nil
}

func (s *Service) Insert(...interface{}) (interface{}, error) {
	fmt.Println("insert from service...")
	return nil, nil
}

func (s *Service) Edit(...interface{}) (interface{}, error) {
	fmt.Println("edit from service...")
	return nil, nil
}

func (s *Service) Remove(...interface{}) (interface{}, error) {
	fmt.Println("remove from service...")
	return nil, nil
}
