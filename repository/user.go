package repository

import (
	"encoding/json"
	"fmt"
	"github.com/22hac07win/route-server.git/domain"
	"github.com/gin-gonic/gin"
)

func (s *supabaseDBClient) UpsertUser(c *gin.Context, userID string) error {

	user := UpsertUser{
		ID: userID,
	}

	body, err := json.Marshal(user)
	if err != nil {
		return err
	}

	fmt.Println(string(body))

	err = s.UpsertContent(c, UserTable, body)
	return err
}

func (s *supabaseDBClient) ChangeUserState(c *gin.Context, userID string, state domain.State) error {

	user := ChangeUserState{
		ID:    userID,
		State: string(state),
	}

	body, err := json.Marshal(user)
	if err != nil {
		return err
	}

	err = s.UpsertContent(c, UserTable, body)
	return err
}

func (s *supabaseDBClient) GetUser(c *gin.Context, userID string) (*domain.User, error) {

	fmt.Println(s.Url)

	body, err := s.ReadEqContent(c, UserTable, UserTableColumns.ID, userID)

	var res []domain.User
	err = json.Unmarshal(body, &res)

	if len(res) == 0 {
		return nil, ErrUnAuthorized
	}

	return &res[0], err
}
