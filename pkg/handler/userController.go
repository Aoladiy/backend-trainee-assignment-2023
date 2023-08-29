package handler

import (
	backendTraineeAssignment2023 "github.com/Aoladiy/backend-trainee-assignment-2023"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
	"strconv"
)

func (h *Handler) getAllUsers(c *gin.Context) {

}
func (h *Handler) getUserById(c *gin.Context) {

}
func (h *Handler) getUserSegments(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	message, err := h.services.GetUserSegments(id)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"message": message,
	})
}
func (h *Handler) createUser(c *gin.Context) {
	var input backendTraineeAssignment2023.User

	if err := c.BindJSON(&input); err != nil {
		if err.Error() != "EOF" {
			NewErrorResponse(c, http.StatusBadRequest, err.Error())
			return
		}
	}

	id, err := h.services.CreateUser(input)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}
func (h *Handler) updateUserById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	var input struct {
		SegmentsToJoin  []string `json:"segmentsToJoin"`
		SegmentsToLeave []string `json:"segmentsToLeave"`
	}

	if err := c.BindJSON(&input); err != nil {
		if err != io.EOF {
			NewErrorResponse(c, http.StatusBadRequest, err.Error())
			return
		}
	}

	_, message, err := h.services.UpdateUserById(input.SegmentsToJoin, input.SegmentsToLeave, id)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"message": message,
	})
}

func (h *Handler) deleteUserById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	message, err := h.services.DeleteUser(id)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	if id, err = strconv.Atoi(message); err != nil {
		c.JSON(http.StatusOK, map[string]interface{}{
			"message": message,
		})
	} else {
		c.JSON(http.StatusOK, map[string]interface{}{
			"id": id,
		})
	}
}
