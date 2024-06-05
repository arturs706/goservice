package controllers

import (
	"example.com/gouserservice/domain"
	"example.com/gouserservice/usecases"
	"github.com/gofiber/fiber/v2"
	"log"
	"example.com/gouserservice/infrastructure/security"
	"strings"
	"github.com/google/uuid"
	"os"
	"example.com/gouserservice/infrastructure/auth"
	"fmt"

)

type UserController struct {
	UserInteractor usecases.UserInteractor
}

func NewUserController(interactor usecases.UserInteractor) UserController {
	return UserController{interactor}
}

func (controller *UserController) CreateLocalUser(c *fiber.Ctx) error {
    user := new(domain.User)

    if err := c.BodyParser(user); err != nil {
        log.Println(err)
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
            "status": "error",
			"message": "Invalid request payload",
        })
    }

	if user.UserID == "" {
        user.UserID = uuid.New().String()
    }

	if user.AuthMethod == "" {
        user.AuthMethod = "LOCAL"
    }

    hashedPassword, err := security.HashPassword(user.Passwd)
    if err != nil {
        log.Println(err)
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
            "status": "error",
			"message": "Failed to hash password",
        })
    }

    user.Passwd = hashedPassword

	secretKey := os.Getenv("VERIFICATION_TOKEN")
	if secretKey == "" {
		log.Fatalf("VERIFICATION_TOKEN not set in .env file")
	}
	jwtHandler := auth.NewJWTHandler(secretKey)
	token, err := jwtHandler.GenerateEmailVerToken(user)
	fmt.Println(token)
	user.EmailVerToken = token


    err = controller.UserInteractor.CreateLocalUser(user)
    if err != nil {
        log.Println(err)
        if strings.Contains(err.Error(), "duplicate key value violates unique constraint") {
            return c.Status(fiber.StatusConflict).JSON(fiber.Map{
                "status": "error",
				"message": "User already exists",
            })
        }
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
            "status": "error",
			"message": "Failed to create user",
        })
    }

    return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"status": "success",
		"email": user.Email,
	})
}


func (controller *UserController) CreateFacebookUser(c *fiber.Ctx) error {
	user := new(domain.User)

    if err := c.BodyParser(user); err != nil {
        log.Println(err)
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
            "status":  "error",
            "message": "Invalid request payload",
        })
    }

    log.Printf("Received user: %+v\n", user)

    if user.UserID == "" {
        user.UserID = uuid.New().String()
    }

    if user.AuthMethod == "" {
        user.AuthMethod = "FACEBOOK"
    }

    err := controller.UserInteractor.CreateGoogleUser(user)
    if err != nil {
        log.Println(err)
        if strings.Contains(err.Error(), "duplicate key value violates unique constraint") {
            return c.Status(fiber.StatusConflict).JSON(fiber.Map{
                "status":  "error",
                "message": "User already exists",
            })
        }
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
            "status":  "error",
            "message": "Failed to create user",
        })
    }

    return c.Status(fiber.StatusCreated).JSON(fiber.Map{
        "status": "success",
        "email": user.Email,
    })
}

func (controller *UserController) CreateGoogleUser(c *fiber.Ctx) error {
    user := new(domain.User)

    if err := c.BodyParser(user); err != nil {
        log.Println(err)
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
            "status":  "error",
            "message": "Invalid request payload",
        })
    }

    log.Printf("Received user: %+v\n", user)

    if user.UserID == "" {
        user.UserID = uuid.New().String()
    }

    if user.AuthMethod == "" {
        user.AuthMethod = "GOOGLE"
    }

    err := controller.UserInteractor.CreateGoogleUser(user)
    if err != nil {
        log.Println(err)
        if strings.Contains(err.Error(), "duplicate key value violates unique constraint") {
            return c.Status(fiber.StatusConflict).JSON(fiber.Map{
                "status":  "error",
                "message": "User already exists",
            })
        }
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
            "status":  "error",
            "message": "Failed to create user",
        })
    }

    return c.Status(fiber.StatusCreated).JSON(fiber.Map{
        "status": "success",
        "email": user.Email,
    })
}




func (controller *UserController) GetUserByID(c *fiber.Ctx) error {
	userID := c.Params("userID")
	user, err := controller.UserInteractor.GetUserByID(userID)
	if err != nil {
		log.Println(err)
		return err
	}
	return c.JSON(user)
}

func (controller *UserController) UpdateUser(c *fiber.Ctx) error {
	user := new(domain.User)
	if err := c.BodyParser(user); err != nil {
		log.Println(err)
		return err
	}

	err := controller.UserInteractor.UpdateUser(c.Context(), user) 
	if err != nil {
		log.Println(err)
		return err
	}
	return c.JSON(user)
}

func (controller *UserController) DeleteUser(c *fiber.Ctx) error {
	userID := c.Params("userID")
	err := controller.UserInteractor.DeleteUser(userID)
	if err != nil {
		log.Println(err)
		return err
	}
	return c.SendStatus(fiber.StatusNoContent)
}

func (controller *UserController) GetAllUsers(c *fiber.Ctx) error {
	users, err := controller.UserInteractor.GetAllUsers()
	if err != nil {
		log.Println(err)
		return err
	}
	return c.JSON(users)
}


func (controller *UserController) LoginUserController(c *fiber.Ctx) error {
    user := new(domain.User)
    if err := c.BodyParser(user); err != nil {
        log.Println(err)
        return err
    }
    dbUser, err := controller.UserInteractor.LoginUserUC(user.Email)
    if err != nil {
        log.Println(err)
        return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
            "status":  "error",
            "message": "Invalid email or password",
        })
    }

	user.Passwd = strings.TrimSpace(user.Passwd)
    err = security.ComparePassword(dbUser.Passwd, user.Passwd)
    if err != nil {
        log.Println(err)
        return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
            "status":  "error",
            "message": "Invalid email or password",
        })
    }
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"status": "success",
		"email": dbUser.Email,
	})
}