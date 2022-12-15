package router

import (
	"github.com/22hac07win/route-server.git/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

func PostMessage(c *gin.Context) {
	var req ApiRequest
	if err := c.BindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	var nextId string
	if req.NextID == "" {
		return
	}
	nextId = req.NextID

	block, err := service.GetNextBlock(c, nextId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	res, err := block.GetContent()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}
