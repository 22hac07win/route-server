package service

import (
	"fmt"
	"github.com/22hac07win/route-server.git/domain"
	"github.com/gin-gonic/gin"
	"strings"
)

func (rp *routeProvider) GetNextBlock(c *gin.Context, userID string, nextID string) (*domain.ApiResponse, error) {

	fmt.Println("nextID", nextID)
	arr := strings.Split(nextID, "-")

	index := len(arr) - 2
	fmt.Println(arr)
	bType := arr[index]

	switch bType {
	case "TEXT":
		tb, err := rp.s.GetTextBlock(c, nextID)
		if err != nil {
			return nil, err
		}
		return rp.GetTextBlockContent(c, *tb)
	case "FN":
		fb, err := rp.s.GetFuncBlock(c, nextID)
		if err != nil {
			return nil, err
		}
		return rp.GetFuncBlockContent(c, userID, fb)
	case "INPUT":
		ib, err := rp.s.GetInputBlock(c, nextID)
		if err != nil {
			return nil, err
		}
		return rp.GetInputBlockContent(c, ib)
	case "OPT":
		ob, err := rp.s.GetOptionBlock(c, nextID)
		if err != nil {
			return nil, err
		}
		return rp.GetOptionBlockContent(c, ob)
	}

	return nil, ErrInvalidID
}

func (rp *routeProvider) GetTextBlockContent(c *gin.Context, b domain.TextBlock) (*domain.ApiResponse, error) {
	res := &domain.ApiResponse{
		ID:        b.ID,
		BlockType: domain.TextBlockType,
		Text:      b.Text,
		NextID:    b.NextID,
	}
	return res, nil
}

func (rp *routeProvider) GetFuncBlockContent(c *gin.Context, userID string, b *domain.FunctionBlock) (*domain.ApiResponse, error) {

	text, err := rp.bfs.GenerateText(c, userID, b)
	if err != nil {
		return nil, err
	}

	res := &domain.ApiResponse{
		ID:        b.ID,
		BlockType: domain.FunctionBlockType,
		Text:      text,
		NextID:    b.NextID,
	}
	return res, nil
}

func (rp *routeProvider) GetInputBlockContent(c *gin.Context, b *domain.InputBlock) (*domain.ApiResponse, error) {
	fmt.Println(b)
	res := &domain.ApiResponse{
		ID:        b.ID,
		BlockType: domain.InputBlockType,
		Text:      b.Text,
		InputKey:  b.Key,
		NextID:    b.NextID,
	}
	return res, nil
}

func (rp *routeProvider) GetOptionBlockContent(c *gin.Context, b *domain.OptionBlock) (*domain.ApiResponse, error) {
	var opts []domain.ResOption
	for _, v := range b.Options {
		opt := domain.ResOption{
			OptionNumber: v.OptionNumber,
			OptionText:   v.OptionText,
			NextBlockID:  v.NextBlockID,
		}
		opts = append(opts, opt)
	}

	res := &domain.ApiResponse{
		ID:        b.ID,
		BlockType: domain.OptionBlockType,
		Text:      b.Text,
		Options:   opts,
	}
	return res, nil
}
