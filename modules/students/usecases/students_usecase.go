package usecases

import "go-clean-arch/modules/entities"

type studentUsecase struct {
	StudentRepository entities.StudentRepository
}

// Constructor
func NewStudentUsecase(studentRepository entities.StudentRepository) *studentUsecase {
	return &studentUsecase{
		StudentRepository: studentRepository,
	}
}

func (s *studentUsecase) RegisterStudent(req *entities.Student) (*entities.Student, error) {
	// Do some business logic here
	return s.StudentRepository.StoreStudent(req)
}

func (s *studentUsecase) GetStudentById(id string) (*entities.Student, error) {
	// Do some business logic here
	return s.StudentRepository.GetStudentById(id)
}
