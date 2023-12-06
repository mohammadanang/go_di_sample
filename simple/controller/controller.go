package controller

import (
	"fmt"

	"github.com/mohammadanang/di_go_sample/simple/service"
)

type Controller struct {
	service service.SimpleService
}

func NewController(s service.SimpleService) *Controller {
	return &Controller{
		service: s,
	}
}

func (c *Controller) List(...interface{}) {
	c.service.FindAll()
	fmt.Println("get list from controller...")
}

func (c *Controller) Show(...interface{}) {
	c.service.Find()
	fmt.Println("get list from controller...")
}

func (c *Controller) Create(...interface{}) {
	c.service.Insert()
	fmt.Println("get list from controller...")
}

func (c *Controller) Update(...interface{}) {
	c.service.Edit()
	fmt.Println("get list from controller...")
}

func (c *Controller) Delete(...interface{}) {
	c.service.Remove()
	fmt.Println("get list from controller...")
}
