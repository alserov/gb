package service

import (
	"context"
	"github.com/alserov/gb/sm_2_4/internal/db/hashtable"
	"github.com/alserov/gb/sm_2_4/internal/log"
	"github.com/alserov/gb/sm_2_4/internal/models"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestService_SuggestFriendship_And_AnswerWithAccept(t *testing.T) {
	s := NewService(log.Logger{}, hashtable.NewHashtableRepository())

	initiatorID, err := s.CreateUser(context.Background(), models.CreateUserReq{
		User: models.User{
			Username: "a",
			Password: "b",
			Age:      3,
		},
	})
	require.NoError(t, err)
	require.NotEmpty(t, initiatorID)

	receiverID, err := s.CreateUser(context.Background(), models.CreateUserReq{
		User: models.User{
			Username: "b",
			Password: "c",
			Age:      4,
		},
	})
	require.NoError(t, err)
	require.NotEmpty(t, receiverID)

	err = s.SuggestFriendship(context.Background(), receiverID, initiatorID)
	require.NoError(t, err)

	sugg, err := s.GetAllFriendshipSuggestions(context.Background(), receiverID)
	require.NoError(t, err)
	require.Equal(t, 1, len(sugg))

	err = s.AnswerOnFriendshipSuggestion(context.Background(), models.AnswerOnFriendshipSuggestionReq{
		InitiatorID: initiatorID,
		ReceiverID:  receiverID,
		Answer:      models.ACCEPTED,
	})
	require.NoError(t, err)

	frnds, err := s.GetAllFriends(context.Background(), receiverID)
	require.NoError(t, err)
	require.Equal(t, 1, len(frnds))
}
