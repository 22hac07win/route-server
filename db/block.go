package db

import (
	"encoding/json"
	"github.com/22hac07win/route-server.git/domain"
	"github.com/gin-gonic/gin"
	"strings"
)

func (s *SupabaseDBClient) GetNextBlock(c *gin.Context, nextId string) (domain.Block, error) {

	slice := strings.Split(nextId, "-")

	index := len(slice) - 2
	bType := slice[index]

	switch bType {
	case "TEXT":
		return s.GetTextBlock(c, nextId)
	case "FN":
		return s.GetFuncBlock(c, nextId)
	case "INPUT":
		return s.GetInputBlock(c, nextId)
	case "OPT":
		return s.GetOptionBlock(c, nextId)
	}

	return nil, ErrInvalidID
}

func (s *SupabaseDBClient) GetTextBlock(c *gin.Context, id string) (*domain.TextBlock, error) {
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

func (s *SupabaseDBClient) GetFuncBlock(c *gin.Context, id string) (*domain.FunctionBlock, error) {
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

func (s *SupabaseDBClient) GetInputBlock(c *gin.Context, id string) (*domain.InputBlock, error) {
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

func (s *SupabaseDBClient) GetOptionBlock(c *gin.Context, id string) (*domain.OptionBlock, error) {
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
