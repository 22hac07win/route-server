package domain

func (b *TextBlock) GetContent() (*ApiResponse, error) {
	res := &ApiResponse{
		ID:        b.ID,
		BlockType: TextBlockType,
		Text:      b.Text,
		IsInput:   false,
		NextID:    b.NextID,
	}
	return res, nil
}

func (b *FunctionBlock) GetContent() (*ApiResponse, error) {

	f := b.Function
	text, err := f()

	if err != nil {
		return nil, err
	}

	res := &ApiResponse{
		ID:        b.ID,
		BlockType: FunctionBlockType,
		Text:      text,
		IsInput:   false,
		NextID:    b.NextID,
	}
	return res, nil
}

func (b *InputBlock) GetContent() (*ApiResponse, error) {
	res := &ApiResponse{
		ID:        b.ID,
		BlockType: InputBlockType,
		Text:      b.Text,
		IsInput:   true,
		NextID:    b.NextID,
	}
	return res, nil
}

func (b *OptionBlock) GetContent() (*ApiResponse, error) {
	var opts []ResOption
	for _, v := range b.Options {
		opt := ResOption{
			OptionalNumber: v.OptionNumber,
			OptionalText:   v.OptionText,
		}
		opts = append(opts, opt)
	}

	res := &ApiResponse{
		ID:        b.ID,
		BlockType: OptionBlockType,
		Text:      b.Text,
		IsInput:   false,
		Options:   opts,
	}
	return res, nil
}
