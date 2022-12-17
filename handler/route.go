package handler

import (
	"net/http"

	"github.com/22hac07win/route-server.git/domain"
	"github.com/22hac07win/route-server.git/service"
	"github.com/gin-gonic/gin"
)

type routeHandler struct {
	rp service.RouteProvider
}

func NewRouteHandler(rp service.RouteProvider) *routeHandler {
	return &routeHandler{rp: rp}
}

func (rh *routeHandler) PostMessage(c *gin.Context) {
	var req domain.ApiRequest
	if err := c.BindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	var nextId string
	if req.NextID == "" {
		userID := c.GetString("userID")
		story, err := rh.rp.GetNextStory(c, userID)
		if err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}
		nextId = story.FirstBlockID
	}
	nextId = req.NextID

	res, err := rh.rp.GetNextBlockContent(c, nextId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}
