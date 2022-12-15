package db

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
)

func (s *SupabaseDBClient) GetAllDBStory(c *gin.Context) ([]*DBStory, error) {
	url := fmt.Sprintf("%s/rest/v1/story/select=*", s.Url)
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Add("apikey", s.apiKey)
	req.Header.Add("Authorization", "Bearer "+s.apiKey)

	client := new(http.Client)
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	var storys []DBStory
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	json.Unmarshal(body, &storys)

	var res []*DBStory
	for _, story := range storys {
		res = append(res, &story)
	}

	return res, nil
}