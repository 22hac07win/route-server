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
	req.Header.Add("ApiKey", s.ApiKey)
	req.Header.Add("Authorization", "Bearer "+s.ApiKey)

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

	err = json.Unmarshal(body, &storys)

	if err != nil {
		return nil, err
	}

	var res []*DBStory
	for _, story := range storys {
		res = append(res, &story)
	}

	return res, nil
}
