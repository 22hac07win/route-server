package db

import (
	"bytes"
	"fmt"
	"github.com/22hac07win/route-server.git/domain"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
	"os"
)

type SupabaseDBClient struct {
	Url    string
	ApiKey string
}

func NewSupabaseDBClient() *SupabaseDBClient {
	url := os.Getenv("SUPABASE_URL")
	ApiKey := os.Getenv("SUPABASE_API_KEY")

	return &SupabaseDBClient{
		Url:    url,
		ApiKey: ApiKey,
	}
}

type InsertUser struct {
	ID string `json:"id"`
}

type InsertStore struct {
	StoreType domain.StoreType `json:"store_type"`
	UserID    string           `json:"user_id"`
	Body      string           `json:"body"`
}

func (s *SupabaseDBClient) ReadEqContent(c *gin.Context, table TableName, col string, value string) ([]byte, error) {
	url := fmt.Sprintf("%s/rest/v1/%s?%s=eq.%s", s.Url, table, col, value)
	req := s.NewGetHttpRequest(url)

	client := new(http.Client)
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}

func (s *SupabaseDBClient) NewGetHttpRequest(url string) *http.Request {
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Add("ApiKey", s.ApiKey)
	req.Header.Add("Authorization", "Bearer "+s.ApiKey)

	return req
}

func (s *SupabaseDBClient) NewPostHttpRequest(url string, reader *bytes.Reader) *http.Request {
	req, _ := http.NewRequest("POST", url, reader)

	req.Header.Add("ApiKey", s.ApiKey)
	req.Header.Add("Authorization", "Bearer "+s.ApiKey)
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Prefer", "return=minimal")

	return req
}
