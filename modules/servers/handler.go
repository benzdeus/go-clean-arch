package servers

import (
	"go-clean-arch/modules/students/controllers"
	"go-clean-arch/modules/students/repositories"
	"go-clean-arch/modules/students/usecases"
)

func (s *Server) MapHandlers() {
	apiV1 := s.Engine.Group("/api/v1")
	studentsGroup := apiV1.Group("/students")
	studentsRepository := repositories.NewStudentRepository(s.Db)
	studentsUsecase := usecases.NewStudentUsecase(studentsRepository)
	controllers.NewStudentsController(studentsGroup, studentsUsecase)

}
