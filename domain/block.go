package domain

func (b *TextBlock) GetContent() (*ApiResponse, error) {
	res := &ApiResponse{
		ID:        b.ID,
		BlockType: TextBlockType,
		Text:      b.Text,
		InputKey:  false,
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
		InputKey:  false,
		NextID:    b.NextID,
	}
	return res, nil
}

func (b *InputBlock) GetContent() (*ApiResponse, error) {
	res := &ApiResponse{
		ID:        b.ID,
		BlockType: InputBlockType,
		Text:      b.Text,
		InputKey:  true,
		NextID:    b.NextID,
	}
	return res, nil
}

func (b *OptionBlock) GetContent() (*ApiResponse, error) {
	var opts []ResOption
	for _, v := range b.Options {
		opt := ResOption{
			OptionNumber: v.OptionNumber,
			OptionText:   v.OptionText,
			NextBlockID:  v.NextBlockID,
		}
		opts = append(opts, opt)
	}

	res := &ApiResponse{
		ID:        b.ID,
		BlockType: OptionBlockType,
		Text:      b.Text,
		InputKey:  false,
		Options:   opts,
	}
	return res, nil
}
