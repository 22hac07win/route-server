package db

import (
	"github.com/22hac07win/route-server.git/domain"
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

type InsertDBUser struct {
	ID string `json:"id"`
}

type DBUser struct {
	ID        string `json:"id"`
	CreatedAt string `json:"created_at"`
}

type DBStory struct {
	ID            string `json:"id"`
	Name          string `json:"name"`
	FireIF        string `json:"fire_if"`
	FirestBlockID string `json:"firest_block_id"`
}

type DBBlock struct {
	ID        string           `json:"id"`
	StoryID   string           `json:"story_id"`
	BlockType domain.BlockType `json:"block_type"`
	Text      string           `json:"text"`
	Func      string           `json:"func"`
	Options   []domain.Option  `json:"options"`
	NextID    string           `json:"next_id"`
}

type StoreType string

const (
	InputStore  StoreType = "input"
	OptionStore StoreType = "option"
)

type InsertDBStore struct {
	StoreType StoreType `json:"store_type"`
	UserID    string    `json:"user_id"`
	Body      string    `json:"body"`
}

type DBStore struct {
	ID        string    `json:"id"`
	StoreType StoreType `json:"store_type"`
	UserID    string    `json:"user_id"`
	Body      string    `json:"body"`
}
