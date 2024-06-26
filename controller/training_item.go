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

func ListTraningItem(c *gin.Context) {
	logger.Logger.Debug("ListTraningItem called.")

	userId, err := ValidateTokenAndGetUserId(c)
	if err != nil {
		logger.Logger.Error("Token Error.", logger.ErrAttr(err))
		c.JSON(http.StatusForbidden, gin.H{"error": "You are not authenticated"})
		return
	}
	logger.Logger.Debug("ListTraningItem userId", logger.Attr("userId", userId))

	ctx := c.Request.Context()
	trainingItems, err := service.GetTraningItems(ctx, userId)
	if err != nil {
		logger.Logger.Error("ListTraningItem Failed.", logger.ErrAttr(err))
		return
	}
	if trainingItems == nil {
		error404 := customerror.NewError404()
		c.JSON(error404.ErrorCode, error404.Body)
		return
	}
	logger.Logger.Debug("ListTraningItem trainingItems", logger.Attr("trainingItems", trainingItems))

	c.JSON(http.StatusOK, trainingItems)
}

func GetTraningItem(c *gin.Context) {
	userId, err := ValidateTokenAndGetUserId(c)
	if err != nil {
		logger.Logger.Error("Token Error.", logger.ErrAttr(err))
		c.JSON(http.StatusForbidden, gin.H{"error": "You are not authenticated"})
		return
	}

	trainingItemId, err := strconv.Atoi(c.Param("training-item-id"))
	if err != nil {
		logger.Logger.Error("GetTraningItem Failed.", logger.ErrAttr(err))
		return
	}

	ctx := c.Request.Context()
	trainingItem, err := service.GetTraningItem(ctx, trainingItemId, userId)
	if err != nil {
		logger.Logger.Error("GetTraningItem Failed.", logger.ErrAttr(err))
		return
	}
	if trainingItem == nil {
		error404 := customerror.NewError404()
		c.JSON(error404.ErrorCode, error404.Body)
		return
	}

	c.JSON(http.StatusOK, trainingItem)
}

func CreateTraningItem(c *gin.Context) {
	userId, err := ValidateTokenAndGetUserId(c)
	if err != nil {
		logger.Logger.Error("Token Error.", logger.ErrAttr(err))
		c.JSON(http.StatusForbidden, gin.H{"error": "You are not authenticated"})
		return
	}
	var requestBody model.TrainingItem
	if err := c.ShouldBindJSON(&requestBody); err != nil {
		logger.Logger.Error("CreateTraningItem Failed. Failed to bind request body.", logger.ErrAttr(err))
		return
	}

	requestBody.UserId = userId

	ctx := c.Request.Context()
	if err := service.CreateTraningItem(ctx, &requestBody); err != nil {
		logger.Logger.Error("CreateTraningItem Failed.", logger.ErrAttr(err))
		return
	}

	c.JSON(http.StatusCreated, requestBody)
}

func UpdateTraningItem(c *gin.Context) {
	userId, err := ValidateTokenAndGetUserId(c)
	if err != nil {
		logger.Logger.Error("Token Error.", logger.ErrAttr(err))
		c.JSON(http.StatusForbidden, gin.H{"error": "You are not authenticated"})
		return
	}

	var requestBody model.TrainingItem
	if err := c.ShouldBindJSON(&requestBody); err != nil {
		logger.Logger.Error("UpdateTraningItem Failed. Failed to bind request body.", logger.ErrAttr(err))
		return
	}

	requestBody.UserId = userId
	requestBody.Id, err = strconv.Atoi(c.Param("training-item-id"))
	if err != nil {
		logger.Logger.Error("UpdateTraningItem Failed. Failed to convert training-item-id string to int.", logger.ErrAttr(err))
		return
	}

	ctx := c.Request.Context()
	if err = service.UpdateTraningItem(ctx, &requestBody, userId); err != nil {
		if customErr, ok := err.(*customerror.Error404); ok {
			c.JSON(customErr.ErrorCode, customErr.Body)
			logger.Logger.Error("UpdateTraningItem 404.", logger.ErrAttr(err))
		} else {
			logger.Logger.Error("UpdateTraningItem Failed.", logger.ErrAttr(err))
		}
		return
	}

	c.JSON(http.StatusCreated, requestBody)
}

func DeleteTraningItem(c *gin.Context) {
	userId, err := ValidateTokenAndGetUserId(c)
	if err != nil {
		logger.Logger.Error("Token Error.", logger.ErrAttr(err))
		c.JSON(http.StatusForbidden, gin.H{"error": "You are not authenticated"})
		return
	}

	trainingItemId, err := strconv.Atoi(c.Param("training-item-id"))
	if err != nil {
		logger.Logger.Error("DeleteTraningItem Failed.", logger.ErrAttr(err))
		return
	}

	ctx := c.Request.Context()
	if err = service.DeleteTraningItem(ctx, trainingItemId, userId); err != nil {
		if customErr, ok := err.(*customerror.Error404); ok {
			c.JSON(customErr.ErrorCode, customErr.Body)
			logger.Logger.Error("DeleteTraningItem 404.", logger.ErrAttr(err))
		} else {
			logger.Logger.Error("DeleteTraningItem Failed.", logger.ErrAttr(err))
		}
		return
	}
	c.JSON(http.StatusNoContent, gin.H{"message": "TrainingItem deleted successfully."})
}
