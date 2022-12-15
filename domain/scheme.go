package domain

type ApiRequest struct {
	NextID string `json:"next_id"`
	Input  string `json:"input"`
	Option string `json:"option"`
}

type ResOption struct {
	OptionalNumber int    `json:"optional_number"`
	OptionalText   string `json:"optional_text"`
	NextBlockID    string `json:"next_block_id"`
}

type ApiResponse struct {
	ID        string      `json:"id"`
	BlockType BlockType   `json:"block_type"`
	Text      string      `json:"text"`
	IsInput   bool        `json:"is_input"`
	Options   []ResOption `json:"options"`
	NextID    string      `json:"next_id"`
}
