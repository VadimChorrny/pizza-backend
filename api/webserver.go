package api

import (
	"context"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"strconv"
)

type Service struct {
	opts Options
}

func New(opts Options) (*Service, error) {
	return &Service{
		opts: opts,
	}, nil
}

func cors(url string) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			c.Response().Header().Set("Access-Control-Allow-Origin", url)
			c.Response().Header().Set("Access-Control-Allow-Credentials", "true")
			c.Response().Header().Set("Access-Control-Allow-Methods", "GET, PUT, POST, DELETE, HEAD")
			return next(c)
		}
	}
}

func (s *Service) Run(ctx context.Context) error {
	e := echo.New()
	e.HideBanner = true
	e.HidePort = true

	e.Use(middleware.Logger())
	e.Use(cors("*")) // allow all origins

	e.GET("api/user", s.getUserByEmail)
	e.POST("api/user", s.createUser)

	addr := ":" + strconv.Itoa(s.opts.HttpPort)

	fmt.Println("RUN: ", addr)
	return e.Start(addr)
}

func (s *Service) createUser(c echo.Context) error {
	err := s.opts.App.CreateUser()
	if err != nil {
		return c.String(500, err.Error())
	}

	return c.JSON(200, "User created!")
}

func (s *Service) getUserByEmail(c echo.Context) error {
	email := c.QueryParams().Get("email")
	if email == "" {
		return c.String(400, "Email is required")
	}

	user, err := s.opts.App.GetUserByEmail(email)
	if err != nil {
		return c.String(500, err.Error())
	}

	return c.JSON(200, user)
}
