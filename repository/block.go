package repository

import (
	"encoding/json"
	"fmt"
	"github.com/22hac07win/route-server.git/domain"
	"github.com/gin-gonic/gin"
	"strings"
)

func (s *supabaseDBClient) GetNextBlock(c *gin.Context, nextID string) (domain.Block, error) {

	fmt.Println("nextID", nextID)
	arr := strings.Split(nextID, "-")

	index := len(arr) - 2
	fmt.Println(arr)
	bType := arr[index]

	switch bType {
	case "TEXT":
		return s.GetTextBlock(c, nextID)
	case "FN":
		return s.GetFuncBlock(c, nextID)
	case "INPUT":
		return s.GetInputBlock(c, nextID)
	case "OPT":
		return s.GetOptionBlock(c, nextID)
	}

	return nil, ErrInvalidID
}

func (s *supabaseDBClient) GetTextBlock(c *gin.Context, id string) (*domain.TextBlock, error) {
	byte, err := s.ReadEqContent(c, TextBlockTable, TextBlockTableColumns.ID, id)
	if err != nil {
		return nil, err
	}

	var res []domain.TextBlock
	err = json.Unmarshal(byte, &res)
	if len(res) == 0 {
		return nil, ErrNotExists
	}

	return &res[0], err
}

func (s *supabaseDBClient) GetFuncBlock(c *gin.Context, id string) (*domain.FunctionBlock, error) {
	byte, err := s.ReadEqContent(c, FuncBlockTable, FuncBlockTableColumns.ID, id)
	if err != nil {
		return nil, err
	}

	var res []domain.FunctionBlock
	err = json.Unmarshal(byte, &res)
	if len(res) == 0 {
		return nil, ErrNotExists
	}

	return &res[0], err
}

func (s *supabaseDBClient) GetInputBlock(c *gin.Context, id string) (*domain.InputBlock, error) {
	byte, err := s.ReadEqContent(c, InputBlockTable, InputBlockTableColumns.ID, id)
	if err != nil {
		return nil, err
	}

	var res []domain.InputBlock
	err = json.Unmarshal(byte, &res)
	if len(res) == 0 {
		return nil, ErrNotExists
	}

	return &res[0], err
}

func (s *supabaseDBClient) GetOptionBlock(c *gin.Context, id string) (*domain.OptionBlock, error) {
	byte, err := s.ReadEqContent(c, OptBlockTable, OptBlockTableColumns.ID, id)
	if err != nil {
		return nil, err
	}

	var res []domain.OptionBlock
	err = json.Unmarshal(byte, &res)
	if len(res) == 0 {
		return nil, ErrNotExists
	}

	return &res[0], err
}
