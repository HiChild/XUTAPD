package dto

import "XUTAPD/models"

//只发送必要的信息给前端
type UserDTO struct {
	Username string `json:"username"`
}

func ToUserDTO(user models.User) UserDTO {
	return UserDTO{
		Username: user.UserName,
	}
}