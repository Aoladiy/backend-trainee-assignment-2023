package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type errorResponse struct {
	Message string `json:""`
}

func NewErrorResponse(c *gin.Context, statusCode int, message string) {
	logrus.Fatal(message)
	c.AbortWithStatusJSON(statusCode, errorResponse{message})
}
