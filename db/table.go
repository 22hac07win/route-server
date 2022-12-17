package db

type TableName string

const (
	UserTable       TableName = "user"
	StoryTable      TableName = "story"
	TextBlockTable  TableName = "text_block"
	FuncBlockTable  TableName = "function_block"
	InputBlockTable TableName = "input_block"
	OptBlockTable   TableName = "option_block"
	StoreTable      TableName = "store"
)

var UserTableColumns = struct {
	ID        string
	CreatedAt string
	State     string
}{
	ID:        "id",
	CreatedAt: "created_at",
	State:     "state",
}

var StoryTableColumns = struct {
	ID           string
	Title        string
	FireIF       string
	FirstBlockID string
}{
	ID:           "id",
	Title:        "title",
	FireIF:       "fire_if",
	FirstBlockID: "first_block_id",
}

var TextBlockTableColumns = struct {
	ID      string
	StoryID string
	Text    string
	NextID  string
}{
	ID:      "id",
	StoryID: "story_id",
	Text:    "text",
	NextID:  "next_id",
}
var FuncBlockTableColumns = struct {
	ID       string
	StoryID  string
	Text     string
	Function string
	Args     string
	NextID   string
}{
	ID:       "id",
	StoryID:  "story_id",
	Text:     "text",
	Function: "function",
	Args:     "args",
	NextID:   "next_id",
}

var InputBlockTableColumns = struct {
	ID      string
	StoryID string
	Text    string
	Key     string
	NextID  string
}{
	ID:      "id",
	StoryID: "story_id",
	Text:    "text",
	Key:     "key",
	NextID:  "next_id",
}

var OptBlockTableColumns = struct {
	ID      string
	StoryID string
	Text    string
	Options string
}{
	ID:      "id",
	StoryID: "story_id",
	Text:    "text",
	Options: "options",
}

var StoreTableColumns = struct {
	UserID    string
	Key       string
	StoreType string
	Body      string
}{
	UserID:    "user_id",
	Key:       "key",
	StoreType: "store_type",
	Body:      "body",
}
