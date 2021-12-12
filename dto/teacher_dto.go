package dto

import "XUTAPD/models"

//只发送必要的信息给前端
type TeacherDTO struct {
	TeacherName string `json:"teacher_name"`
}

func ToTeacherDTO(teacher models.Teacher) TeacherDTO {
	return TeacherDTO{
		TeacherName: teacher.TeacherName,
	}
}