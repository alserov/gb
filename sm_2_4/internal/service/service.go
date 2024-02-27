package service

import (
	"context"
	"fmt"
	"github.com/alserov/gb/sm_2_4/internal/db"
	"github.com/alserov/gb/sm_2_4/internal/log"
	"github.com/alserov/gb/sm_2_4/internal/models"
	"github.com/google/uuid"
)

type Service interface {
	Friendship
	User
}

type Friendship interface {
	SuggestFriendship(ctx context.Context, receiverID models.ID, initiatorID models.ID) error
	AnswerOnFriendshipSuggestion(ctx context.Context, req models.AnswerOnFriendshipSuggestionReq) error
	GetAllFriendshipSuggestions(ctx context.Context, userID models.ID) (models.Users, error)
	GetAllFriends(ctx context.Context, userID models.ID) (models.Users, error)
}

type User interface {
	CreateUser(ctx context.Context, req models.CreateUserReq) (models.ID, error)
	DeleteUser(ctx context.Context, userID models.ID) error
}

func NewService(log log.Logger, repo db.Repository) Service {
	return &service{
		log:  log,
		repo: repo,
	}
}

type service struct {
	log  log.Logger
	repo db.Repository
}

func (s service) SuggestFriendship(ctx context.Context, receiverID models.ID, initiatorID models.ID) error {
	if err := s.repo.SuggestFriendship(ctx, receiverID, initiatorID); err != nil {
		return fmt.Errorf("repo error: %w", err)
	}

	return nil
}

func (s service) AnswerOnFriendshipSuggestion(ctx context.Context, req models.AnswerOnFriendshipSuggestionReq) error {
	if err := s.repo.AnswerOnFriendshipSuggestion(ctx, req); err != nil {
		return fmt.Errorf("repo error: %w", err)
	}

	return nil
}

func (s service) GetAllFriendshipSuggestions(ctx context.Context, userID models.ID) (models.Users, error) {
	suggs, err := s.repo.GetAllFriendshipSuggestions(ctx, userID)
	if err != nil {
		return nil, fmt.Errorf("repo error: %w", err)
	}

	return suggs, nil
}

func (s service) GetAllFriends(ctx context.Context, userID models.ID) (models.Users, error) {
	frnds, err := s.repo.GetAllFriends(ctx, userID)
	if err != nil {
		return nil, fmt.Errorf("repo error: %w", err)
	}

	return frnds, nil
}

func (s service) CreateUser(ctx context.Context, req models.CreateUserReq) (models.ID, error) {
	req.ID = models.ID(uuid.New().String())

	err := s.repo.CreteUser(ctx, req)
	if err != nil {
		return "", fmt.Errorf("repo error: %w", err)
	}

	return req.ID, nil
}

func (s service) DeleteUser(ctx context.Context, userID models.ID) error {
	err := s.repo.DeleteUser(ctx, userID)
	if err != nil {
		return fmt.Errorf("repo error: %w", err)
	}

	return nil
}
