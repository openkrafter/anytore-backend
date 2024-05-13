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

func ListTrainingRecords(c *gin.Context) {
	userId, err := ValidateTokenAndGetUserId(c)
	if err != nil {
		logger.Logger.Error("Token Error.", logger.ErrAttr(err))
		c.JSON(http.StatusForbidden, gin.H{"error": "You are not authenticated"})
		return
	}

	ctx := c.Request.Context()
	trainingRecords, err := service.GetTrainingRecords(ctx, userId)
	if err != nil {
		logger.Logger.Error("ListTrainingRecords Failed.", logger.ErrAttr(err))
		return
	}
	if trainingRecords == nil {
		error404 := customerror.NewError404()
		c.JSON(error404.ErrorCode, error404.Body)
		return
	}

	c.JSON(http.StatusOK, trainingRecords)
}

func GetTrainingRecord(c *gin.Context) {
	userId, err := ValidateTokenAndGetUserId(c)
	if err != nil {
		logger.Logger.Error("Token Error.", logger.ErrAttr(err))
		c.JSON(http.StatusForbidden, gin.H{"error": "You are not authenticated"})
		return
	}

	trainingRecordId, err := strconv.Atoi(c.Param("training-record-id"))
	if err != nil {
		logger.Logger.Error("GetTrainingRecord Failed.", logger.ErrAttr(err))
		return
	}

	ctx := c.Request.Context()
	trainingRecord, err := service.GetTrainingRecord(ctx, trainingRecordId, userId)
	if err != nil {
		logger.Logger.Error("GetTrainingRecord Failed.", logger.ErrAttr(err))
		return
	}

	c.JSON(http.StatusOK, trainingRecord)
}

func CreateTrainingRecord(c *gin.Context) {
	userId, err := ValidateTokenAndGetUserId(c)
	if err != nil {
		logger.Logger.Error("Token Error.", logger.ErrAttr(err))
		c.JSON(http.StatusForbidden, gin.H{"error": "You are not authenticated"})
		return
	}

	trainingRecord := new(model.TrainingRecord)
	if err := c.BindJSON(trainingRecord); err != nil {
		logger.Logger.Error("CreateTrainingRecord Failed.", logger.ErrAttr(err))
		return
	}

	trainingRecord.UserId = userId

	ctx := c.Request.Context()
	err = service.CreateTrainingRecord(ctx, trainingRecord)
	if err != nil {
		logger.Logger.Error("CreateTrainingRecord Failed.", logger.ErrAttr(err))
		return
	}

	c.JSON(http.StatusCreated, trainingRecord)
}

func UpdateTrainingRecord(c *gin.Context) {
	userId, err := ValidateTokenAndGetUserId(c)
	if err != nil {
		logger.Logger.Error("Token Error.", logger.ErrAttr(err))
		c.JSON(http.StatusForbidden, gin.H{"error": "You are not authenticated"})
		return
	}

	trainingRecord := new(model.TrainingRecord)
	if err := c.BindJSON(trainingRecord); err != nil {
		logger.Logger.Error("UpdateTrainingRecord Failed.", logger.ErrAttr(err))
		return
	}

	trainingRecord.UserId = userId
	trainingRecord.Id, err = strconv.Atoi(c.Param("training-record-id"))
	if err != nil {
		logger.Logger.Error("UpdateTrainingRecord Failed.", logger.ErrAttr(err))
		return
	}

	ctx := c.Request.Context()
	err = service.UpdateTrainingRecord(ctx, trainingRecord, userId)
	if err != nil {
		if customErr, ok := err.(*customerror.Error404); ok {
			c.JSON(customErr.ErrorCode, customErr.Body)
			logger.Logger.Error("UpdateTraningRecord 404.", logger.ErrAttr(err))
		} else {
			logger.Logger.Error("UpdateTraningRecord Failed.", logger.ErrAttr(err))
		}
		return
	}

	c.JSON(http.StatusCreated, trainingRecord)
}

func DeleteTrainingRecord(c *gin.Context) {
	userId, err := ValidateTokenAndGetUserId(c)
	if err != nil {
		logger.Logger.Error("Token Error.", logger.ErrAttr(err))
		c.JSON(http.StatusForbidden, gin.H{"error": "You are not authenticated"})
		return
	}

	trainingRecordId, err := strconv.Atoi(c.Param("training-record-id"))
	if err != nil {
		logger.Logger.Error("DeleteTrainingRecord Failed.", logger.ErrAttr(err))
		return
	}

	ctx := c.Request.Context()
	err = service.DeleteTrainingRecord(ctx, trainingRecordId, userId)
	if err != nil {
		if customErr, ok := err.(*customerror.Error404); ok {
			c.JSON(customErr.ErrorCode, customErr.Body)
			logger.Logger.Error("DeleteTrainingRecord 404.", logger.ErrAttr(err))
		} else {
			logger.Logger.Error("DeleteTrainingRecord Failed.", logger.ErrAttr(err))
		}
		return
	}

	c.JSON(http.StatusNoContent, gin.H{"message": "TrainingRecord deleted successfully."})
}
