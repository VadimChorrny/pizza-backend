package api

import "github.com/pkg/errors"

var ErrInvalidOption = errors.New("invalid option")

type Options struct {
	HttpPort int `option:"required,not-empty"`
	//App      *app.App `option:"required,not-empty"`
}
