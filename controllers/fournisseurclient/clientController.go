package fournisseurclient

import (
	"fmt"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/kgermando/i-pos-stock/database"
	"github.com/kgermando/i-pos-stock/models"
)

// Paginate
func GetPaginatedClient(c *fiber.Ctx) error {
	db := database.DB

	page, err := strconv.Atoi(c.Query("page", "1"))
	if err != nil || page <= 0 {
		page = 1 // Default page number
	}
	limit, err := strconv.Atoi(c.Query("limit", "15"))
	if err != nil || limit <= 0 {
		limit = 15
	}
	offset := (page - 1) * limit

	search := c.Query("search", "")

	var dataList []models.Client

	var length int64
	db.Model(dataList).Count(&length)
	db.
		Where("fullname ILIKE ?", "%"+search+"%").
		Offset(offset).
		Limit(limit).
		Order("clients.updated_at DESC").
		Preload("Commandes").
		Find(&dataList)

	if err != nil {
		fmt.Println("error s'est produite: ", err)
		return c.Status(500).SendString(err.Error())
	}

	// Calculate total number of pages
	totalPages := len(dataList) / limit
	if remainder := len(dataList) % limit; remainder > 0 {
		totalPages++
	}
	pagination := map[string]interface{}{
		"total_pages": totalPages,
		"page":        page,
		"page_size":   limit,
		"length":      length,
	}

	return c.JSON(fiber.Map{
		"status":     "success",
		"message":    "All clients",
		"data":       dataList,
		"pagination": pagination,
	})
}

// Get All data
func GetAllClients(c *fiber.Ctx) error {
	db := database.DB
	var data []models.Client
	db.Preload("Commandes").Find(&data)
	return c.JSON(fiber.Map{
		"status":  "success",
		"message": "All clients",
		"data":    data,
	})
}

// Get one data
func GetClient(c *fiber.Ctx) error {
	id := c.Params("id")
	db := database.DB
	var client models.Client
	db.Preload("Commandes").Find(&client, id)
	if client.FullName == "" {
		return c.Status(404).JSON(
			fiber.Map{
				"status":  "error",
				"message": "No client FullName found",
				"data":    nil,
			},
		)
	}
	return c.JSON(
		fiber.Map{
			"status":  "success",
			"message": "client found",
			"data":    client,
		},
	)
}

// Create data
func CreateClient(c *fiber.Ctx) error {
	p := &models.Client{}

	if err := c.BodyParser(&p); err != nil {
		return err
	}

	database.DB.Create(p)

	return c.JSON(
		fiber.Map{
			"status":  "success",
			"message": "client created success",
			"data":    p,
		},
	)
}

// Update data
func UpdateClient(c *fiber.Ctx) error {
	id := c.Params("id")
	db := database.DB

	type UpdateData struct {
		FullName  string `json:"fullname"`
		Telephone string `json:"telephone"`
		Email     string `json:"email"`
		Signature string `json:"signature"`
	}

	var updateData UpdateData

	if err := c.BodyParser(&updateData); err != nil {
		return c.Status(500).JSON(
			fiber.Map{
				"status":  "error",
				"message": "Review your iunput",
				"data":    nil,
			},
		)
	}

	client := new(models.Client)

	db.First(&client, id)
	client.FullName = updateData.FullName
    client.Telephone = updateData.Telephone
    client.Email = updateData.Email
	client.Signature = updateData.Signature

	db.Save(&client)

	return c.JSON(
		fiber.Map{
			"status":  "success",
			"message": "client updated success",
			"data":    client,
		},
	)

}

// Delete data
func DeleteClient(c *fiber.Ctx) error {
	id := c.Params("id")

	db := database.DB

	var client models.Client
	db.First(&client, id)
	if client.FullName == "" {
		return c.Status(404).JSON(
			fiber.Map{
				"status":  "error",
				"message": "No client FullName found",
				"data":    nil,
			},
		)
	}

	db.Delete(&client)

	return c.JSON(
		fiber.Map{
			"status":  "success",
			"message": "client deleted success",
			"data":    nil,
		},
	)
}
