package service

import (
	"github.com/22hac07win/route-server.git/db"
	"github.com/22hac07win/route-server.git/domain"
	"github.com/gin-gonic/gin"
)

func GetNextBlock(c *gin.Context, nextId string) (domain.Block, error) {
	sc := db.NewSupabaseDBClient()
	block, err := sc.GetBlock(c, nextId)
	if err != nil {
		return nil, err
	}

	return *block, nil
}

func GetNextStory(c *gin.Context, req domain.ApiRequest) (*domain.Story, error) {
	userID := c.GetString("userID")
	print(userID)
	// TODO: get next story
	return &domain.Story{}, nil
}
