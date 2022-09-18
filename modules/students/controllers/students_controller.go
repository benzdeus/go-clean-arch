package controllers

import (
	"go-clean-arch/modules/entities"

	"github.com/gin-gonic/gin"
)

type studentsController struct {
	StudentUsecase entities.StudentUsecase
}

func NewStudentsController(ctx *gin.RouterGroup, studentUsecase entities.StudentUsecase) {

	controllers := &studentsController{
		StudentUsecase: studentUsecase,
	}
	ctx.GET("/:id", controllers.GetStudentById)
	ctx.POST("", controllers.RegisterStudent)
	ctx.GET("/ping", controllers.Ping)
}

func (s *studentsController) RegisterStudent(c *gin.Context) {
	// Do some business logic here

	req := &entities.Student{}
	if err := c.ShouldBindJSON(req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	res, err := s.StudentUsecase.RegisterStudent(req)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{
		"data": res,
	})

}

func (s *studentsController) Ping(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}

func (s *studentsController) GetStudentById(c *gin.Context) {
	// Do some business logic here
	student, err := s.StudentUsecase.GetStudentById(c.Param("id"))

	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{
		"data": student,
	})
}
