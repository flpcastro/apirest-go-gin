package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/flpcastro/apirest-go-gin/controllers"
	"github.com/flpcastro/apirest-go-gin/database"
	"github.com/flpcastro/apirest-go-gin/models"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

var ID int

func TestRoutesSetup() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	rotas := gin.Default()
	return rotas
}

func CreateStudentMock() {
	student := models.Student{Name: "Name Test", CPF: "12345678901", RG: "123456789"}
	database.DB.Create(&student)
	ID = int(student.ID)
}

func DeleteStudentMock() {
	var student models.Student
	database.DB.Delete(&student, ID)
}

func TestVerifyStatusCodeOfGreetingsWithParameter(t *testing.T) {
	r := TestRoutesSetup()
	r.GET("/:name", controllers.Greetings)
	req, _ := http.NewRequest("GET", "/gui", nil)
	res := httptest.NewRecorder()
	r.ServeHTTP(res, req)
	assert.Equal(t, http.StatusOK, res.Code, "Must be equal")
	resMock := `{"API say":"Hey, whats up?"}`
	resBody, _ := ioutil.ReadAll(res.Body)
	assert.Equal(t, resMock, string(resBody))
}

func TestListAllStudentsHandler(t *testing.T) {
	database.DatabaseConnect()
	CreateStudentMock()
	defer DeleteStudentMock()
	r := TestRoutesSetup()
	r.GET("/students", controllers.AllStudents)
	req, _ := http.NewRequest("GET", "/students", nil)
	res := httptest.NewRecorder()
	r.ServeHTTP(res, req)
	assert.Equal(t, http.StatusOK, res.Code)
}

func TestGetStudentByCPFHandler(t *testing.T) {
	database.DatabaseConnect()
	CreateStudentMock()
	defer DeleteStudentMock()
	r := TestRoutesSetup()
	r.GET("/students/cpf/:cpf", controllers.GetStudentByCPF)
	req, _ := http.NewRequest("GET", "/students/cpf/12345678901", nil)
	res := httptest.NewRecorder()
	r.ServeHTTP(res, req)
	assert.Equal(t, http.StatusOK, res.Code)
}

func TestGetStudentByIDHandler(t *testing.T) {
	database.DatabaseConnect()
	CreateStudentMock()
	defer DeleteStudentMock()
	r := TestRoutesSetup()
	r.GET("/students/:id", controllers.GetStudentById)
	searchPath := "/students/" + strconv.Itoa(ID)
	req, _ := http.NewRequest("GET", searchPath, nil)
	res := httptest.NewRecorder()
	r.ServeHTTP(res, req)
	var studentMock models.Student
	json.Unmarshal(res.Body.Bytes(), &studentMock)
	assert.Equal(t, "Student Test Name", studentMock.Name, "Name must be equal")
	assert.Equal(t, "12345678901", studentMock.CPF)
	assert.Equal(t, "123456789", studentMock.RG)
	assert.Equal(t, http.StatusOK, res.Code)
}

func TestDeleteStudentHandler(t *testing.T) {
	database.DatabaseConnect()
	CreateStudentMock()
	r := TestRoutesSetup()
	r.DELETE("/students/:id", controllers.DeleteStudent)
	searchPath := "/students/" + strconv.Itoa(ID)
	req, _ := http.NewRequest("DELETE", searchPath, nil)
	res := httptest.NewRecorder()
	r.ServeHTTP(res, req)
	assert.Equal(t, http.StatusOK, res.Code)
}

func TestEditStudentHandler(t *testing.T) {
	database.DatabaseConnect()
	CreateStudentMock()
	defer DeleteStudentMock()
	r := TestRoutesSetup()
	r.PATCH("/students/:id", controllers.EditStudent)
	student := models.Student{Name: "Student Test Name", CPF: "47123456789", RG: "123456700"}
	jsonValue, _ := json.Marshal(student)
	editPath := "/students/" + strconv.Itoa(ID)
	req, _ := http.NewRequest("PATCH", editPath, bytes.NewBuffer(jsonValue))
	res := httptest.NewRecorder()
	r.ServeHTTP(res, req)
	var studentUpdatedMock models.Student
	json.Unmarshal(res.Body.Bytes(), &studentUpdatedMock)
	assert.Equal(t, "47123456789", studentUpdatedMock.CPF)
	assert.Equal(t, "123456700", studentUpdatedMock.RG)
	assert.Equal(t, "Student Name Test", studentUpdatedMock.Name)
}
