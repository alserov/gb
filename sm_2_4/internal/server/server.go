package server

import (
	"github.com/alserov/fuze"
	"github.com/alserov/gb/sm_2_4/internal/log"
	"github.com/alserov/gb/sm_2_4/internal/models"
	"github.com/alserov/gb/sm_2_4/internal/service"
	"net/http"
)

type Server interface {
	User
	Friendship
}

type Friendship interface {
	SuggestFriendship(c *fuze.Ctx)
	AnswerOnFriendshipSuggestion(c *fuze.Ctx)
	GetAllFriendshipSuggestions(c *fuze.Ctx)
	GetAllFriends(c *fuze.Ctx)
}

type User interface {
	CreateUser(c *fuze.Ctx)
	DeleteUser(c *fuze.Ctx)
}

func NewServer(log log.Logger, s service.Service) Server {
	return &server{
		log:  log,
		srvc: s,
	}
}

type server struct {
	log log.Logger

	srvc service.Service
}

func (s server) CreateUser(c *fuze.Ctx) {
	// также в DeleteUser, SuggestFriendship, AnswerOnFriendshipSuggestion, GetAllFriendshipSuggestions нужно проверять
	// может ли клиент редактировать такие данные, (например) в куки jwt токен засунуть и проверять его валидность
	var req models.CreateUserReq
	if err := c.Decode(&req); err != nil {
		c.SendStatus(http.StatusBadRequest)
		return
	}

	id, err := s.srvc.CreateUser(c.Request.Context(), req)
	if err != nil {
		handleErr(c, err)
		return
	}

	c.SendValue(map[string]string{
		"id": string(id),
	}, http.StatusCreated)
}

func (s server) DeleteUser(c *fuze.Ctx) {
	id := c.Parameters["user_id"]
	if id == "" {
		c.SendStatus(http.StatusNotFound)
		return
	}

	err := s.srvc.DeleteUser(c.Request.Context(), models.ID(id))
	if err != nil {
		handleErr(c, err)
		return
	}
}

func (s server) SuggestFriendship(c *fuze.Ctx) {
	receiverID, initiatorID := c.Parameters["receiver_id"], c.Parameters["initiator_id"]
	if receiverID == "" || initiatorID == "" {
		c.SendStatus(http.StatusNotFound)
		return
	}

	err := s.srvc.SuggestFriendship(c.Request.Context(), models.ID(receiverID), models.ID(initiatorID))
	if err != nil {
		handleErr(c, err)
		return
	}
}

func (s server) AnswerOnFriendshipSuggestion(c *fuze.Ctx) {
	receiverID, initiatorID := c.Parameters["receiver_id"], c.Parameters["initiator_id"]
	if receiverID == "" || initiatorID == "" {
		c.SendStatus(http.StatusNotFound)
		return
	}

	var req models.AnswerOnFriendshipSuggestionReq
	if err := c.Decode(&req); err != nil {
		c.SendStatus(http.StatusBadRequest)
		return
	}

	req.InitiatorID = models.ID(initiatorID)
	req.ReceiverID = models.ID(receiverID)

	err := s.srvc.AnswerOnFriendshipSuggestion(c.Request.Context(), req)
	if err != nil {
		handleErr(c, err)
		return
	}
}

func (s server) GetAllFriendshipSuggestions(c *fuze.Ctx) {
	id := c.Parameters["user_id"]
	if id == "" {
		c.SendStatus(http.StatusNotFound)
		return
	}

	suggs, err := s.srvc.GetAllFriendshipSuggestions(c.Request.Context(), models.ID(id))
	if err != nil {
		handleErr(c, err)
		return
	}

	c.SendValue(map[string][]models.UserInfo{
		"suggestions": suggs,
	}, http.StatusOK)
}

func (s server) GetAllFriends(c *fuze.Ctx) {
	id := c.Parameters["user_id"]
	if id == "" {
		c.SendStatus(http.StatusNotFound)
		return
	}

	frnds, err := s.srvc.GetAllFriends(c.Request.Context(), models.ID(id))
	if err != nil {
		handleErr(c, err)
		return
	}

	c.SendValue(map[string][]models.UserInfo{
		"friends": frnds,
	}, http.StatusOK)
}
