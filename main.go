package main

import (
	//"fmt"
	"fmt"
	"log"

	//"net/http"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"github.com/satti999/todoapp/config"
	"github.com/satti999/todoapp/src/route"
)

func main() {

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(err)
	}
	configration := &config.Config{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		Password: os.Getenv("DB_PASS"),
		User:     os.Getenv("DB_USER"),
		SSLMode:  os.Getenv("DB_SSLMODE"),
		DBName:   os.Getenv("DB_NAME"),
	}
	db, err := config.NewConnection(configration)
	if err != nil {
		log.Fatal("could not load the data base")
	}
	err = config.MigrateModels(db)
	if err != nil {
		log.Fatal("could not migrate db")
	}

	// repo := repository.NewRepository(db)

	//ser := service.NewService(repo)
	//hand := handler.NewHandler(ser)
	app := fiber.New()
	route.MainRouteHandler(db, app)

	// userRepo := repository.NewUserRepository(repo)
	// userSer := service.NewUserService(userRepo)

	// userHandler := handler.NewUserHandler(userSer)

	// handler.UserHandle(app, userHandler)
	//handler.UserHandler.UserHandle(app, userHand)
	err = app.Listen(":8080")
	if err != nil {
		fmt.Println("Error starting server:", err)
	}

}
