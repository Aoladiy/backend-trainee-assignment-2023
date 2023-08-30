package handler

import (
	"encoding/csv"
	"fmt"
	backendTraineeAssignment2023 "github.com/Aoladiy/backend-trainee-assignment-2023"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
	"os"
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
func (h *Handler) getUserLog(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	var input struct {
		Period string `json:"period"`
	}
	if err := c.BindJSON(&input); err != nil {
		if err != io.EOF {
			c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		}
	}

	status, userLog, err := h.services.GetUserLog(id, input.Period)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	if !status {
		c.JSON(http.StatusOK, gin.H{"message": fmt.Sprintf("there's no user with id=%v", id)})
		return
	}

	tempFileName := fmt.Sprintf("user_log_%d.csv", id)
	tempFile, err := os.Create(tempFileName)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	defer tempFile.Close()

	csvWriter := csv.NewWriter(tempFile)
	defer csvWriter.Flush()

	for _, logEntry := range userLog {
		csvRow := []string{
			strconv.Itoa(logEntry.UserID),
			strconv.Itoa(logEntry.SegmentID),
			logEntry.Action,
			logEntry.Datetime.Format("2006-01-02 15:04:05"),
		}

		err := csvWriter.Write(csvRow)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			return
		}

		csvWriter.Flush()
		if err := csvWriter.Error(); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			return
		}
	}
	err = tempFile.Close()
	if err != nil {
		fmt.Println("Error closing temporary file:", err)
	}
	c.Header("Content-Description", "File Transfer")
	c.Header("Content-Disposition", fmt.Sprintf("attachment; filename=%s", tempFileName))
	c.Header("Content-Type", "application/octet-stream")
	c.File(tempFileName)
	err = os.Remove(tempFileName)
	if err != nil {
		fmt.Println("Error deleting temporary file:", err)
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
