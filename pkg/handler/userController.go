package handler

import (
	"fmt"
	backendTraineeAssignment2023 "github.com/Aoladiy/backend-trainee-assignment-2023"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
	"strconv"
	"time"
)

func (h *Handler) getAllUsers(c *gin.Context) {
	users, err := h.services.GetAllUsers()
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"message": users,
	})
}
func (h *Handler) getUserById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	status, user, err := h.services.GetUserById(id)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
	}
	if status {
		c.JSON(http.StatusOK, map[string]interface{}{
			"message": user,
		})
	} else {
		c.JSON(http.StatusOK, map[string]interface{}{
			"message": fmt.Sprintf("there's no user with id=%v", id),
		})
	}
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
		TimeToLeave     string   `json:"timeToLeave"`
	}

	if err := c.BindJSON(&input); err != nil {
		if err != io.EOF {
			NewErrorResponse(c, http.StatusBadRequest, err.Error())
			return
		}
	}

	if input.TimeToLeave == "" {
		_, message, err := h.services.UpdateUserById(input.SegmentsToJoin, input.SegmentsToLeave, id, -time.Second)
		if err != nil {
			NewErrorResponse(c, http.StatusInternalServerError, err.Error())
			return
		}

		c.JSON(http.StatusOK, map[string]interface{}{
			"message": message,
		})
	} else {
		timeToLeave, err := time.ParseDuration(input.TimeToLeave)
		if err != nil {
			c.JSON(http.StatusOK, map[string]interface{}{
				"message": err.Error(),
			})
		} else {

			_, message, err := h.services.UpdateUserById(input.SegmentsToJoin, input.SegmentsToLeave, id, timeToLeave)
			if err != nil {
				NewErrorResponse(c, http.StatusInternalServerError, err.Error())
				return
			}

			c.JSON(http.StatusOK, map[string]interface{}{
				"message": message,
			})
		}
	}
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
