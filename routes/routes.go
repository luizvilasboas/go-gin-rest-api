package routes

import (
	"github.com/gin-gonic/gin"
	"gitlab.com/alura-courses-code/golang/go-gin-rest-api/controllers"
)

func HandleRequests() {
	r := gin.Default()
	r.GET("/api/students", controllers.HandleGetAllStudents)
	r.GET("/api/students/:id", controllers.HandleGetStudent)
	r.POST("/api/students/create", controllers.HandleCreateStudent)
	r.DELETE("/api/students/delete/:id", controllers.HandleDeleteStudent)
	r.PATCH("/api/students/update/:id", controllers.HandleUpdateStudent)
	r.GET("/api/students/search/:cpf", controllers.HandleSearchStudentCPF)
	r.Run()
}
