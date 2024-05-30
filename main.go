package main

import (
	"example.com/gouserservice/infrastructure/db"
	"example.com/gouserservice/infrastructure/router"
	"example.com/gouserservice/interface/controllers"
	"example.com/gouserservice/interface/repository"
	"example.com/gouserservice/usecases"
	"github.com/gofiber/fiber/v2"
	"log"
	"github.com/joho/godotenv"

)

func main() {
	app := fiber.New()
	if err := godotenv.Load(); err != nil {
        log.Fatal("Error loading .env file")
    }
	dbHandler := db.NewDBHandler()
	defer dbHandler.Close()
	userRepo := repository.NewUserRepo(&dbHandler)
	userInteractor := usecases.NewUserInteractor(userRepo)
	userController := controllers.NewUserController(userInteractor)
	userRouter := router.NewRouter(userController)
	userRouter.Route(app)
	log.Fatal(app.Listen(":2001"))
}