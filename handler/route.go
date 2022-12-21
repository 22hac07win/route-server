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

func (rh *routeHandler) GetInit(c *gin.Context) {
	userID := c.GetString("userID")

	err := rh.rp.UpdateUser(c, userID)

	fmt.Println("GetInit")

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	var req domain.ApiRequest
	rh.ReturnMessage(c, req)
}

func (rh *routeHandler) PostMessage(c *gin.Context) {
	var req domain.ApiRequest

	fmt.Println("PostMessage")

	if err := c.BindJSON(&req); err != nil {
		fmt.Println("BindJSON error")
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	rh.ReturnMessage(c, req)
}

func (rh *routeHandler) ReturnMessage(c *gin.Context, req domain.ApiRequest) {
	var nextID string
	nextID = req.NextID

	userID := c.GetString("userID")

	if req.Input.Key != "" {
		err := rh.rp.CreateStore(c, userID, req.Input)
		if err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}
	}

	if req.NextID == "" {

		story, err := rh.rp.GetNextStory(c, userID)
		if err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}
		nextID = story.FirstBlockID
		fmt.Println("nextId", nextID)
	}

	res, err := rh.rp.GetNextBlock(c, userID, nextID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}
