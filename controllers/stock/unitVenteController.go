package stock

import (
	"fmt"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/kgermando/i-pos-stock/database"
	"github.com/kgermando/i-pos-stock/models"
)

// Paginate
func GetPaginatedUniteVente(c *fiber.Ctx) error {
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

	var dataList []models.UniteVente

	var length int64
	db.Model(dataList).Count(&length)
	db.
		Where("name ILIKE ?", "%"+search+"%").
		Offset(offset).
		Limit(limit).
		Order("unite_ventes.updated_at DESC").
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
		"message":    "All uniteVentes",
		"data":       dataList,
		"pagination": pagination,
	})
}

// Get All data
func GetAllUniteVentes(c *fiber.Ctx) error {
	db := database.DB
	var data []models.UniteVente
	db.Find(&data)
	return c.JSON(fiber.Map{
		"status":  "success",
		"message": "All uniteVentes",
		"data":    data,
	})
}

// Get one data
func GetUniteVente(c *fiber.Ctx) error {
	id := c.Params("id")
	db := database.DB
	var uniteVente models.UniteVente
	db.Find(&uniteVente, id)
	if uniteVente.Name == "" {
		return c.Status(404).JSON(
			fiber.Map{
				"status":  "error",
				"message": "No uniteVente name found",
				"data":    nil,
			},
		)
	}
	return c.JSON(
		fiber.Map{
			"status":  "success",
			"message": "uniteVente found",
			"data":    uniteVente,
		},
	)
}

// Create data
func CreateUniteVente(c *fiber.Ctx) error {
	p := &models.UniteVente{}

	if err := c.BodyParser(&p); err != nil {
		return err
	}

	database.DB.Create(p)

	return c.JSON(
		fiber.Map{
			"status":  "success",
			"message": "uniteVente created success",
			"data":    p,
		},
	)
}

// Update data
func UpdateUniteVente(c *fiber.Ctx) error {
	id := c.Params("id")
	db := database.DB

	type UpdateData struct {
		Name      string `gorm:"not null" json:"name"`
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

	uniteVente := new(models.UniteVente)

	db.First(&uniteVente, id)
	uniteVente.Name = updateData.Name
	uniteVente.Signature = updateData.Signature

	db.Save(&uniteVente)

	return c.JSON(
		fiber.Map{
			"status":  "success",
			"message": "uniteVente updated success",
			"data":    uniteVente,
		},
	)

}

// Delete data
func DeleteUniteVente(c *fiber.Ctx) error {
	id := c.Params("id")

	db := database.DB

	var uniteVente models.UniteVente
	db.First(&uniteVente, id)
	if uniteVente.Name == "" {
		return c.Status(404).JSON(
			fiber.Map{
				"status":  "error",
				"message": "No uniteVente name found",
				"data":    nil,
			},
		)
	}

	db.Delete(&uniteVente)

	return c.JSON(
		fiber.Map{
			"status":  "success",
			"message": "uniteVente deleted success",
			"data":    nil,
		},
	)
}
