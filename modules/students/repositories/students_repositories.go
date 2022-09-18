package repositories

import (
	"go-clean-arch/modules/entities"

	"github.com/jinzhu/gorm"
)

type studentRepository struct {
	Db *gorm.DB
}

func NewStudentRepository(db *gorm.DB) *studentRepository {
	return &studentRepository{Db: db}
}

func (r *studentRepository) StoreStudent(req *entities.Student) (*entities.Student, error) {
	err := r.Db.Create(req).Error
	if err != nil {
		return nil, err
	}
	return req, nil
}

func (r *studentRepository) GetStudentById(id string) (*entities.Student, error) {
	var student entities.Student
	err := r.Db.Where("id = ?", id).First(&student).Error
	if err != nil {
		return nil, err
	}
	return &student, nil
}
