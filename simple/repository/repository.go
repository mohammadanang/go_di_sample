package repository

import (
	"fmt"

	"github.com/mohammadanang/di_go_sample/database"
)

type SimpleRepository interface {
	Create(...interface{}) error
	Read(...interface{}) (interface{}, error)
	Update(...interface{}) (interface{}, error)
	Delete(...interface{}) error
}

type Repository struct {
	db database.SqlDb
}

func NewRepository(db database.SqlDb) *Repository {
	return &Repository{
		db: db,
	}
}

func (r *Repository) Create(...interface{}) error {
	fmt.Println("create from repository...")
	return nil
}

func (r *Repository) Read(...interface{}) (interface{}, error) {
	fmt.Println("read from repository...")
	return nil, nil
}

func (r *Repository) Update(...interface{}) (interface{}, error) {
	fmt.Println("update from repository...")
	return nil, nil
}

func (r *Repository) Delete(...interface{}) error {
	fmt.Println("delete from repository...")
	return nil
}
