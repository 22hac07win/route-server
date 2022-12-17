package repository

import (
	"bytes"
	"fmt"
	"github.com/22hac07win/route-server.git/domain"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
	"os"
)

type SupabaseDBClient interface {
	ReadEqContent(c *gin.Context, table TableName, col string, value string) ([]byte, error)
	ReadMultiEqContent(c *gin.Context, table TableName, args []ReadMultiEqArg) ([]byte, error)
	ReadAllContent(c *gin.Context, table TableName) ([]byte, error)
	UpsertContent(c *gin.Context, table TableName, body []byte) error
	NewGetHttpRequest(url string) *http.Request
	NewUpsertPostHttpRequest(url string, reader *bytes.Reader) *http.Request
	GetNextBlock(c *gin.Context, nextId string) (domain.Block, error)
	GetTextBlock(c *gin.Context, id string) (*domain.TextBlock, error)
	GetFuncBlock(c *gin.Context, id string) (*domain.FunctionBlock, error)
	GetInputBlock(c *gin.Context, id string) (*domain.InputBlock, error)
	GetOptionBlock(c *gin.Context, id string) (*domain.OptionBlock, error)
	GetAllStory(c *gin.Context) ([]*domain.Story, error)
	UpsertUser(c *gin.Context, userID string, state string) error
	GetUser(c *gin.Context, userID string) (*domain.User, error)
	UpsertStore(c *gin.Context, data *UpsertStore) error
	GetStore(c *gin.Context, userID string, key string) (*domain.Store, error)
}

type supabaseDBClient struct {
	Url    string
	ApiKey string
}

func NewSupabaseDBClient() *supabaseDBClient {
	url := os.Getenv("SUPABASE_URL")
	ApiKey := os.Getenv("SUPABASE_API_KEY")

	return &supabaseDBClient{
		Url:    url,
		ApiKey: ApiKey,
	}
}

type UpsertUser struct {
	ID    string `json:"id"`
	State string `json:"state"`
}

type UpsertStore struct {
	UserID    string           `json:"user_id"`
	Key       string           `json:"key"`
	StoreType domain.StoreType `json:"store_type"`
	Body      string           `json:"body"`
}

func (s *supabaseDBClient) ReadEqContent(c *gin.Context, table TableName, col string, value string) ([]byte, error) {
	url := fmt.Sprintf("%s/rest/v1/%s?%s=eq.%s", s.Url, table, col, value)
	req := s.NewGetHttpRequest(url)

	client := new(http.Client)
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}

type ReadMultiEqArg struct {
	Col   string
	Value string
}

func (s *supabaseDBClient) ReadMultiEqContent(c *gin.Context, table TableName, args []ReadMultiEqArg) ([]byte, error) {

	query := ""
	for i, q := range args {
		if i != 0 {
			query += "&"
		}
		query += fmt.Sprintf("%s=eq.%s", q.Col, q.Value)
	}

	url := fmt.Sprintf("%s/rest/v1/%s?%s", s.Url, table, query)

	fmt.Println(url)

	req := s.NewGetHttpRequest(url)

	client := new(http.Client)
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	fmt.Println(string(body))

	return body, nil
}

func (s *supabaseDBClient) ReadAllContent(c *gin.Context, table TableName) ([]byte, error) {
	url := fmt.Sprintf("%s/rest/v1/%s?select=*", s.Url, table)
	fmt.Println(url)

	req := s.NewGetHttpRequest(url)

	client := new(http.Client)
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	fmt.Printf("%+v\n", resp)
	body, err := io.ReadAll(resp.Body)
	fmt.Println(string(body))
	if err != nil {
		return nil, err
	}
	return body, nil
}

func (s *supabaseDBClient) UpsertContent(c *gin.Context, table TableName, body []byte) error {

	url := fmt.Sprintf("%s/rest/v1/%s", s.Url, table)
	reader := bytes.NewReader(body)

	req := s.NewUpsertPostHttpRequest(url, reader)

	client := new(http.Client)
	_, err := client.Do(req)

	return err
}

func (s *supabaseDBClient) NewGetHttpRequest(url string) *http.Request {
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Add("ApiKey", s.ApiKey)
	req.Header.Add("Authorization", "Bearer "+s.ApiKey)

	return req
}

func (s *supabaseDBClient) NewUpsertPostHttpRequest(url string, reader *bytes.Reader) *http.Request {
	req, _ := http.NewRequest("POST", url, reader)

	req.Header.Add("ApiKey", s.ApiKey)
	req.Header.Add("Authorization", "Bearer "+s.ApiKey)
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Prefer", "resolution=merge-duplicates")

	return req
}
