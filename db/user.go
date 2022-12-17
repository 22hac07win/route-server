package db

import (
	"encoding/json"
	"github.com/22hac07win/route-server.git/domain"
	"github.com/gin-gonic/gin"
)

func (s *SupabaseDBClient) UpsertUser(c *gin.Context, userID string, state string) error {

	user := UpsertUser{
		ID:    userID,
		State: state,
	}

	body, err := json.Marshal(user)
	if err != nil {
		return err
	}

	err = s.UpsertContent(c, UserTable, body)
	return err
}

func (s *SupabaseDBClient) GetUser(c *gin.Context, userID string) (*domain.User, error) {

	body, err := s.ReadEqContent(c, UserTable, UserTableColumns.ID, userID)

	var res []domain.User
	err = json.Unmarshal(body, &res)

	return &res[0], err
}
