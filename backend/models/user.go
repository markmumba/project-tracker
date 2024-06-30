package models

import (

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name     string `json:"name"`
	Email    string `json:"email" gorm:"unique"`
	Password string `json:"password"`
	RoleID   uint   `json:"role_id"`
	Role     Role `gorm:"foreignKey:RoleID"`
}

type UserDTO struct {
	Id 	 uint   `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Role     string `json:"role"`
}


func UserToDTO(u *User) UserDTO {
	return UserDTO{
		Id : u.ID,
		Name:  u.Name,
		Email: u.Email,
		Role:  u.Role.Name,
	}
}
func UserToDTOs(users []User) []UserDTO {
	userDTOs := make([]UserDTO, len(users))
	for i, user := range users {
		userDTOs[i] = UserToDTO(&user)
	}
	return userDTOs
}
