package models

type CreateUserReq struct {
	User
}

type DeleteUserReq struct {
	UserID   ID     `json:"initiatorID"`
	Password string `json:"password,omitempty"`
}

type AnswerOnFriendshipSuggestionReq struct {
	InitiatorID ID               `json:"initiatorID"`
	ReceiverID  ID               `json:"receiverID"`
	Answer      FriendshipStatus `json:"answer"`
}
