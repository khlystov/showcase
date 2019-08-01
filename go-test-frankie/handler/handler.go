package handler

import (
	"go-test-frankie/request"
	"net/http"

	"github.com/gin-gonic/gin"
)

// IsGood - Http handler for POST /isgood
func IsGood(c *gin.Context) {
	var r request.DeviceCheck

	if err := c.ShouldBindJSON(&r.Item); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    0,
			"message": "Everything is wrong. Go fix it.",
		})
		return
	}

	ok, error := r.Validate()

	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    0,
			"message": error.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"puppy": true,
	})
}
