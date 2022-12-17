package repository

import (
	"encoding/json"
	"fmt"
	"github.com/22hac07win/route-server.git/domain"
	"github.com/gin-gonic/gin"
)

func (s *supabaseDBClient) GetAllStory(c *gin.Context) ([]*domain.Story, error) {

	body, err := s.ReadAllContent(c, StoryTable)
	if err != nil {
		return nil, err
	}

	var stories []domain.Story
	err = json.Unmarshal(body, &stories)

	fmt.Println(stories)

	if err != nil {
		return nil, err
	}

	var res []*domain.Story
	for _, story := range stories {
		res = append(res, &story)
	}

	return res, nil
}
