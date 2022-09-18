package entities

type StudentUsecase interface {
	RegisterStudent(req *Student) (*Student, error) // RegisterStudent is a method to register a new student
	GetStudentById(id string) (*Student, error)     // GetStudentById is a method to get a student by id
}

type StudentRepository interface {
	StoreStudent(req *Student) (*Student, error) // StoreStudent is a method to store a new student
	GetStudentById(id string) (*Student, error)  // GetStudentById is a method to get a student by id
}

type Student struct {
	Id        uint   `json:"id"`
	Username  string `json:"username" binding:"required"`
	Password  string `json:"password" gorm:"omitempty"  binding:"required"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}
