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

func ListUser(c *gin.Context) {
	ctx := c.Request.Context()
	users, err := service.GetUsers(ctx)
	if err != nil {
		logger.Logger.Error("ListUser Failed.", logger.ErrAttr(err))
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
	userIdString := c.Param("user-id")
	userId, err := strconv.Atoi(userIdString)
	if err != nil {
		logger.Logger.Error("GetUser Failed. Failed to convert userId string to int.", logger.ErrAttr(err))
		return
	}

	ctx := c.Request.Context()
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
	var requestBody model.User
	if err := c.ShouldBindJSON(&requestBody); err != nil {
		logger.Logger.Error("CreateUser Failed. Failed to bind request body.", logger.ErrAttr(err))
		return
	}

	ctx := c.Request.Context()
	if err := service.CreateUser(ctx, &requestBody); err != nil {
		logger.Logger.Error("CreateUser Failed.", logger.ErrAttr(err))
		return
	}

	c.JSON(http.StatusCreated, requestBody)
}

func UpdateUser(c *gin.Context) {
	var requestBody model.User
	if err := c.ShouldBindJSON(&requestBody); err != nil {
		logger.Logger.Error("UpdateUser Failed. Failed to bind request body.", logger.ErrAttr(err))
		return
	}

	ctx := c.Request.Context()
	if err := service.UpdateUser(ctx, &requestBody); err != nil {
		logger.Logger.Error("UpdateUser Failed.", logger.ErrAttr(err))
		return
	}

	c.JSON(http.StatusCreated, requestBody)
}

func DeleteUser(c *gin.Context) {
	userIdString := c.Param("user-id")
	userId, err := strconv.Atoi(userIdString)
	if err != nil {
		logger.Logger.Error("DeleteUser Failed. Failed to convert userId string to int.", logger.ErrAttr(err))
		return
	}

	ctx := c.Request.Context()
	if err := service.DeleteUser(ctx, userId); err != nil {
		logger.Logger.Error("DeleteUser Failed.", logger.ErrAttr(err))
		return
	}

	c.JSON(http.StatusNoContent, gin.H{"message": "User deleted successfully."})
}
