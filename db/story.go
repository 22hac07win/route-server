package db

import (
	"encoding/json"
	"fmt"
	"github.com/22hac07win/route-server.git/domain"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
)

func (s *SupabaseDBClient) GetAllDBStory(c *gin.Context) ([]*domain.Story, error) {
	url := fmt.Sprintf("%s/rest/v1/story/select=*", s.Url)

	req := s.NewGetHttpRequest(url)

	client := new(http.Client)
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	var storys []domain.Story
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(body, &storys)

	if err != nil {
		return nil, err
	}

	var res []*domain.Story
	for _, story := range storys {
		res = append(res, &story)
	}

	return res, nil
}
