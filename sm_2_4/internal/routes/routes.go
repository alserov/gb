package routes

import (
	"fmt"
	"github.com/alserov/fuze"
	"github.com/alserov/gb/sm_2_4/internal/server"
	"net/http"
)

func Setup(ctrl *fuze.Controller, s server.Server) {
	ctrl.POST("frshp/suggest/{initiator_id}/{receiver_id}", s.SuggestFriendship)
	ctrl.POST("frshp/answer/{initiator_id}/{receiver_id}", s.AnswerOnFriendshipSuggestion)
	ctrl.GET("frshp/{user_id}", s.GetAllFriendshipSuggestions)

	ctrl.POST("user", s.CreateUser)
	ctrl.DELETE("user/{user_id}", s.DeleteUser)
	ctrl.GET("user/{user_id}", s.GetAllFriends)

	ctrl.GET("/test", func(c *fuze.Ctx) {
		fmt.Println("received request")
		c.SendValue([]byte("ok"), http.StatusOK)
	})
}
