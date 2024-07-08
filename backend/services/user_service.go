package services

import (
	"errors"

	"github.com/markmumba/project-tracker/auth"
	"github.com/markmumba/project-tracker/database"
	"github.com/markmumba/project-tracker/models"
)

func CreateUser(user *models.User) error {
	hashedPassword, err := auth.HashPassword(user.Password)
	if err != nil {
		return err
	}
	user.Password = hashedPassword
	result := database.DB.Create(user)
	return result.Error
}

func LoginUser(email, password string) (string, error) {
	var user models.User
	err := database.DB.Where("email = ?", email).First(&user).Error
	if err != nil {
		return "", err
	}

	if !auth.CheckPasswordHash(password, user.Password) {
		return "", errors.New("invalid credentials")
	}

	token, err := auth.GenerateJWT(&user)
	if err != nil {
		return "", err
	}
	return token, nil
}
func GetUser(id uint) (*models.User, error) {
	var user models.User
	result := database.DB.Preload("Role").First(&user, id)
	return &user, result.Error
}

func GetAllUsers() ([]models.User, error) {
	var users []models.User
	result := database.DB.Find(&users)
	return users, result.Error
}

func GetStudentsByLecturer(lecturerID uint) ([]models.User, error) {
	var projects []models.Project
	err := database.DB.Preload("Student").Where("lecturer_id = ?", lecturerID).Find(&projects).Error
	if err != nil {
		return nil, err
	}

	var students []models.User
	for _, project := range projects {
		students = append(students, project.Student)
	}

	return students, nil
}

func GetLecturers() ([]models.User, error) {
	var lecturers []models.User
	result := database.DB.Preload("Role").Where("role_id = 1").Find(&lecturers)
	return lecturers, result.Error
}

func UpdateUser(id uint, user *models.User) error {
	result := database.DB.Save(user).Where("id = ?", id)
	return result.Error
}

func DeleteUser(id uint) error {
	var user models.User
	result := database.DB.Delete(&user, id)
	return result.Error
}
