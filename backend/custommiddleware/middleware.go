package custommiddleware

import (
	"net/http"
	"os"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"github.com/markmumba/project-tracker/auth"
)

var jwtSecret = []byte(os.Getenv("JWT_SECRET"))

func Authentication(next echo.HandlerFunc) echo.HandlerFunc {

	return func(c echo.Context) error {

		cookie, err := c.Cookie("token")
		if err != nil || cookie == nil {
			c.JSON(http.StatusBadRequest, err.Error())
			return err
		}
		token, err := jwt.ParseWithClaims(cookie.Value, &auth.JwtCustomClaims{}, func(t *jwt.Token) (interface{}, error) {
			return jwtSecret, nil
		})
		if err != nil {
			return c.JSON(http.StatusUnauthorized, echo.Map{
				"message": "unauthorized",
				"error":   err.Error(),
			})
		}

		if claims, ok := token.Claims.(*auth.JwtCustomClaims); ok && token.Valid {
			c.Set("userId", uint(claims.UserId))
			c.Set("userRole", claims.UserRole)
			return next(c)

		} else {
			return c.JSON(http.StatusUnauthorized, echo.Map{
				"message": "unauthorized",
			})

		}
	}
}
