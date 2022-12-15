package service

import (
	"github.com/22hac07win/route-server.git/db"
	"github.com/22hac07win/route-server.git/domain"
	"github.com/gin-gonic/gin"
)

func GetNextBlock(c *gin.Context, nextId string) (*domain.Block, error) {
	sc := db.NewSupabaseDBClient()
	block, err := sc.GetBlock(c, nextId)
	if err != nil {
		return nil, err
	}

	return block, nil
}

func GetNextStory(c *gin.Context) (*domain.Story, error) {
	return &domain.Story{}, nil
}
