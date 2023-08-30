package handler

import (
	"github.com/Aoladiy/backend-trainee-assignment-2023/pkg/service"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()
	router.Group("/api")
	{
		segments := router.Group("/segments")
		{
			segments.GET("/", h.getAllSegments)
			segments.GET("/:slug", h.getSegmentBySlug)
			segments.POST("/", h.createSegment)
			segments.PUT("/:slug", h.updateSegmentBySlug)
			segments.DELETE("/:slug", h.deleteSegmentBySlug)
		}
		users := router.Group("users")
		{
			users.GET("/", h.getAllUsers)
			users.GET("/:id", h.getUserById)
			users.GET("/:id/segments", h.getUserSegments)
			users.GET("/:id/log", h.getUserLog)
			users.POST("/", h.createUser)
			users.PUT("/:id", h.updateUserById)
			users.DELETE("/:id", h.deleteUserById)
		}
	}
	return router
}
