package repository

import (
	"encoding/json"
	"github.com/22hac07win/route-server.git/domain"
	"github.com/gin-gonic/gin"
)

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
