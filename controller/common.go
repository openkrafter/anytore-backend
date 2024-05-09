package controller

import (
	"context"
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/openkrafter/anytore-backend/logger"
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

func CheckAdminAuthorization(ctx context.Context, userId int) error {
	user, err := service.GetUserById(ctx, userId)
	if err != nil {
		logger.Logger.Error("Failed to get user.", logger.ErrAttr(err))
		return err
	}

	if user.Name != os.Getenv("ADMIN_NAME") {
		errMsg := fmt.Sprintf("User %d (%s) is not admin.", userId, user.Name)
		logger.Logger.Error(errMsg, logger.ErrAttr(err))
		return errors.New(errMsg)
	}
	return nil
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
