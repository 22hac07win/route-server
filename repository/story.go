package repository

import (
	"encoding/json"
	"fmt"
	"github.com/22hac07win/route-server.git/domain"
	"github.com/gin-gonic/gin"
)

func (s *supabaseDBClient) GetAllStory(c *gin.Context) ([]domain.Story, error) {

	body, err := s.ReadAllContent(c, StoryTable)
	if err != nil {
		return nil, err
	}

	var str []domain.Story
	print(string(body))
	err = json.Unmarshal(body, &str)

	if err != nil {
		return nil, err
	}

	fmt.Println(str)
	var res []domain.Story
	for _, v := range str {
		res = append(res, v)
	}

	return res, nil
}
