package services

import (
	"github.com/google/uuid"
	"pizza-backend/storage"
	"pizza-backend/storage/models"
	"pizza-backend/utils"
	"time"
)

type Opts struct {
	Storage *storage.Storage
}

type App struct {
	storage *storage.Storage
}

func New(opts Opts) (*App, error) {
	return &App{
		storage: opts.Storage,
	}, nil
}

func (a *App) CreateUser() error {
	user := &models.User{
		ID:        uuid.New().String(),
		Email:     "tempmail@gmail.com",
		Password:  "12345678",
		FirstName: "John",
		LastName:  "Doe",
		Role:      "admin",
		CreatedAt: time.Now(),
	}

	err := a.storage.CreateUser(user)
	if err != nil {
		utils.Logger().Error().Msg(err.Error())
		return err
	}

	return nil
}

func (a *App) GetUserByEmail(email string) (*models.User, error) {
	user, err := a.storage.GetUserByEmail(email)
	if err != nil {
		utils.Logger().Error().Msg(err.Error())
		return nil, err
	}

	return user, nil
}
