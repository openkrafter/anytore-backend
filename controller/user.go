package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/openkrafter/anytore-backend/customerror"
	"github.com/openkrafter/anytore-backend/logger"
	"github.com/openkrafter/anytore-backend/model"
	"github.com/openkrafter/anytore-backend/service"
)

func ListUsers(c *gin.Context) {
	logger.Logger.Debug("ListUsers called.")
	tokenUserId, err := ValidateTokenAndGetUserId(c)
	if err != nil {
		logger.Logger.Error("Token Error.", logger.ErrAttr(err))
		c.JSON(http.StatusForbidden, gin.H{"error": "You are not authenticated"})
		return
	}

	ctx := c.Request.Context()
	err = CheckAdminAuthorization(ctx, tokenUserId)
	if err != nil {
		logger.Logger.Error("ListUsers Failed.", logger.ErrAttr(err))
		c.JSON(http.StatusForbidden, gin.H{"error": "You are not authorized"})
		return
	}

	users, err := service.GetUsers(ctx)
	if err != nil {
		logger.Logger.Error("ListUsers Failed.", logger.ErrAttr(err))
		return
	}
	if users == nil {
		error404 := customerror.NewError404()
		c.JSON(error404.ErrorCode, error404.Body)
		return
	}

	c.JSON(http.StatusOK, users)
}

func GetUser(c *gin.Context) {
	tokenUserId, err := ValidateTokenAndGetUserId(c)
	if err != nil {
		logger.Logger.Error("Token Error.", logger.ErrAttr(err))
		c.JSON(http.StatusForbidden, gin.H{"error": "You are not authenticated"})
		return
	}

	ctx := c.Request.Context()
	err = CheckAdminAuthorization(ctx, tokenUserId)
	if err != nil {
		logger.Logger.Error("ListUsers Failed.", logger.ErrAttr(err))
		c.JSON(http.StatusForbidden, gin.H{"error": "You are not authorized"})
		return
	}

	userIdString := c.Param("user-id")
	userId, err := strconv.Atoi(userIdString)
	if err != nil {
		logger.Logger.Error("GetUser Failed. Failed to convert userId string to int.", logger.ErrAttr(err))
		return
	}

	user, err := service.GetUserById(ctx, userId)
	if err != nil {
		logger.Logger.Error("GetUser Failed.", logger.ErrAttr(err))
		return
	}
	if user == nil {
		error404 := customerror.NewError404()
		c.JSON(error404.ErrorCode, error404.Body)
		return
	}

	c.JSON(http.StatusOK, user)
}

func CreateUser(c *gin.Context) {
	tokenUserId, err := ValidateTokenAndGetUserId(c)
	if err != nil {
		logger.Logger.Error("Token Error.", logger.ErrAttr(err))
		c.JSON(http.StatusForbidden, gin.H{"error": "You are not authenticated"})
		return
	}

	ctx := c.Request.Context()
	err = CheckAdminAuthorization(ctx, tokenUserId)
	if err != nil {
		logger.Logger.Error("ListUsers Failed.", logger.ErrAttr(err))
		c.JSON(http.StatusForbidden, gin.H{"error": "You are not authorized"})
		return
	}

	var requestBody model.User
	if err := c.ShouldBindJSON(&requestBody); err != nil {
		logger.Logger.Error("CreateUser Failed. Failed to bind request body.", logger.ErrAttr(err))
		return
	}

	if err := service.CreateUser(ctx, &requestBody); err != nil {
		logger.Logger.Error("CreateUser Failed.", logger.ErrAttr(err))
		return
	}

	c.JSON(http.StatusCreated, requestBody)
}

func UpdateUser(c *gin.Context) {
	tokenUserId, err := ValidateTokenAndGetUserId(c)
	if err != nil {
		logger.Logger.Error("Token Error.", logger.ErrAttr(err))
		c.JSON(http.StatusForbidden, gin.H{"error": "You are not authenticated"})
		return
	}

	ctx := c.Request.Context()
	err = CheckAdminAuthorization(ctx, tokenUserId)
	if err != nil {
		logger.Logger.Error("ListUsers Failed.", logger.ErrAttr(err))
		c.JSON(http.StatusForbidden, gin.H{"error": "You are not authorized"})
		return
	}

	var requestBody model.User
	if err := c.ShouldBindJSON(&requestBody); err != nil {
		logger.Logger.Error("UpdateUser Failed. Failed to bind request body.", logger.ErrAttr(err))
		return
	}

	requestBody.Id, err = strconv.Atoi(c.Param("user-id"))
	if err != nil {
		logger.Logger.Error("UpdateUser Failed. Failed to convert userId string to int.", logger.ErrAttr(err))
		return
	}

	if err := service.UpdateUser(ctx, &requestBody); err != nil {
		logger.Logger.Error("UpdateUser Failed.", logger.ErrAttr(err))
		return
	}

	c.JSON(http.StatusCreated, requestBody)
}

func DeleteUser(c *gin.Context) {
	tokenUserId, err := ValidateTokenAndGetUserId(c)
	if err != nil {
		logger.Logger.Error("Token Error.", logger.ErrAttr(err))
		c.JSON(http.StatusForbidden, gin.H{"error": "You are not authenticated"})
		return
	}

	ctx := c.Request.Context()
	err = CheckAdminAuthorization(ctx, tokenUserId)
	if err != nil {
		logger.Logger.Error("ListUsers Failed.", logger.ErrAttr(err))
		c.JSON(http.StatusForbidden, gin.H{"error": "You are not authorized"})
		return
	}

	userId, err := strconv.Atoi(c.Param("user-id"))
	if err != nil {
		logger.Logger.Error("DeleteUser Failed. Failed to convert userId string to int.", logger.ErrAttr(err))
		return
	}

	if err := service.DeleteUser(ctx, userId); err != nil {
		logger.Logger.Error("DeleteUser Failed.", logger.ErrAttr(err))
		return
	}

	c.JSON(http.StatusNoContent, gin.H{"message": "User deleted successfully."})
}
