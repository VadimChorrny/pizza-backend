package main

import (
	"context"
	"fmt"
	"github.com/joho/godotenv"
	"github.com/urfave/cli"
	"os"
	"pizza-backend/api"
	"pizza-backend/services"
	"pizza-backend/storage"
	"pizza-backend/storage/migrations"
	"pizza-backend/utils"
	"time"
)

func main() {
	a := cli.NewApp()
	a.Version = "1.0.0"
	a.Name = "vf app"
	a.Usage = "cli for vf app"
	a.Commands = []cli.Command{
		{
			Name:   "run",
			Usage:  "Run web server",
			Action: cmdRun,
			Flags:  []cli.Flag{},
		},
		{
			Name:   "migrate",
			Usage:  "Generate and write result",
			Action: cmdMigrate,
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
		fmt.Println(err)
	}

	store, err := storage.New(os.Getenv("DATABASE_URL"), storage.Config{
		MaxOpenConns:    20,
		MaxIdleConns:    5,
		ConnMaxLifetime: 1 * time.Hour,
	})
	if err != nil {
		utils.Logger().Err(err).Msg("cannot create storage")
		return err
	}

	app, err := services.New(services.Opts{
		Storage: store,
	})
	if err != nil {
		utils.Logger().Err(err).Msg("cannot create app")
		return err
	}

	api, err := api.New(api.Options{
		HttpPort: 80,
		App:      app,
	})

	if err != nil {
		panic(err)
	}

	err = api.Run(context.Background())
	if err != nil {
		utils.Logger().Err(err).Msg("cannot run api")
		return err
	}

	return nil
}

func cmdMigrate(c *cli.Context) error {
	fmt.Println("ENVIRON")
	fmt.Println(os.Environ())
	err := godotenv.Load(".env")
	if err != nil {
		utils.Logger().Err(err).Msg("failed to load .env file")
		return err
	}

	utils.Logger().Info().Msg("migrations started, please wait...")
	utils.Logger().Info().Msgf("database url: %s", os.Getenv("DATABASE_URL"))

	db, err := storage.New(os.Getenv("DATABASE_URL"), storage.Config{})
	if err != nil {
		return err
	}
	return migrations.Migrate(db)
}
