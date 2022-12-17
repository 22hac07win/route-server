package handler

import (
	"fmt"
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

	fmt.Println("PostMessage")

	if err := c.BindJSON(&req); err != nil {
		fmt.Println("BindJSON error")
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	var nextID string
	nextID = req.NextID
	if req.NextID == "" {
		userID := c.GetString("userID")
		story, err := rh.rp.GetNextStory(c, userID)
		if err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}
		nextID = story.FirstBlockID
		fmt.Println("nextId", nextID)
	}

	res, err := rh.rp.GetNextBlockContent(c, nextID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}
