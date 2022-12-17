package domain

import (
	"time"
)

type State string

const (
	StartState = "start"
	LifeState  = "life"
)

type User struct {
	ID        string    `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	State     State     `json:"state"`
}

type FireIf string

const (
	StartIf  FireIf = "start"
	RandomIf FireIf = "random"
)

type Story struct {
	ID           string `json:"id"`
	Title        string `json:"title"`
	FireIf       FireIf `json:"fire_if"`
	FirstBlockID string `json:"first_block_id"`
}

type BlockType string

const (
	TextBlockType     BlockType = "text"
	FunctionBlockType BlockType = "function"
	InputBlockType    BlockType = "input"
	OptionBlockType   BlockType = "option"
)

type Block interface {
	GetContent() (*ApiResponse, error)
}

type TextBlock struct {
	ID      string `json:"id"`
	StoryID string `json:"story_id"`
	Text    string `json:"text"`
	NextID  string `json:"next_id"`
}

type FunctionBlock struct {
	ID       string   `json:"id"`
	StoryID  string   `json:"story_id"`
	Text     string   `json:"text"`
	Function string   `json:"function"`
	Args     []string `json:"args"`
	NextID   string   `json:"next_id"`
}

type InputBlock struct {
	ID      string `json:"id"`
	StoryID string `json:"story_id"`
	Text    string `json:"text"`
	Key     string `json:"key"`
	NextID  string `json:"next_id"`
}

type Option struct {
	OptionNumber int    `json:"option_number"`
	OptionText   string `json:"option_text"`
	NextBlockID  string `json:"next_block_id"`
}

type OptionBlock struct {
	ID      string   `json:"id"`
	StoryID string   `json:"story_id"`
	Text    string   `json:"text"`
	Options []Option `json:"options"`
}

type StoreType string

const (
	InputStore  StoreType = "input"
	OptionStore StoreType = "option"
)

type Store struct {
	UserID    string    `json:"user_id"`
	Key       string    `json:"key"`
	StoreType StoreType `json:"store_type"`
	Body      string    `json:"body"`
}
