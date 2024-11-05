package controllers

import (
	"strconv"
	"time"

	db "github.com/briannkhata/eschool/config"
	models "github.com/briannkhata/eschool/models"
	"github.com/gofiber/fiber/v2"
)

func CreateCashier(c *fiber.Ctx) error {
	var data map[string]string
	err := c.BodyParser(&data)

	if err != nil {
		return c.Status(400).JSON(
			fiber.Map{
				"success": false,
				"message": "Invalid data",
			})
	}

	if data["name"] == "" {
		return c.Status(400).JSON(
			fiber.Map{
				"success": false,
				"message": "Name is Required",
			})
	}

	if data["passcode"] == "" {
		return c.Status(400).JSON(
			fiber.Map{
				"success": false,
				"message": "Passcode is Required",
			})
	}

	cashier := models.Cashier{
		Name:      data["name"],
		Passcode:  data["passcode"],
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
	}

	db.DB.Create(&cashier)
	return c.Status(200).JSON(fiber.Map{
		"success": true,
		"message": "Cashier added successfully",
		"data":    cashier,
	})

}

func GetCashierDetails(c *fiber.Ctx) error {
	cashierId := c.Params("cashierId")
	var cashier models.Cashier

	db.DB.Select("id,name,created_at,updated_at").Where("id=?", cashierId).First(&cashier)

	cashierData := make(map[string]interface{})
	cashierData["cashierId"] = cashier.Id
	cashierData["name"] = cashier.Name
	cashierData["createdAt"] = cashier.CreatedAt
	cashierData["updatedAt"] = cashier.UpdatedAt

	if cashier.Id == 0 {
		return c.Status(404).JSON(fiber.Map{
			"success": false,
			"message": "Cashier Not Found",
			"data":    map[string]interface{}{},
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"success": true,
		"message": "Success",
		"data":    cashierData,
	})

}

func CashiersList(c *fiber.Ctx) error {
	var cashier []models.Cashier

	limit, _ := strconv.Atoi(c.Query("limit"))
	skip, _ := strconv.Atoi(c.Query("skip"))
	var count int64

	db.DB.Select("*").Limit(limit).Offset(skip).Find(&cashier).Count(&count)
	return c.Status(200).JSON(
		fiber.Map{
			"success": true,
			"message": "Cashier list api",
			"data":    cashier,
		})
}

func EditCashier(c *fiber.Ctx) error {
	return nil
}

func UpdateCashier(c *fiber.Ctx) error {
	cashierId := c.Params("cashierId")
	var cashier models.Cashier

	db.DB.Find(&cashier, "id=?", cashierId)

	if cashier.Name == "" {
		return c.Status(404).JSON(fiber.Map{
			"success": false,
			"message": "Cashier Not Found",
		})
	}

	var updateCashier models.Cashier

	err := c.BodyParser(&updateCashier)
	if err != nil {
		return err
	}

	if updateCashier.Name == "" {
		return c.Status(404).JSON(fiber.Map{
			"success": false,
			"message": "Name is Required",
		})
	}

	cashier.Name = updateCashier.Name
	db.DB.Save(&cashier)

	return c.Status(200).JSON(fiber.Map{
		"success": true,
		"message": "success",
		"data":    cashier,
	})
}

func DeleteCashier(c *fiber.Ctx) error {
	return nil
}
