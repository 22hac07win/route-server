package service

import (
	"github.com/22hac07win/route-server.git/db"
	"github.com/22hac07win/route-server.git/domain"
	"github.com/gin-gonic/gin"
)

func GetNextBlockContent(c *gin.Context, nextID string) (*domain.ApiResponse, error) {
	sc := db.NewSupabaseDBClient()
	b, err := sc.GetNextBlock(c, nextID)
	if err != nil {
		return nil, err
	}

	content, err := b.GetContent()
	if err != nil {
		return nil, err
	}
	return content, nil
}

func GetNextStory(c *gin.Context, req domain.ApiRequest) (*domain.Story, error) {
	userID := c.GetString("userID")
	print(userID)
	// TODO: get next story
	return &domain.Story{}, nil
}
