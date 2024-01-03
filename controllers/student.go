package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gitlab.com/alura-courses-code/golang/go-gin-rest-api/database"
	"gitlab.com/alura-courses-code/golang/go-gin-rest-api/models"
)

func HandleGetAllStudents(c *gin.Context) {
	var students []models.Student
	database.DB.Find(&students)
	c.JSON(http.StatusOK, students)
}

func HandleGetStudent(c *gin.Context) {
	var student models.Student
	id := c.Params.ByName("id")
	database.DB.First(&student, id)

	if student.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"message":   "Student not found",
			"errorCode": "404",
		})

		return
	}

	c.JSON(http.StatusOK, student)
}

func HandleCreateStudent(c *gin.Context) {
	var student models.Student

	if err := c.ShouldBindJSON(&student); err != nil {
		c.JSON(400, gin.H{
			"message":   err.Error(),
			"errorCode": "400",
		})

		return
	}

	if err := models.ValidateStudentData(&student); err != nil {
		c.JSON(400, gin.H{
			"message":   err.Error(),
			"errorCode": "400",
		})
	}

	database.DB.Create(&student)
	c.JSON(http.StatusOK, student)
}

func HandleDeleteStudent(c *gin.Context) {
	var student models.Student
	id := c.Params.ByName("id")
	database.DB.Delete(&student, id)

	c.JSON(http.StatusOK, gin.H{
		"id":      id,
		"message": "Student deleted successfully",
	})
}

func HandleUpdateStudent(c *gin.Context) {
	var student models.Student
	id := c.Params.ByName("id")
	database.DB.First(&student, id)

	if err := c.ShouldBindJSON(&student); err != nil {
		c.JSON(400, gin.H{
			"message":   err.Error(),
			"errorCode": "400",
		})

		return
	}

	if err := models.ValidateStudentData(&student); err != nil {
		c.JSON(400, gin.H{
			"message":   err.Error(),
			"errorCode": "400",
		})
	}

	database.DB.Model(&student).UpdateColumns(student)
	c.JSON(http.StatusOK, student)
}

func HandleSearchStudentCPF(c *gin.Context) {
	var student models.Student
	cpf := c.Params.ByName("cpf")
	database.DB.Where(&models.Student{CPF: cpf}).First(&student)

	if student.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"message":   "Student not found",
			"errorCode": "404",
		})

		return
	}

	c.JSON(http.StatusOK, student)
}
