package db

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
)

func (s *SupabaseDBClient) AddDBStore(c *gin.Context, data *DBStore) error {
	url := fmt.Sprintf("%s/rest/v1/store", s.Url)
	body, err := json.Marshal(data)
	if err != nil {
		return err
	}
	reader := bytes.NewReader(body)

	req, _ := http.NewRequest("POST", url, reader)

	req.Header.Add("ApiKey", s.ApiKey)
	req.Header.Add("Authorization", "Bearer "+s.ApiKey)
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Prefer", "return=minimal")

	client := new(http.Client)
	_, err = client.Do(req)

	return err
}

func (s *SupabaseDBClient) GetDBStore(c *gin.Context, userID string, key string) (*DBStore, error) {
	url := fmt.Sprintf("%s/rest/v1/store?user_id=eq.%s&key=eq.%s&order=id.desc&limit=1", s.Url, userID, key)
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Add("ApiKey", s.ApiKey)
	req.Header.Add("Authorization", "Bearer "+s.ApiKey)

	client := new(http.Client)
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	var res DBStore
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(body, &res)
	return &res, err
}
