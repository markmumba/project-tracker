package controllers

import (
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/markmumba/project-tracker/helpers"
	"github.com/markmumba/project-tracker/models"
	"github.com/markmumba/project-tracker/services"
)

// TODO : Get all the lecturers that is get all users where role id is 1

func Login(c echo.Context) error {
	var credentials struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := c.Bind(&credentials); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	token, err := services.LoginUser(credentials.Email, credentials.Password)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, err.Error())
	}
	c.SetCookie(&http.Cookie{
		Name:    "token",
		Value:   token,
		Expires: time.Now().Add(time.Hour * 72),
	})

	return c.JSON(http.StatusOK, echo.Map{
		"token": token,
	})
}
func Logout(c echo.Context) error {
	cookie := &http.Cookie{
		Name:    "token",
		Value:   "",
		Expires: time.Now().Add(-time.Hour),
	}

	c.SetCookie(cookie)
	return c.JSON(http.StatusOK, echo.Map{
		"message": "Logout successful",
	})
}

func CreateUser(c echo.Context) error {
	var user models.User
	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	if err := services.CreateUser(&user); err != nil { // Print type for debugging

		// Safely assert and handle the userID type
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusCreated, models.UserToDTO(&user))
}

func GetUser(c echo.Context) error {
	userID, err := helpers.ConvertUserID(c, "userId")
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	user, err := services.GetUser(userID)
	if err != nil {
		return c.JSON(http.StatusNotFound, err.Error())
	}

	return c.JSON(http.StatusOK, models.UserToDTO(user))
}

func GetAllUsers(c echo.Context) error {
	users, err := services.GetAllUsers()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, models.UserToDTOs(users))
}
func GetLecturers(c echo.Context) error {
	lecturers, err := services.GetLecturers()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, models.UserToDTOs(lecturers))
}

func GetStudentsByLecturerId(c echo.Context) error {

	userID, err := helpers.ConvertUserID(c, "userId")
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	students, err := services.GetStudentsByLecturer(userID)
	if err != nil {
		return c.JSON(http.StatusNotFound, err.Error())
	}
	return c.JSON(http.StatusOK, models.UserToDTOs(students))
}
func UpdateUser(c echo.Context) error {
	userID, err := helpers.ConvertUserID(c, "userId")
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	var updateUser models.User
	if err := c.Bind(&updateUser); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	if err := services.UpdateUser(userID, &updateUser); err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, models.UserToDTO(&updateUser))
}

func UpdateUserProfileImage(c echo.Context) error {
	var image struct {
		ProfileImage string `json:"profile_image"`
	}
	userID, err := helpers.ConvertUserID(c, "userId")
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	if err := c.Bind(&image); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	err = services.UpdateUserProfileImage(userID, image.ProfileImage)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, "Profile image updated successfully")
}


func DeleteUser(c echo.Context) error {

	userID, err := helpers.ConvertUserID(c, "userId")
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	err = services.DeleteUser(userID)
	if err != nil {
		return c.JSON(http.StatusNotFound, err.Error())
	}
	return c.JSON(http.StatusOK, "User deleted successfully")
}
