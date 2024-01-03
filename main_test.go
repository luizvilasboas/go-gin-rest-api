package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"gitlab.com/alura-courses-code/golang/go-gin-rest-api/controllers"
	"gitlab.com/alura-courses-code/golang/go-gin-rest-api/database"
	"gitlab.com/alura-courses-code/golang/go-gin-rest-api/models"
)

func SetupTestRoutes() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	routes := gin.Default()
	return routes
}

var ID int

func CreateStudentMock() {
	student := models.Student{
		Name: "Test",
		Age:  20,
		CPF:  "12345678901",
		RG:   "123456789",
	}

	database.DB.Create(&student)
	ID = int(student.ID)
}

func DeleteStudentMock() {
	var student models.Student
	database.DB.Delete(&student, ID)
}

func TestHandleGetAllStudents(t *testing.T) {
	database.ConnectDatabase()

	CreateStudentMock()
	defer DeleteStudentMock()

	r := SetupTestRoutes()
	r.GET("/api/students", controllers.HandleGetAllStudents)

	req, _ := http.NewRequest("GET", "/api/students", nil)
	res := httptest.NewRecorder()

	r.ServeHTTP(res, req)

	assert.Equal(t, http.StatusOK, res.Code)
}

func TestHandleGetStudent(t *testing.T) {
	database.ConnectDatabase()

	CreateStudentMock()
	defer DeleteStudentMock()

	r := SetupTestRoutes()
	r.GET("/api/students/:id", controllers.HandleGetStudent)

	req, _ := http.NewRequest("GET", "/api/students/"+strconv.Itoa(ID), nil)
	res := httptest.NewRecorder()

	r.ServeHTTP(res, req)

	var studentMock models.Student
	json.Unmarshal(res.Body.Bytes(), &studentMock)

	assert.Equal(t, http.StatusOK, res.Code)
	assert.Equal(t, "Test", studentMock.Name)
	assert.Equal(t, 20, studentMock.Age)
	assert.Equal(t, "12345678901", studentMock.CPF)
	assert.Equal(t, "123456789", studentMock.RG)
}

func TestHandleDeleteStudent(t *testing.T) {
	database.ConnectDatabase()

	CreateStudentMock()

	r := SetupTestRoutes()
	r.DELETE("/api/students/:id", controllers.HandleDeleteStudent)

	path := "/api/students/" + strconv.Itoa(ID)

	req, _ := http.NewRequest("DELETE", path, nil)
	res := httptest.NewRecorder()
	r.ServeHTTP(res, req)

	assert.Equal(t, http.StatusOK, res.Code)
	assert.Equal(t, "{\"id\":\""+strconv.Itoa(ID)+"\",\"message\":\"Student deleted successfully\"}", res.Body.String())
}
