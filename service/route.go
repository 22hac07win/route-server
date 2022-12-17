package service

import (
	"github.com/22hac07win/route-server.git/domain"
	"github.com/22hac07win/route-server.git/repository"
	"github.com/gin-gonic/gin"
	"math/rand"
	"time"
)

type RouteProvider interface {
	CreateStore(c *gin.Context, userID string, input domain.Input) error
	GetNextBlockContent(c *gin.Context, nextID string) (*domain.ApiResponse, error)
	GetNextStory(c *gin.Context, userID string) (*domain.Story, error)
}

type routeProvider struct {
	s repository.SupabaseDBClient
}

func NewRouteProvider(s repository.SupabaseDBClient) *routeProvider {
	return &routeProvider{s: s}
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

func (rp *routeProvider) GetNextBlockContent(c *gin.Context, nextID string) (*domain.ApiResponse, error) {

	b, err := rp.s.GetNextBlock(c, nextID)
	if err != nil {
		return nil, err
	}

	content, err := b.GetContent()
	if err != nil {
		return nil, err
	}
	return content, nil
}

func (rp *routeProvider) GetNextStory(c *gin.Context, userID string) (*domain.Story, error) {
	user, err := rp.s.GetUser(c, userID)
	if err != nil {
		return nil, err
	}

	stories, err := rp.s.GetAllStory(c)
	if err != nil {
		return nil, err
	}

	var res *domain.Story
	if user.State == domain.StartState {
		for _, story := range stories {
			if story.FireIf == domain.StartIf {
				res = story
				break
			}
		}
	} else if user.State == domain.LifeState {
		var s []*domain.Story
		for _, story := range stories {
			if story.FireIf == domain.RandomIf {
				s = append(s, story)
			}
		}

		rand.Seed(time.Now().UnixNano())
		res = s[rand.Intn(len(s))]
	}

	return res, nil
}
