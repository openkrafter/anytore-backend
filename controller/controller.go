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
	r.GET("/sample", SampleTraningItem) // for debug

	r.POST("/login", Login)

	r.GET("/admin/users", ListUsers)
	r.GET("/admin/users/:user-id", GetUser)
	r.POST("/admin/users", CreateUser)
	r.PUT("/admin/users/:user-id", UpdateUser)
	r.DELETE("/admin/users/:user-id", DeleteUser)

	r.GET("/training-items", ListTraningItem)
	r.GET("/training-items/:training-item-id", GetTraningItem)
	r.POST("/training-items", CreateTraningItem)
	r.PUT("/training-items/:training-item-id", UpdateTraningItem)
	r.DELETE("/training-items/:training-item-id", DeleteTraningItem)
}

func SetCors(r *gin.Engine) {
	logger.Logger.Debug("Setting CORS.")
	config := cors.DefaultConfig()
	if os.Getenv("CORS_ORIGIN") == "*" {
		logger.Logger.Debug("CORS setting: allow all origins")
		config.AllowAllOrigins = true
	} else {
		logger.Logger.Debug("CORS setting: allow specific origin")
		logger.Logger.Debug(os.Getenv("CORS_ORIGIN"))
		config.AllowOrigins = []string{os.Getenv("CORS_ORIGIN")}
		config.AllowCredentials = true
	}
	config.AllowMethods = []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}
	config.AllowHeaders = []string{"Origin", "Content-Type", "Authorization", "Accept"}
	config.ExposeHeaders = []string{"Content-Length"}
	config.MaxAge = 6 * time.Hour

	r.Use(cors.New(config))
}

func SetCSP(r *gin.Engine) {
	logger.Logger.Debug("Setting CSP.")
	r.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Content-Security-Policy", "default-src 'self'; script-src 'self'; style-src 'self';")
		c.Next()
	})
}

func Run() {
	logger.Logger.Info("Controller thread start.")

	r := gin.Default()

	SetCors(r)
	SetCSP(r)
	RegisterRoutes(r)

	// listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
	if err := r.Run(); err != nil {
		logger.Logger.Error("Failed to start the server.", logger.ErrAttr(err))
		return
	}
}
