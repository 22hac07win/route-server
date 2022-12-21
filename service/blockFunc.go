package service

import (
	"fmt"
	"github.com/22hac07win/route-server.git/domain"
	"github.com/22hac07win/route-server.git/repository"
	"github.com/gin-gonic/gin"
)

type BlockFuncService interface {
	GenerateText(c *gin.Context, userID string, fb *domain.FunctionBlock) (string, error)
	SetArgsFunc(c *gin.Context, userID string, fb *domain.FunctionBlock) (string, error)
	NullResFunc(c *gin.Context) (string, error)
}

type blockFuncService struct {
	s repository.SupabaseDBClient
}

func NewBlockFuncService(s repository.SupabaseDBClient) *blockFuncService {
	return &blockFuncService{s: s}
}

func (bfs *blockFuncService) GenerateText(c *gin.Context, userID string, fb *domain.FunctionBlock) (string, error) {
	switch fb.Function {
	case "SetArgsFunc":
		return bfs.SetArgsFunc(c, userID, fb)
	default:
		return bfs.NullResFunc(c)
	}
}

func (bfs *blockFuncService) SetArgsFunc(c *gin.Context, userID string, fb *domain.FunctionBlock) (string, error) {

	var strs []any
	for _, arg := range fb.Args {
		store, err := bfs.s.GetStore(c, userID, arg)
		if err != nil {
			return "", err
		}
		strs = append(strs, store.Body)
	}
	res := fmt.Sprintf(fb.Text, strs...)
	return res, nil
}

func (bfs *blockFuncService) NullResFunc(c *gin.Context) (string, error) {
	return "", ErrFunctionNotFound
}
