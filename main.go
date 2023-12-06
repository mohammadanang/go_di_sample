package main

import (
	"fmt"

	"github.com/mohammadanang/di_go_sample/config"
	"github.com/mohammadanang/di_go_sample/database"
	"github.com/mohammadanang/di_go_sample/simple/controller"
	"github.com/mohammadanang/di_go_sample/simple/repository"
	"github.com/mohammadanang/di_go_sample/simple/service"
)

func main() {
	conf := config.InitConfig()
	db := database.InitDB(conf)
	repo := repository.NewRepository(db)
	service := service.NewService(repo)
	controller := controller.NewController(service)

	controller.List()
	controller.Show()
	controller.Create()
	controller.Update()
	controller.Delete()

	fmt.Println("server running...")
}
