package controller

import (
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/openkrafter/anytore-backend/service"
)

func ValidateTokenAndGetUserId(c *gin.Context) (int, error) {
	receivedToken := getTokenFromAuthorizationHeader(c)
	userId, err := service.ValidateToken(receivedToken)
	if err != nil {
		return 0, err
	}

	return userId, nil
}

func getTokenFromAuthorizationHeader(c *gin.Context) string {
	authHeader := c.GetHeader("Authorization")
	if authHeader == "" {
		return ""
	}

	// values[0] == Bearer, values[1] == Token
	values := strings.Split(authHeader, " ")
	if len(values) != 2 || strings.ToLower(values[0]) != "bearer" {
		return ""
	}

	return values[1]
}
