package dto

import "XUTAPD/models"

//只发送必要的信息给前端
type StudentDTO struct {
	StudentName string `json:"student_name"`
}

func ToStudentDTO(student models.Student) StudentDTO {
	return StudentDTO{
		StudentName: student.StudentName,
	}
}