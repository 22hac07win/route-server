package db

import (
	"encoding/json"
	"fmt"
	"github.com/22hac07win/route-server.git/domain"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
	"os"
)

type SupabaseDBClient struct {
	Url    string
	apiKey string
}

func NewSupabaseDBClient() *SupabaseDBClient {
	url := os.Getenv("SUPABASE_URL")
	apiKey := os.Getenv("SUPABASE_API_KEY")

	return &SupabaseDBClient{
		Url:    url,
		apiKey: apiKey,
	}
}

func (s *SupabaseDBClient) GetBlock(c *gin.Context, blocId string) (*domain.Block, error) {
	b, err := s.GetDBBlock(c, blocId)
	if err != nil {
		return nil, err
	}

	var res domain.Block

	switch b.BlockType {
	case domain.TextBlockType:
		res = &domain.TextBlock{
			ID:      b.ID,
			StoryID: b.StoryID,
			Text:    b.Text,
			NextID:  b.NextID,
		}
	case domain.FunctionBlockType:
		fn, err := b.GetFunction(b.Func)
		if err != nil {
			return nil, err
		}

		res = &domain.FunctionBlock{
			ID:       b.ID,
			StoryID:  b.StoryID,
			Function: fn,
			NextID:   b.NextID,
		}

	case domain.OptionBlockType:
		res = &domain.OptionBlock{
			ID:      b.ID,
			StoryID: b.StoryID,
			Text:    b.Text,
			Options: b.Options,
		}
	}

	return &res, nil
}

func (s *SupabaseDBClient) GetDBBlock(c *gin.Context, blockId string) (*DBBlock, error) {
	url := fmt.Sprintf("%s/rest/v1/blocks/id=%s", s.Url, blockId)
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Add("apikey", s.apiKey)
	req.Header.Add("Authorization", "Bearer "+s.apiKey)

	client := new(http.Client)
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	var res DBBlock
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	json.Unmarshal(body, &res)
	return &res, nil
}
