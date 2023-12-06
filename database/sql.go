package database

import (
	"fmt"

	"github.com/mohammadanang/di_go_sample/config"
)

type SqlDb interface {
	Query(...interface{})
	Exec(...interface{})
}

type Sql struct {
	conf *config.Config
}

func InitDB(c *config.Config) *Sql {
	return &Sql{
		conf: c,
	}
}

func (s *Sql) Query(...interface{}) {
	fmt.Println("sql db query...", s.conf.DbName)
}

func (s *Sql) Exec(...interface{}) {
	fmt.Println("sql db exec...", s.conf.DbName)
}
