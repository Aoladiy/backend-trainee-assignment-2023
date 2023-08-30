package handler

import (
	"fmt"
	backendTraineeAssignment2023 "github.com/Aoladiy/backend-trainee-assignment-2023"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func (h *Handler) getAllSegments(c *gin.Context) {
	segments, err := h.services.GetAllSegments()
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"message": segments,
	})
}
func (h *Handler) getSegmentBySlug(c *gin.Context) {
	slug := c.Param("slug")
	status, segment, err := h.services.GetSegmentBySlug(slug)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
	}
	if status {
		c.JSON(http.StatusOK, map[string]interface{}{
			"message": segment,
		})
	} else {
		c.JSON(http.StatusOK, map[string]interface{}{
			"message": fmt.Sprintf("there's no segment with slug=%v", slug),
		})
	}
}
func (h *Handler) createSegment(c *gin.Context) {
	var input struct {
		Id                   int
		Slug                 string
		AutoAssignPercentage string
	}
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
	if input.AutoAssignPercentage == "" {
		segment := backendTraineeAssignment2023.Segment{
			Id:   input.Id,
			Slug: input.Slug,
		}
		if firstError == nil {
			if input.Slug == "" {
				c.JSON(http.StatusOK, map[string]interface{}{
					"message": "slug must not be empty",
				})
			} else {
				slug, err := h.services.CreateSegment(segment, 0)
				if err != nil {
					NewErrorResponse(c, http.StatusInternalServerError, err.Error())
					return
				}
				c.JSON(http.StatusOK, map[string]interface{}{
					"message": slug,
				})
			}
		}
	} else {
		if autoAssignPercentage, err := strconv.Atoi(input.AutoAssignPercentage); err != nil {
			c.JSON(http.StatusOK, map[string]interface{}{
				"message": "autoAssignPercentage must be integer",
			})
		} else {
			segment := backendTraineeAssignment2023.Segment{
				Id:   input.Id,
				Slug: input.Slug,
			}
			if firstError == nil {
				if input.Slug == "" {
					c.JSON(http.StatusOK, map[string]interface{}{
						"message": "slug must not be empty",
					})
				} else {
					slug, err := h.services.CreateSegment(segment, autoAssignPercentage)
					if err != nil {
						NewErrorResponse(c, http.StatusInternalServerError, err.Error())
						return
					}
					c.JSON(http.StatusOK, map[string]interface{}{
						"message": slug,
					})
				}
			}
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
