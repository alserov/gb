package server

import (
	"errors"
	"github.com/alserov/fuze"
	"github.com/alserov/gb/sm_2_4/internal/models"
	"net/http"
)

func handleErr(c *fuze.Ctx, err error) {
	var e *models.Error
	errors.As(err, &e)

	switch e.Type {
	case models.ERR_INTERNAL:
		c.SendStatus(http.StatusInternalServerError)
		break
	case models.ERR_BAD_REQUEST:
		c.SendValue(map[string]string{
			"error": e.Msg,
		}, http.StatusBadRequest)
		break
	case models.ERR_NOT_FOUND:
		c.SendValue(map[string]string{
			"error": e.Msg,
		}, http.StatusNotFound)
		break
	}
}
