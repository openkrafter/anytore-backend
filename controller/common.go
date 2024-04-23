package controller

import (
	"strings"

	"github.com/gin-gonic/gin"
)

func GetTokenFromAuthorizationHeader(c *gin.Context) string {
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