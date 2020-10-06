package http

import (
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func Run() (err error) {
	router := gin.New()
	if os.Getenv("GIN_MODE") != "" {
		gin.SetMode(os.Getenv("GIN_MODE"))
	}

	// CORS
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowAllOrigins = true // TODO: Dev

	// Public routes: /
	apiPublic := router.Group("/")
	apiPublic.Use(cors.New(corsConfig))
	{
		apiPublic.GET("/health", GetHealth)

		// GAEN - React
		apiPublic.POST("/publish", PostPublish)

		apiPublic.POST("/api/verify", PostVerify)

		apiPublic.POST("/api/certificate", PostCertificate)

		apiPublic.GET("/download", GetDownload)

		apiPublic.GET("/configuration", GetConfiguration)
	}

	return router.Run(os.Getenv("port"))
}
