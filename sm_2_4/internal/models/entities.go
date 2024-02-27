package models

type User struct {
	ID       ID     `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Age      int    `json:"age"`
}

type UserInfo struct {
	ID       ID     `json:"id"`
	Username string `json:"username"`
	Age      int    `json:"age"`
}

type Friendship struct {
	InitiatorID ID               `json:"initiatorID"`
	ReceiverID  ID               `json:"targetID"`
	Status      FriendshipStatus `json:"status"`
}

const (
	REFUSED FriendshipStatus = iota
	ACCEPTED
	WAITING
)

type (
	Users            []UserInfo
	ID               string
	FriendshipStatus uint
)
