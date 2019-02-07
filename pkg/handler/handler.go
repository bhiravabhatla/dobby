package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/thecasualcoder/dobby/pkg/config"
	"os"
	"strconv"
)

var statusCode = 200

// Health return the dobby health status
func Health(c *gin.Context) {
	c.JSON(statusCode, gin.H{"healthy": true})
}

// Version return dobby version
func Version(c *gin.Context) {
	envVersion := os.Getenv("VERSION")
	version := config.BuildVersion()
	if envVersion != "" {
		version = envVersion
	}
	c.JSON(200, gin.H{"version": version})
}

func init() {
	healthy, err := strconv.ParseBool(os.Getenv("HEALTH"))

	if err != nil {
		statusCode = 200
	} else if !healthy {
		statusCode = 500
	}
}
