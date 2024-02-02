// handler/user.go
package handler

import (
	"log"
	"postgress_api/database"
    "postgress_api/model"
    "postgress_api/config"
	"github.com/gofiber/fiber/v2"
    "github.com/dgrijalva/jwt-go"
    "time"
)

func SignUp(c *fiber.Ctx) error {
	user := new(model.User)
	db := database.DB

	// Parse the request body to get user input
	if err := c.BodyParser(user); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Invalid request")
	}

	// Create the user in the database
	if err := db.Create(user).Error; err != nil {
		log.Println(err)
		return c.Status(fiber.StatusInternalServerError).SendString("Error creating user")
	}

	// Load JWT secret key from environment variable using the Config function
	secretKey := config.Config("JWT_SECRET")
	if secretKey == "" {
		log.Fatal("JWT_SECRET not set in .env file")
	}

	// Create a new JWT token
	token := jwt.New(jwt.SigningMethodHS256)

	// Set claims
	claims := token.Claims.(jwt.MapClaims)
	claims["id"] = user.ID
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix() // Token expires in 24 hours

	// Sign and get the complete encoded token as a string
	tokenString, err := token.SignedString([]byte(secretKey))
	if err != nil {
		log.Println(err)
		return c.Status(fiber.StatusInternalServerError).SendString("Failed to create token")
	}

	return c.JSON(fiber.Map{"token": tokenString})
}


func Login(c *fiber.Ctx) error {
	var userInput model.LoginInput
	db := database.DB

	// Parse the request body to get user input (email and password)
	if err := c.BodyParser(&userInput); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Invalid request")
	}

	// Fetch user from the database based on the provided email
	var userFromDB model.User
	if err := db.Where("email = ?", userInput.Email).First(&userFromDB).Error; err != nil {
		log.Println(err)
		return c.Status(fiber.StatusUnauthorized).SendString("Invalid email or password")
	}

	// Validate the provided password against the stored password
	if userInput.Password != userFromDB.Password {
		return c.Status(fiber.StatusUnauthorized).SendString("Invalid email or password")
	}

	// Load JWT secret key from environment variable using the Config function
	secretKey := config.Config("JWT_SECRET")
	if secretKey == "" {
		log.Fatal("JWT_SECRET not set in .env file")
	}

	// Create a new JWT token
	token := jwt.New(jwt.SigningMethodHS256)

	// Set claims
	claims := token.Claims.(jwt.MapClaims)
	claims["id"] = userFromDB.ID
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix() // Token expires in 24 hours

	// Sign and get the complete encoded token as a string
	tokenString, err := token.SignedString([]byte(secretKey))
	if err != nil {
		log.Println(err)
		return c.Status(fiber.StatusInternalServerError).SendString("Failed to create token")
	}

	return c.JSON(fiber.Map{"token": tokenString})
}