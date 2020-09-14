package http

import (
	"log"
	"net/http"

	"github.com/rise-center/en-cwa-middleware/internal/config"
	"github.com/rise-center/en-cwa-middleware/internal/utils"
	"github.com/rise-center/en-cwa-middleware/pkg/types"

	"github.com/gin-gonic/gin"
	"github.com/google/exposure-notifications-server/pkg/api/v1"
)

// Server health check
func GetHealth(c *gin.Context) {
	c.Status(http.StatusOK)
}

// Publish exposure diagnostic keys
func PostPublish(c *gin.Context) {
	var body v1.Publish
	if err := c.BindJSON(&body); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	// TODO: Convert to CWA payload
	// TODO: Edit pkg/types CWAPublishBody to match CWA requirements
	cwaBody := types.CWAPublishBody{
		Keys:                 nil,
		HealthAuthorityID:    "",
		VerificationPayload:  "",
		HMACKey:              "",
		SymptomOnsetInterval: 0,
		Traveler:             false,
		RevisionToken:        "",
		Padding:              "",
	}

	// TODO: Forward to CWA server
	// Add headers if you require something beyond content-type:application/json
	cwaHeaders := make(map[string]string)

	res, err := utils.HttpPost(config.CWAPublishUrl, cwaBody, cwaHeaders)
	if err != nil {
		log.Println(err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	if res.StatusCode() != 200 {
		log.Println(res.Status())
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	// TODO: Populate app response
	appRes := v1.PublishResponse{
		RevisionToken:     "",
		InsertedExposures: 0,
		ErrorMessage:      "",
		Code:              "",
		Padding:           "",
	}

	c.JSON(http.StatusOK, appRes)
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
