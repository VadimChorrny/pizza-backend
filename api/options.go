package api

import (
	"pizza-backend/services"
)

type Options struct {
	HttpPort int           `option:"required,not-empty"`
	App      *services.App `option:"required,not-empty"`
}
