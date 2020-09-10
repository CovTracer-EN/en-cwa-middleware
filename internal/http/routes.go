package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetHealth(c *gin.Context) {
	c.Status(http.StatusOK)
}

func PostPublish(c *gin.Context) {
	// TODO: Implement

	c.Status(http.StatusOK)
}

func PostVerify(c *gin.Context) {
	// TODO: Implement

	c.Status(http.StatusOK)
}

func PostCertificate(c *gin.Context) {
	// TODO: Implement

	c.Status(http.StatusOK)
}

func GetDownload(c *gin.Context) {
	// TODO: Implement

	c.Status(http.StatusOK)
}

func GetConfiguration(c *gin.Context) {
	// TODO: Implement

	c.Status(http.StatusOK)
}
