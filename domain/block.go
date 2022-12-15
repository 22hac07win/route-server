package domain

import (
	"github.com/22hac07win/route-server.git/router"
)

func (b *TextBlock) GetContent() (*router.ApiResponse, error) {
	res := &router.ApiResponse{
		ID:        b.ID,
		BlockType: TextBlockType,
		Text:      b.Text,
		IsInput:   false,
		NextID:    b.NextID,
	}
	return res, nil
}

func (b *FunctionBlock) GetContent() (*router.ApiResponse, error) {

	f := b.Function
	text, err := f()

	if err != nil {
		return nil, err
	}

	res := &router.ApiResponse{
		ID:        b.ID,
		BlockType: FunctionBlockType,
		Text:      text,
		IsInput:   false,
		NextID:    b.NextID,
	}
	return res, nil
}

func (b *InputBlock) GetContent() (*router.ApiResponse, error) {
	res := &router.ApiResponse{
		ID:        b.ID,
		BlockType: InputBlockType,
		Text:      b.Text,
		IsInput:   true,
		NextID:    b.NextID,
	}
	return res, nil
}

func (b *OptionBlock) GetContent() (*router.ApiResponse, error) {
	var opts []router.ResOption
	for _, v := range b.Options {
		opt := router.ResOption{
			OptionalNumber: v.OptionNumber,
			OptionalText:   v.OptionText,
		}
		opts = append(opts, opt)
	}

	res := &router.ApiResponse{
		ID:        b.ID,
		BlockType: OptionBlockType,
		Text:      b.Text,
		IsInput:   false,
		Options:   opts,
	}
	return res, nil
}
