package db

import (
	"github.com/22hac07win/route-server.git/domain"
)

type DBUser struct {
	ID        string
	CreatedAt string
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

type DBStore struct {
	ID        string    `json:"id"`
	StoreType StoreType `json:"store_type"`
	UserID    string    `json:"user_id"`
	Body      string    `json:"body"`
}
