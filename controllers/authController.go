package controllers

// import (
// 	"os"
// 	"strconv"
// 	"time"

// 	db "github.com/briannkhata/eschool/config"
// 	models "github.com/briannkhata/eschool/models"
// 	"github.com/gofiber/fiber/v2"
// 	"github.com/golang-jwt/jwt"
// )

// func Login(c *fiber.Ctx) error {
// 	cashierId := c.Params("cashierId")

// 	var data map[string]string

// 	err := c.BodyParser(&data)

// 	if err != nil {
// 		return c.Status(400).JSON(fiber.Map{
// 			"success": false,
// 			"message": "Invalid post request",
// 		})
// 	}

// 	if data["passcode"] == "" {
// 		return c.Status(400).JSON(fiber.Map{
// 			"success": false,
// 			"message": "Passcode is required",
// 			"error":   map[string]interface{}{},
// 		})
// 	}

// 	var cashier models.Cashier

// 	db.DB.Where("id=?", cashierId).First(&cashier)

// 	if cashier.Id == 0 {
// 		return c.Status(400).JSON(fiber.Map{
// 			"success": false,
// 			"message": "Cashier not Found",
// 			"error":   map[string]interface{}{},
// 		})
// 	}

// 	if cashier.Passcode != data["passcode"] {
// 		return c.Status(401).JSON(fiber.Map{
// 			"success": false,
// 			"message": "Passcode do not Match",
// 			"error":   map[string]interface{}{},
// 		})
// 	}

// 	token := jwt.NewWithClaims(jwt.SigningMethodES256, jwt.MapClaims{
// 		"Issuer":    strconv.Itoa(int(cashier.Id)),
// 		"ExpiresAt": time.Now().Add(time.Hour * 24).Unix(),
// 	})

// 	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))

// 	if err != nil {
// 		return c.Status(401).JSON(fiber.Map{
// 			"success": false,
// 			"message": "Token Expired",
// 		})
// 	}

// 	cashierData := make(map[string]interface{})
// 	cashierData["token"] = tokenString

// 	return c.Status(401).JSON(fiber.Map{
// 		"success": true,
// 		"message": "success",
// 		"data":    cashierData,
// 	})

// }

// func Logout(c *fiber.Ctx) error {
// 	return nil

// }

// func Passcode(c *fiber.Ctx) error {
// 	return nil
// }

import (
	"os"
	"strconv"
	"time"

	db "github.com/briannkhata/eschool/config"
	models "github.com/briannkhata/eschool/models"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
)

func Login(c *fiber.Ctx) error {
	cashierId := c.Params("cashierId")

	var data map[string]string
	err := c.BodyParser(&data)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"success": false,
			"message": "Invalid request body",
		})
	}

	if data["passcode"] == "" {
		return c.Status(400).JSON(fiber.Map{
			"success": false,
			"message": "Passcode is required",
		})
	}

	var cashier models.Cashier
	db.DB.Where("id = ?", cashierId).First(&cashier)

	if cashier.Id == 0 {
		return c.Status(404).JSON(fiber.Map{
			"success": false,
			"message": "Cashier not found",
		})
	}

	if cashier.Passcode != data["passcode"] {
		return c.Status(401).JSON(fiber.Map{
			"success": false,
			"message": "Incorrect passcode",
		})
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"Issuer":    strconv.Itoa(int(cashier.Id)),
		"ExpiresAt": time.Now().Add(time.Hour * 24).Unix(),
	})

	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"success": false,
			"message": "Could not generate token",
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"success": true,
		"message": "Login successful",
		"data":    map[string]interface{}{"token": tokenString},
	})
}

func Logout(c *fiber.Ctx) error {
	// Implementation for Logout (e.g., by revoking token or setting an expiry)
	return c.Status(200).JSON(fiber.Map{
		"success": true,
		"message": "Logout successful",
	})
}

func Passcode(c *fiber.Ctx) error {
	// Implementation for passcode-related operations
	return nil
}
