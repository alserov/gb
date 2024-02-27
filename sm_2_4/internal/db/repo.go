package db

import (
	"context"
	"github.com/alserov/gb/sm_2_4/internal/models"
)

type Repository interface {
	CreteUser(ctx context.Context, req models.CreateUserReq) error
	DeleteUser(ctx context.Context, userID models.ID) error

	SuggestFriendship(ctx context.Context, receiverID models.ID, initiatorID models.ID) error
	AnswerOnFriendshipSuggestion(ctx context.Context, req models.AnswerOnFriendshipSuggestionReq) error
	GetAllFriendshipSuggestions(ctx context.Context, userID models.ID) (models.Users, error)
	GetAllFriends(ctx context.Context, userID models.ID) (models.Users, error)
}
