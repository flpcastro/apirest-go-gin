package routes

import (
	"github.com/flpcastro/apirest-go-gin/controllers"
	"github.com/gin-gonic/gin"
)

func HandleRequest() {
	r := gin.Default()
	r.LoadHTMLGlob("templates/*")
	r.Static("/assets", "./assets")
	r.GET("/:name", controllers.Greetings)
	r.GET("/students", controllers.AllStudents)
	r.GET("/students/:id", controllers.GetStudentById)
	r.POST("/students", controllers.CreateNewStudent)
	r.DELETE("/students/:id", controllers.DeleteStudent)
	r.PATCH("/students/:id", controllers.EditStudent)
	r.GET("/students/cpf/:cpf", controllers.GetStudentByCPF)
	r.GET("/students/", controllers.GetStudentByCPF)
	r.GET("/index", controllers.ShowPageIndex)
	r.NoRoute(controllers.RouteNotFound)
	r.Run()
}
