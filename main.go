package main

import (
	"context"
	"fmt"
	"github.com/joho/godotenv"
	"github.com/urfave/cli"
	"os"
	"pizza-backend/api"
	"pizza-backend/utils"
)

func main() {
	a := cli.NewApp()
	a.Version = "1.0.0"
	a.Name = "vf app"
	a.Usage = "cli for vf app"
	a.Commands = []cli.Command{
		{
			Name:   "run",
			Usage:  "Generate and write result",
			Action: cmdRun,
			Flags:  []cli.Flag{},
		},
	}

	if err := a.Run(os.Args); err != nil {
		utils.Logger().Err(err).Msg("failed to run command")
		fmt.Printf("%v\n", err)
		//panic("cannot run command: " + err.Error())
		os.Exit(1)
	}
}

func cmdRun(c *cli.Context) error {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println("Error loading .env file", err)
	}
	//stor, err := storage.New(os.Getenv("DATABASE_URL"), storage.Config{
	//	MaxOpenConns:    20,
	//	MaxIdleConns:    5,
	//	ConnMaxLifetime: 1 * time.Hour,
	//})
	if err != nil {
		utils.Logger().Err(err).Msg("failed to init storage")
		return err
	}
	//app, err := app.New(stor)
	//if err != nil {
	//	utils.Logger().Err(err).Msg("cannot create app")
	//	return err
	//}

	api, err := api.New(api.Options{
		HttpPort: 80,
		//App:      app,
	})

	if err != nil {
		panic(err)
	}

	api.Run(context.Background())
	return nil
}
