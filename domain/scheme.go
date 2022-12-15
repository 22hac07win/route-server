package domain

type ApiRequest struct {
	NextID     string `json:"nextId"`
	Input      string `json:"input"`
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
	HaveInput bool        `json:"haveInput"`
	Options   []ResOption `json:"options"`
	NextID    string      `json:"nextId"`
}
