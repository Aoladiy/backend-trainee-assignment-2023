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
	var firstError error = nil

	if firstError = c.BindJSON(&input); firstError != nil {
		if firstError.Error() == "EOF" {
			c.JSON(http.StatusOK, map[string]interface{}{
				"message": "Request must not be empty",
			})
		} else {
			NewErrorResponse(c, http.StatusBadRequest, firstError.Error())
			return
		}
	}
	if firstError == nil {
		if input.Slug == "" {
			c.JSON(http.StatusOK, map[string]interface{}{
				"message": "slug must not be empty",
			})
		} else {
			slug, err := h.services.CreateSegment(input)
			if err != nil {
				NewErrorResponse(c, http.StatusInternalServerError, err.Error())
				return
			}
			c.JSON(http.StatusOK, map[string]interface{}{
				"slug": slug,
			})
		}
	}
}
func (h *Handler) updateSegmentBySlug(c *gin.Context) {

}
func (h *Handler) deleteSegmentBySlug(c *gin.Context) {
	slug := c.Param("slug")

	status, message, err := h.services.DeleteSegment(slug)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	if status {
		c.JSON(http.StatusOK, map[string]interface{}{
			"slug": message,
		})
	} else {
		c.JSON(http.StatusOK, map[string]interface{}{
			"message": message,
		})
	}
}
