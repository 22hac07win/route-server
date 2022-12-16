package db

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/22hac07win/route-server.git/domain"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
)

func (s *SupabaseDBClient) AddDBStore(c *gin.Context, data *domain.Store) error {
	url := fmt.Sprintf("%s/rest/v1/store", s.Url)
	body, err := json.Marshal(data)
	if err != nil {
		return err
	}
	reader := bytes.NewReader(body)

	req := s.NewPostHttpRequest(url, reader)

	client := new(http.Client)
	_, err = client.Do(req)

	return err
}

func (s *SupabaseDBClient) GetDBStore(c *gin.Context, userID string, key string) (*domain.Store, error) {
	url := fmt.Sprintf("%s/rest/v1/store?user_id=eq.%s&key=eq.%s&order=id.desc&limit=1", s.Url, userID, key)
	req := s.NewGetHttpRequest(url)

	client := new(http.Client)
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	var res domain.Store
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(body, &res)
	return &res, err
}
