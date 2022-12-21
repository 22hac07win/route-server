package repository

import (
	"encoding/json"
	"fmt"
	"github.com/22hac07win/route-server.git/domain"
	"github.com/gin-gonic/gin"
)

func (s *supabaseDBClient) UpsertStore(c *gin.Context, data *UpsertStore) error {
	fmt.Println("store:", data)
	body, err := json.Marshal(data)
	if err != nil {
		return err
	}

	err = s.UpsertContent(c, StoreTable, body)
	return err
}

func (s *supabaseDBClient) GetStore(c *gin.Context, userID string, key string) (*domain.Store, error) {

	args := []ReadMultiEqArg{
		{Col: StoreTableColumns.UserID, Value: userID},
		{Col: StoreTableColumns.Key, Value: key},
	}

	body, err := s.ReadMultiEqContent(c, StoreTable, args)

	var res []domain.Store
	err = json.Unmarshal(body, &res)
	return &res[0], err
}
