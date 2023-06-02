package middleware

import (
	"capstone/constant"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
)

func CreateToken(userId uint, role string) (string, error) {
	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["userId"] = userId
	claims["role"] = role

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(constant.SECRET_JWT))
}

func ExtractTokenUserId(e echo.Context) (uint, error) {
	user := e.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	if user.Valid {
		userId := uint(claims["userId"].(float64))
		if claims["role"] == "USER" {
			return userId, nil
		}
	}
	return 0, echo.NewHTTPError(http.StatusUnauthorized, "Not Authorized")
}

func ExtractTokenAdminId(e echo.Context) (string, uint, error) {
	user := e.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	if user.Valid {
		userId := uint(claims["userId"].(float64))
		role := claims["role"].(string)
		return role, userId, nil

	}
	return "", 0, echo.NewHTTPError(http.StatusUnauthorized, "Not Authorized")
}
