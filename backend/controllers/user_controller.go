package controllers

import (
	"net/http"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/markmumba/project-tracker/models"
	"github.com/markmumba/project-tracker/services"
)

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
		Name:     "token",
		Value:    token,
		Expires:  time.Now().Add(time.Hour * 72),
		SameSite: http.SameSiteNoneMode,
		HttpOnly: true, // Disallow JavaScript access

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

	if err := services.CreateUser(&user); err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusCreated, models.UserToDTO(&user))
}

func GetUser(c echo.Context) error {

	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid ID")
	}

	user, err := services.GetUser(uint(id))
	if err != nil {
		return c.JSON(http.StatusNotFound, err.Error())
	}
	return c.JSON(http.StatusOK, models.UserToDTO(user))
}

func GetStudentsByLecturerId(c echo.Context) error {
	if c.Get("UserRole") != "lecturer" {
		return c.JSON(http.StatusUnauthorized, "Unauthorized")
	}
	userId := c.Get("userId").(uint)
	users, err := services.GetStudentsByLecturerId(userId)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, models.UserToDTOs(users))
}

func DeleteUser (c echo.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, "Invalid ID")
	}
	err = services.DeleteUser(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, err.Error())
	}
	c.JSON(http.StatusOK, "User deleted successfully")
}