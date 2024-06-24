package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name     string `json:"name"`
	Email    string `json:"email" gorm:"unique"`
	Password string `json:"password"`
	RoleId   uint   `json:"role_id"`
	Role     Role   `json:"role"`
}

type UserDTO struct {
	Id     uint   `json:"id"`
	Name   string `json:"name"`
	Email  string `json:"email"`
	RoleId uint   `json:"role"`
}

func UserToDTO(u *User) UserDTO {
	return UserDTO{
		Id:    u.ID,
		Name:  u.Name,
		Email: u.Email,
		RoleId: u.RoleId,
	}
}
func UserToDTOs(users []User) []UserDTO {
	userDTOs := make([]UserDTO, len(users))
	for i, user := range users {
		userDTOs[i] = UserToDTO(&user)
	}
	return userDTOs
}
