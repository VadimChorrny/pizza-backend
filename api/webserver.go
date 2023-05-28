package api

import (
	"context"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"strconv"
)

const jsonContentType = "application/json"

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
	e.Use(cors("*"))

	e.GET("/api/v1/hello-world", s.helloWorld)

	addr := ":" + strconv.Itoa(s.opts.HttpPort)

	fmt.Println("RUN: ", addr)
	return e.Start(addr)
}

func (s *Service) helloWorld(c echo.Context) error {
	return c.String(200, "Hello world")
}
