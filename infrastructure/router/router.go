package router

import (
    "example.com/gouserservice/interface/controllers"
    "github.com/gofiber/fiber/v2"
)

type Router struct {
    Controller controllers.UserController
}

func NewRouter(controller controllers.UserController) Router {
    return Router{controller}
}

func (router *Router) Route(app *fiber.App) {
    apiV1 := app.Group("/api/v1")

    userRoutes := apiV1.Group("/users")
    userRoutes.Post("/", router.Controller.CreateLocalUser)
    userRoutes.Get("/id/:userID", router.Controller.GetUserByID)
    userRoutes.Put("/", router.Controller.UpdateUser)
    userRoutes.Delete("/:userID", router.Controller.DeleteUser)
    userRoutes.Get("/", router.Controller.GetAllUsers)
    userRoutes.Post("/login", router.Controller.LoginUserController)
    userRoutes.Post("/login/google", router.Controller.CreateGoogleUser)
    userRoutes.Post("/login/facebook", router.Controller.CreateFacebookUser)
}