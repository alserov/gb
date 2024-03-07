package postgres

import (
	"context"
	"fmt"
	"github.com/alserov/gb/sm_2_4/internal/db"
	"github.com/alserov/gb/sm_2_4/internal/models"
	"github.com/jmoiron/sqlx"
)

func NewRepository(db *sqlx.DB) db.Repository {
	return &repository{
		db: db,
	}
}

type repository struct {
	db *sqlx.DB
}

func (r repository) CreteUser(ctx context.Context, req models.CreateUserReq) error {
	query := `INSERT INTO users (id,username,password,age) VALUES($1,$2,$3,$4)`

	if err := r.db.QueryRowx(query, req.ID, req.Username, req.Password, req.Age).Err(); err != nil {
		return &models.Error{
			Msg:  fmt.Sprintf("failed to execute equery: %v", err),
			Type: models.ERR_BAD_REQUEST,
		}
	}

	return nil
}

func (r repository) DeleteUser(ctx context.Context, userID models.ID) error {
	query := `DELETE FROM users WHERE id = $1`

	if err := r.db.QueryRowx(query, userID).Err(); err != nil {
		return &models.Error{
			Msg:  fmt.Sprintf("failed to execute equery: %v", err),
			Type: models.ERR_BAD_REQUEST,
		}
	}

	return nil
}

func (r repository) SuggestFriendship(ctx context.Context, receiverID models.ID, initiatorID models.ID) error {
	query := `INSERT INTO friendship (initiator_id,receiver_id,status) VALUES($1,$2,$3)`

	if err := r.db.QueryRowx(query, initiatorID, receiverID, models.WAITING).Err(); err != nil {
		return &models.Error{
			Msg:  fmt.Sprintf("failed to execute equery: %v", err),
			Type: models.ERR_BAD_REQUEST,
		}
	}

	return nil
}

func (r repository) AnswerOnFriendshipSuggestion(ctx context.Context, req models.AnswerOnFriendshipSuggestionReq) error {
	query := `UPDATE friendship SET status = $1 WHERE receiver_id = $2 AND initiator_id = $3`

	if err := r.db.QueryRowx(query, req.Answer, req.ReceiverID, req.InitiatorID).Err(); err != nil {
		return &models.Error{
			Msg:  fmt.Sprintf("failed to execute equery: %v", err),
			Type: models.ERR_BAD_REQUEST,
		}
	}

	return nil
}

func (r repository) GetAllFriendshipSuggestions(ctx context.Context, userID models.ID) (models.Users, error) {
	query := `SELECT id, username, age FROM users RIGHT JOIN friendship ON users.id = friendship.receiver_id WHERE receiver_id = $1`

	res, err := r.db.Queryx(query, userID)
	if err != nil {
		return nil, &models.Error{
			Msg:  fmt.Sprintf("failed to execute equery: %v", err),
			Type: models.ERR_BAD_REQUEST,
		}
	}

	var users models.Users
	for res.Next() {
		var u models.UserInfo
		if err = res.StructScan(&u); err != nil {
			return nil, &models.Error{
				Msg:  fmt.Sprintf("failed to execute equery: %v", err),
				Type: models.ERR_BAD_REQUEST,
			}
		}
		users = append(users, u)
	}

	return users, nil
}

func (r repository) GetAllFriends(ctx context.Context, userID models.ID) (models.Users, error) {
	query := `SELECT users.id,users.username, users.age FROM users RIGHT JOIN friendship on friendship.initiator_id = users.id  
                           WHERE receiver_id = $1 AND status = 1`

	res, err := r.db.Queryx(query, userID)
	if err != nil {
		return nil, &models.Error{
			Msg:  fmt.Sprintf("failed to execute equery: %v", err),
			Type: models.ERR_BAD_REQUEST,
		}
	}

	var users models.Users
	for res.Next() {
		var u models.UserInfo
		if err = res.StructScan(&u); err != nil {
			return nil, &models.Error{
				Msg:  fmt.Sprintf("failed to execute equery: %v", err),
				Type: models.ERR_BAD_REQUEST,
			}
		}
		users = append(users, u)
	}

	return users, nil
}
