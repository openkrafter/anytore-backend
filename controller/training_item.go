package controller

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/openkrafter/anytore-backend/customerror"
	"github.com/openkrafter/anytore-backend/logger"
	"github.com/openkrafter/anytore-backend/model"
	"github.com/openkrafter/anytore-backend/service"
)

func ListTraningItem(c *gin.Context) {
	userIdString := GetTokenFromAuthorizationHeader(c)
	userId, err := strconv.Atoi(userIdString)
	if err != nil {
		logger.Logger.Error("Failed to convert userId int to int.", logger.ErrAttr(err))
		return
	}

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

	var response []map[string]interface{}
	for _, trainingItem := range trainingItems {
		response = append(response, trainingItem.GetResponse())
	}
	log.Println(response)
	c.JSON(http.StatusOK, response)
}

func GetTraningItem(c *gin.Context) {
	userIdString := GetTokenFromAuthorizationHeader(c)
	userId, err := strconv.Atoi(userIdString)
	if err != nil {
		logger.Logger.Error("Failed to convert userId int to int.", logger.ErrAttr(err))
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

	response := trainingItem.GetResponse()
	c.JSON(http.StatusOK, response)
}

func CreateTraningItem(c *gin.Context) {
	var requestBody model.TrainingItem
	requestBodyBytes, err := io.ReadAll(c.Request.Body)
	if err != nil {
		logger.Logger.Error("UpdateTraningItem Failed.", logger.ErrAttr(err))
		return
	}

	err = json.Unmarshal(requestBodyBytes, &requestBody)
	if err != nil {
		logger.Logger.Error("UpdateTraningItem Failed.", logger.ErrAttr(err))
		return
	}

	ctx := c.Request.Context()
	err = service.CreateTraningItem(ctx, &requestBody)
	if err != nil {
		logger.Logger.Error("CreateTraningItem Failed.", logger.ErrAttr(err))
		return
	}

	c.JSON(http.StatusCreated, requestBody.GetResponse())
}

func UpdateTraningItem(c *gin.Context) {
	userIdString := GetTokenFromAuthorizationHeader(c)
	userId, err := strconv.Atoi(userIdString)
	if err != nil {
		logger.Logger.Error("Failed to convert userId int to int.", logger.ErrAttr(err))
		return
	}

	var requestBody model.TrainingItem
	requestBodyBytes, err := io.ReadAll(c.Request.Body)
	if err != nil {
		logger.Logger.Error("UpdateTraningItem Failed.", logger.ErrAttr(err))
		return
	}

	err = json.Unmarshal(requestBodyBytes, &requestBody)
	if err != nil {
		logger.Logger.Error("UpdateTraningItem Failed.", logger.ErrAttr(err))
		return
	}

	requestBody.Id, err = strconv.Atoi(c.Param("training-item-id"))
	if err != nil {
		logger.Logger.Error("UpdateTraningItem Failed.", logger.ErrAttr(err))
		return
	}

	ctx := c.Request.Context()
	err = service.UpdateTraningItem(ctx, &requestBody, userId)
	if err != nil {
		if customErr, ok := err.(*customerror.Error404); ok {
			c.JSON(customErr.ErrorCode, customErr.Body)
			logger.Logger.Error("UpdateTraningItem 404.", logger.ErrAttr(err))
		} else {
			logger.Logger.Error("UpdateTraningItem Failed.", logger.ErrAttr(err))
		}
		return
	}

	c.JSON(http.StatusCreated, requestBody.GetResponse())
}

func DeleteTraningItem(c *gin.Context) {
	userIdString := GetTokenFromAuthorizationHeader(c)
	userId, err := strconv.Atoi(userIdString)
	if err != nil {
		logger.Logger.Error("Failed to convert userId int to int.", logger.ErrAttr(err))
		return
	}

	trainingItemId, err := strconv.Atoi(c.Param("training-item-id"))
	if err != nil {
		logger.Logger.Error("DeleteTraningItem Failed.", logger.ErrAttr(err))
		return
	}

	ctx := c.Request.Context()
	err = service.DeleteTraningItem(ctx, trainingItemId, userId)
	if err != nil {
		if customErr, ok := err.(*customerror.Error404); ok {
			c.JSON(customErr.ErrorCode, customErr.Body)
			logger.Logger.Error("DeleteTraningItem 404.", logger.ErrAttr(err))
		} else {
			logger.Logger.Error("DeleteTraningItem Failed.", logger.ErrAttr(err))
		}
		return
	}
	c.JSON(http.StatusNoContent, "Delete to update.")
}
