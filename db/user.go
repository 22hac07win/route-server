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

func (s *SupabaseDBClient) InsertDBUser(c *gin.Context, userID string) error {
	url := fmt.Sprintf("%s/rest/v1/user", s.Url)

	user := InsertUser{
		ID: userID,
	}

	body, err := json.Marshal(user)
	if err != nil {
		return err
	}

	reader := bytes.NewReader(body)

	req := s.NewPostHttpRequest(url, reader)

	client := new(http.Client)
	_, err = client.Do(req)

	return err
}

func (s *SupabaseDBClient) GetDBUser(c *gin.Context, userID string) (*domain.User, error) {
	url := fmt.Sprintf("%s/rest/v1/user?id=eq.%s", s.Url, userID)

	req := s.NewGetHttpRequest(url)

	client := new(http.Client)
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	var res []domain.User
	body, err := io.ReadAll(resp.Body)

	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(body, &res)

	return &res[0], err
}
