package service

import (
	"fmt"
	"github.com/22hac07win/route-server.git/domain"
	"github.com/22hac07win/route-server.git/repository"
	"github.com/gin-gonic/gin"
	"math/rand"
	"time"
)

type RouteProvider interface {
	CreateStore(c *gin.Context, userID string, input domain.Input) error
	UpdateUser(c *gin.Context, userID string) error
	GetNextStory(c *gin.Context, userID string) (*domain.Story, error)
	GetNextBlock(c *gin.Context, userID string, nextID string) (*domain.ApiResponse, error)
	GetTextBlockContent(c *gin.Context, b domain.TextBlock) (*domain.ApiResponse, error)
	GetFuncBlockContent(c *gin.Context, userID string, b *domain.FunctionBlock) (*domain.ApiResponse, error)
	GetInputBlockContent(c *gin.Context, b *domain.InputBlock) (*domain.ApiResponse, error)
	GetOptionBlockContent(c *gin.Context, b *domain.OptionBlock) (*domain.ApiResponse, error)
}

type routeProvider struct {
	s   repository.SupabaseDBClient
	bfs BlockFuncService
}

func NewRouteProvider(s repository.SupabaseDBClient, bfs BlockFuncService) *routeProvider {
	return &routeProvider{s: s, bfs: bfs}
}

func (rp *routeProvider) CreateStore(c *gin.Context, userID string, input domain.Input) error {
	data := &repository.UpsertStore{
		UserID:    userID,
		Key:       input.Key,
		StoreType: domain.InputStore,
		Body:      input.Body,
	}
	err := rp.s.UpsertStore(c, data)
	return err
}

func (rp *routeProvider) UpdateUser(c *gin.Context, userID string) error {
	err := rp.s.UpsertUser(c, userID)
	return err
}

func (rp *routeProvider) GetNextStory(c *gin.Context, userID string) (*domain.Story, error) {

	user, err := rp.s.GetUser(c, userID)
	if err != nil {
		return nil, err
	}

	fmt.Println("pass")

	str, err := rp.s.GetAllStory(c)
	if err != nil {
		return nil, err
	}

	fmt.Println(str)
	var res *domain.Story
	if user.State == domain.StartState {
		for _, v := range str {
			if v.FireIf == "start" {
				res = &v
				break
			}
		}
	} else if user.State == domain.LifeState {
		var s []domain.Story
		for _, v := range str {
			if v.FireIf == "random" {
				s = append(s, v)
			}
		}

		rand.Seed(time.Now().UnixNano())
		res = &s[rand.Intn(len(s))]
	}

	if res == nil {
		return nil, ErrNotFoundStory
	}

	return res, nil
}
