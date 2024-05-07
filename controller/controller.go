package controller

import (
	"net/http"
	"os"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/openkrafter/anytore-backend/logger"
	"github.com/openkrafter/anytore-backend/model"
)

func SampleTraningItem(c *gin.Context) {
	trainingItem := new(model.TrainingItem)
	trainingItem.Id = 10
	trainingItem.UserId = 1
	trainingItem.Name = "running"
	trainingItem.Type = "aerobic"
	trainingItem.Unit = "hour"
	trainingItem.Kcal = 100

	c.JSON(http.StatusOK, trainingItem)
}

func RegisterRoutes(r *gin.Engine) {
	r.POST("/login", Login)

	r.GET("/sample", SampleTraningItem) // for debug
	r.GET("/training-items", ListTraningItem)
	r.GET("/training-items/:training-item-id", GetTraningItem)
	r.POST("/training-items", CreateTraningItem)
	r.PUT("/training-items/:training-item-id", UpdateTraningItem)
	r.DELETE("/training-items/:training-item-id", DeleteTraningItem)
}

func RegisterAdminRoutes(r *gin.Engine) {
	adminGroup := r.Group("/admin")
	username := os.Getenv("ADMIN_USERNAME")
	password := os.Getenv("ADMIN_PASSWORD")
	adminGroup.Use(gin.BasicAuth(gin.Accounts{
		username: password,
	}))

	adminGroup.GET("/users", ListUsers)
	adminGroup.GET("/users/:user-id", GetUser)
	adminGroup.POST("/users", CreateUser)
	adminGroup.PUT("/users/:user-id", UpdateUser)
	adminGroup.DELETE("/users/:user-id", DeleteUser)
}

func Run() {
	logger.Logger.Info("Controller thread start.")

	r := gin.Default()

	setCors(r)
	setCSP(r)
	RegisterRoutes(r)
	RegisterAdminRoutes(r)

	// listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
	if err := r.Run(); err != nil {
		logger.Logger.Error("Failed to start the server.", logger.ErrAttr(err))
		return
	}
}

func setCors(r *gin.Engine) {
	if os.Getenv("GIN_MODE") == "release" {
		r.Use(cors.New(cors.Config{
			AllowOrigins:     []string{os.Getenv("PROD_CORS_ORIGIN")},
			AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
			AllowHeaders:     []string{"Origin", "Content-Type", "Authorization", "Accept"},
			ExposeHeaders:    []string{"Content-Length"},
			AllowCredentials: true,
			MaxAge:           6 * time.Hour,
		}))
	} else {
		logger.Logger.Debug("CORS setting: debug mode")

		r.Use(cors.New(cors.Config{
			AllowOrigins:     []string{os.Getenv("DEV_CORS_ORIGIN")},
			AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
			AllowHeaders:     []string{"Origin", "Content-Type", "Authorization", "Accept"},
			ExposeHeaders:    []string{"Content-Length"},
			AllowCredentials: true,
			MaxAge:           6 * time.Hour,
		}))
	}
}

func setCSP(r *gin.Engine) {
	r.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Content-Security-Policy", "default-src 'self'; script-src 'self'; style-src 'self';")
		c.Next()
	})
}
