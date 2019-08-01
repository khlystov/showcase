package router

import (
	"go-test-frankie/handler"

	"github.com/gin-gonic/gin"
)

// Setup used to set up gin router
func Setup() *gin.Engine {
	r := gin.Default()
	r.POST("/isgood", handler.IsGood)

	return r
}
