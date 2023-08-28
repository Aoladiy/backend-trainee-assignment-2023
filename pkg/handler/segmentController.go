package handler

import (
	backendTraineeAssignment2023 "github.com/Aoladiy/backend-trainee-assignment-2023"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) getAllSegments(c *gin.Context) {

}
func (h *Handler) getSegmentBySlug(c *gin.Context) {

}
func (h *Handler) createSegment(c *gin.Context) {
	var input backendTraineeAssignment2023.Segment

	if err := c.BindJSON(&input); err != nil {
		NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	slug, err := h.services.CreateSegment(input)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"slug": slug,
	})
}
func (h *Handler) updateSegmentBySlug(c *gin.Context) {

}
func (h *Handler) deleteSegmentBySlug(c *gin.Context) {
	slug := c.Param("slug")

	slug, err := h.services.DeleteSegment(slug)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"slug": slug,
	})
}
