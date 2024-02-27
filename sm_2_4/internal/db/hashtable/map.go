package hashtable

import (
	"context"
	"fmt"
	"github.com/alserov/gb/sm_2_4/internal/db"
	"github.com/alserov/gb/sm_2_4/internal/models"
	"strings"
)

func NewHashtableRepository() db.Repository {
	return &repository{
		names:       map[string]struct{}{},
		users:       map[string]models.User{},
		friendships: map[string]models.Friendship{},
	}
}

type repository struct {
	names       map[string]struct{}
	users       map[string]models.User
	friendships map[string]models.Friendship
}

func (r *repository) CreteUser(ctx context.Context, req models.CreateUserReq) error {
	if _, ok := r.names[req.Username]; ok {
		return &models.Error{
			Msg:  fmt.Sprintf("user with username: %s already exists", req.Username),
			Type: models.ERR_BAD_REQUEST,
		}
	}

	r.users[string(req.ID)] = models.User{
		ID:       req.ID,
		Username: req.Username,
		Password: req.Password,
		Age:      req.Age,
	}

	r.names[req.Username] = struct{}{}

	return nil
}

func (r *repository) DeleteUser(ctx context.Context, userID models.ID) error {
	v, ok := r.users[string(userID)]
	if !ok {
		return &models.Error{
			Msg:  fmt.Sprintf("user not found"),
			Type: models.ERR_NOT_FOUND,
		}
	}

	delete(r.names, v.Username)
	delete(r.users, string(userID))

	return nil
}

func (r *repository) SuggestFriendship(ctx context.Context, receiverID models.ID, initiatorID models.ID) error {
	frshipID := fmt.Sprintf("%s%s", receiverID, initiatorID)

	if _, ok := r.friendships[frshipID]; ok {
		return &models.Error{
			Msg:  fmt.Sprintf("you have already sent friendship suggestion to this user"),
			Type: models.ERR_BAD_REQUEST,
		}
	}

	r.friendships[frshipID] = models.Friendship{
		InitiatorID: initiatorID,
		ReceiverID:  receiverID,
		Status:      models.WAITING,
	}

	return nil
}

func (r *repository) AnswerOnFriendshipSuggestion(ctx context.Context, req models.AnswerOnFriendshipSuggestionReq) error {
	frshipID := fmt.Sprintf("%s%s", req.ReceiverID, req.InitiatorID)

	sugg, ok := r.friendships[frshipID]
	if !ok {
		return &models.Error{
			Msg:  fmt.Sprintf("suggestion not found"),
			Type: models.ERR_NOT_FOUND,
		}
	}
	if sugg.Status != models.WAITING {
		return &models.Error{
			Msg:  fmt.Sprintf("you have already answered on this suggestion"),
			Type: models.ERR_NOT_FOUND,
		}
	}

	r.friendships[frshipID] = models.Friendship{
		ReceiverID:  req.ReceiverID,
		InitiatorID: req.InitiatorID,
		Status:      req.Answer,
	}

	return nil
}

func (r *repository) GetAllFriendshipSuggestions(ctx context.Context, userID models.ID) (models.Users, error) {
	var usrs models.Users
	for id, frshp := range r.friendships {
		if strings.Contains(id, string(userID)) {
			usr, _ := r.users[string(frshp.InitiatorID)]
			usrs = append(usrs, models.UserInfo{
				ID:       models.ID(id),
				Username: usr.Username,
				Age:      usr.Age,
			})
		}
	}

	return usrs, nil
}

func (r *repository) GetAllFriends(ctx context.Context, userID models.ID) (models.Users, error) {
	var usrs models.Users
	for id, frshp := range r.friendships {
		if strings.Contains(id, string(userID)) && frshp.Status == models.ACCEPTED {
			usr, _ := r.users[string(userID)]
			usrs = append(usrs, models.UserInfo{
				ID:       models.ID(id),
				Username: usr.Username,
				Age:      usr.Age,
			})
		}
	}

	return usrs, nil
}
