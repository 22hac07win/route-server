package domain

import (
	"github.com/22hac07win/route-server.git/router"
	"time"
)

type User struct {
	ID        string
	CreatedAt time.Time
}

type FireIf string

type Story struct {
	ID           string
	Name         string
	FireIf       FireIf
	FirstBlockID string
}

type BlockType string

const (
	TextBlockType     BlockType = "text"
	FunctionBlockType BlockType = "function"
	InputBlockType    BlockType = "input"
	OptionBlockType   BlockType = "option"
)

type Block interface {
	GetContent() (*router.ApiResponse, error)
}

type TextBlock struct {
	ID      string
	StoryID string
	Text    string
	NextID  string
}

type BlockFunc func() (string, error)

type FunctionBlock struct {
	ID       string
	StoryID  string
	Function BlockFunc
	NextID   string
}

type InputBlock struct {
	ID      string
	StoryID string
	Text    string
	NextID  string
}

type Option struct {
	OptionNumber int    `json:"option_number"`
	OptionText   string `json:"option_text"`
	NextBlockID  string `json:"next_block_id"`
}

type OptionBlock struct {
	ID      string
	StoryID string
	Text    string
	Options []Option
}
