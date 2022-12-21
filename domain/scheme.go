package domain

type Input struct {
	Key  string `json:"key"`
	Body string `json:"body"`
}

type ApiRequest struct {
	NextID     string `json:"nextId"`
	Input      Input  `json:"input"`
	OptionText string `json:"option"`
}

type ResOption struct {
	OptionNumber int    `json:"optionNumber"`
	OptionText   string `json:"optionText"`
	NextBlockID  string `json:"nextBlockId"`
}

type ApiResponse struct {
	ID        string      `json:"id"`
	BlockType BlockType   `json:"blockType"`
	Text      string      `json:"text"`
	InputKey  string      `json:"inputKey"`
	Options   []ResOption `json:"options"`
	NextID    string      `json:"nextId"`
}
