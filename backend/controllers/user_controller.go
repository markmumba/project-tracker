package controllers

import (

	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/markmumba/project-tracker/models"
	"github.com/markmumba/project-tracker/services"
)
func HomePage (c echo.Context) error {
	return c.Render(http.StatusOK, "base", nil)
}

func Login(c echo.Context) error {
	var credentials struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := c.Bind(&credentials); err != nil {
		return c.Render(http.StatusBadRequest, "error.html", map[string]interface{}{
			"error": err.Error(),
		})
	}

	token, err := services.LoginUser(credentials.Email, credentials.Password)
	if err != nil {
		return c.Render(http.StatusUnauthorized, "error.html", map[string]interface{}{
			"error": err.Error(),
		})
	}
	c.SetCookie(&http.Cookie{
		Name:     "token",
		Value:    token,
		Expires:  time.Now().Add(time.Hour * 72),
		SameSite: http.SameSiteNoneMode,
		HttpOnly: true, // Disallow JavaScript access
	})

	return c.Redirect(http.StatusFound, "/")
}

func Logout(c echo.Context) error {
	cookie := &http.Cookie{
		Name:    "token",
		Value:   "",
		Expires: time.Now().Add(-time.Hour),
	}

	c.SetCookie(cookie)
	return c.Render(http.StatusOK, "logout_success.html", map[string]interface{}{
		"message": "Logout successful",
	})
}

func CreateUser(c echo.Context) error {
	var user models.User
	if err := c.Bind(&user); err != nil {
		return c.Render(http.StatusBadRequest, "error.html", map[string]interface{}{
			"error": err.Error(),
		})
	}

	if err := services.CreateUser(&user); err != nil {
		return c.Render(http.StatusInternalServerError, "error.html", map[string]interface{}{
			"error": err.Error(),
		})
	}

	return c.Redirect(http.StatusFound, "/")
}

func GetUser(c echo.Context) error {
	var userParams models.User
	err := c.Bind(&userParams)
	if err != nil {
		return c.Render(http.StatusBadRequest, "error.html", map[string]interface{}{
			"error": err.Error(),
		})
	}
	id := userParams.ID
	user, err := services.GetUser(uint(id))
	if err != nil {
		return c.Render(http.StatusNotFound, "error.html", map[string]interface{}{
			"error": err.Error(),
		})
	}
	return c.Render(http.StatusOK, "user_details.html", models.UserToDTO(user))
}

func GetStudentsByLecturerId(c echo.Context) error {
	var userParams models.User
	err := c.Bind(&userParams)
	if err != nil {
		return c.Render(http.StatusBadRequest, "error.html", map[string]interface{}{
			"error": err.Error(),
		})
	}
	id := userParams.ID
	students, err := services.GetStudentsByLecturerId(uint(id))
	if err != nil {
		return c.Render(http.StatusNotFound, "error.html", map[string]interface{}{
			"error": err.Error(),
		})
	}
	return c.Render(http.StatusOK, "students_list.html", models.UserToDTOs(students))
}

func DeleteUser(c echo.Context) error {
	var userParams models.User
	err := c.Bind(&userParams)
	if err != nil {
		return c.Render(http.StatusBadRequest, "error.html", map[string]interface{}{
			"error": err.Error(),
		})
	}
	id := userParams.ID
	err = services.DeleteUser(uint(id))
	if err != nil {
		return c.Render(http.StatusNotFound, "error.html", map[string]interface{}{
			"error": err.Error(),
		})
	}
	return c.Render(http.StatusOK, "delete_success.html", map[string]interface{}{
		"message": "User deleted successfully",
	})
}
