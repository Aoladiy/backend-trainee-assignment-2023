package handler

import "github.com/gin-gonic/gin"

type Handler struct {
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
			users.POST("/", h.createUser)
			users.PUT("/:id", h.updateUserById)
			users.DELETE("/:id", h.deleteUserById)
		}
	}
	return router
}
