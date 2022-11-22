package main

import (
	"github.com/gin-gonic/gin"
	"github.com/sabmile/zhashkevych/employee/handler"
	"github.com/sabmile/zhashkevych/employee/storage"
)

func main() {
	ms := storage.NewMemoryStorage()
	handler := handler.NewHandler(ms)

	router := gin.Default()

	router.POST("/employee", handler.CreateEmployee)
	router.GET("/employee/:id", handler.GetEmployee)
	router.PUT("/employee/:id", handler.UpdateEmployee)
	router.DELETE("/employee/:id", handler.DeleteEmployee)
	router.GET("/employee", handler.GetEmployees)

	router.Run()
}
