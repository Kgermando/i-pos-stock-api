package product

import (
	"fmt"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/kgermando/i-pos-stock/database"
	"github.com/kgermando/i-pos-stock/models"
)

// Paginate
func GetPaginatedCategory(c *fiber.Ctx) error {
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

	var dataList []models.Category

	var length int64
	db.Model(dataList).Count(&length)
	db.
		Where("name ILIKE ?", "%"+search+"%").
		Offset(offset).
		Limit(limit).
		Order("categories.updated_at DESC").
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
		"message":    "All categorys",
		"data":       dataList,
		"pagination": pagination,
	})
}

// Get All data
func GetAllCategorys(c *fiber.Ctx) error {
	db := database.DB
	var data []models.Category
	db.Find(&data)
	return c.JSON(fiber.Map{
		"status":  "success",
		"message": "All categorys",
		"data":    data,
	})
}

// Get one data
func GetCategory(c *fiber.Ctx) error {
	id := c.Params("id")
	db := database.DB
	var category models.Category
	db.Find(&category, id)
	if category.Name == "" {
		return c.Status(404).JSON(
			fiber.Map{
				"status":  "error",
				"message": "No category name found",
				"data":    nil,
			},
		)
	}
	return c.JSON(
		fiber.Map{
			"status":  "success",
			"message": "category found",
			"data":    category,
		},
	)
}

// Create data
func CreateCategory(c *fiber.Ctx) error {
	p := &models.Category{}

	if err := c.BodyParser(&p); err != nil {
		return err
	}

	database.DB.Create(p)

	return c.JSON(
		fiber.Map{
			"status":  "success",
			"message": "category created success",
			"data":    p,
		},
	)
}

// Update data
func UpdateCategory(c *fiber.Ctx) error {
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

	category := new(models.Category)

	db.First(&category, id)
	category.Name = updateData.Name
	category.Signature = updateData.Signature

	db.Save(&category)

	return c.JSON(
		fiber.Map{
			"status":  "success",
			"message": "category updated success",
			"data":    category,
		},
	)

}

// Delete data
func DeleteCategory(c *fiber.Ctx) error {
	id := c.Params("id")

	db := database.DB

	var category models.Category
	db.First(&category, id)
	if category.Name == "" {
		return c.Status(404).JSON(
			fiber.Map{
				"status":  "error",
				"message": "No category name found",
				"data":    nil,
			},
		)
	}

	db.Delete(&category)

	return c.JSON(
		fiber.Map{
			"status":  "success",
			"message": "category deleted success",
			"data":    nil,
		},
	)
}
