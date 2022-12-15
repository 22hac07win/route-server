package db

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
)

func (s *SupabaseDBClient) InsertDBUser(c *gin.Context, userID string) error {
	url := fmt.Sprintf("%s/rest/v1/user", s.Url)

	user := InsertDBUser{
		ID: userID,
	}

	body, err := json.Marshal(user)
	if err != nil {
		return err
	}

	reader := bytes.NewReader(body)

	req, _ := http.NewRequest("POST", url, reader)

	req.Header.Add("apikey", s.apiKey)
	req.Header.Add("Authorization", "Bearer "+s.apiKey)
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Prefer", "return=minimal")

	client := new(http.Client)
	_, err = client.Do(req)

	return err
}

func (s *SupabaseDBClient) GetDBUser(c *gin.Context, userID string) (*DBUser, error) {
	url := fmt.Sprintf("%s/rest/v1/user?id=%s", s.Url, userID)
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Add("apikey", s.apiKey)
	req.Header.Add("Authorization", "Bearer "+s.apiKey)

	client := new(http.Client)
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	var res DBUser
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	json.Unmarshal(body, &res)
	return &res, nil
}
