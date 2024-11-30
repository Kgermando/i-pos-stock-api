package boncommande

import (
	"fmt"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/kgermando/i-pos-stock/database"
	"github.com/kgermando/i-pos-stock/models"
)

// Paginate
func GetPaginatedBonCommandeLine(c *fiber.Ctx) error {
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

	var dataList []models.BonCommandeLine

	var length int64
	db.Model(dataList).Count(&length)
	db.Joins("JOIN commandes ON bon_commande_lines.bon_commande_id=commandes.id").
		Joins("JOIN products ON bon_commande_lines.product_id=products.id"). 
		Where("products.name ILIKE ? OR products.reference ILIKE ?", "%"+search+"%", "%"+search+"%").
		Select(`
			bon_commande_lines.id AS id,
			products.reference AS reference,
			products.name AS name,
			products.description AS description,
			products.unite_ventes.name AS unite,
			bon_commande_lines.quantity AS quantity,
			bon_commande_lines.price_unit AS price_unit
		`).
		Offset(offset).
		Limit(limit).
		Order("bon_commande_lines.updated_at DESC").
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
		"message":    "All bonCommandeLines",
		"data":       dataList,
		"pagination": pagination,
	})
}

// Query all data ID
func GetPaginatedBonCommandeLineByID(c *fiber.Ctx) error {
	db := database.DB
	BonCommandeID := c.Params("bon_commande_id")

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

	var dataList []models.BonCommandeLine

	var length int64
	var data []models.BonCommandeLine
	db.Model(data).Where("bon_commande_id = ?", BonCommandeID).Count(&length)
	db.Joins("JOIN commandes ON bon_commande_lines.bon_commande_id=commandes.id").
		Joins("JOIN products ON bon_commande_lines.product_id=products.id").
		Where("bon_commande_lines.bon_commande_id = ?", BonCommandeID).
		Where("products.name ILIKE ? OR products.reference ILIKE ?", "%"+search+"%", "%"+search+"%").
		Select(`
			bon_commande_lines.id AS id,
			products.reference AS reference,
			products.name AS name,
			products.description AS description,
			products.unite_ventes.name AS unite,
			bon_commande_lines.quantity AS quantity,
			bon_commande_lines.price_unit AS price_unit
		`).
		Offset(offset).
		Limit(limit).
		Order("bon_commande_lines.updated_at DESC").
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
		"message":    "All bonCommandeLine by BonCommande",
		"data":       dataList,
		"pagination": pagination,
	})
}

// Get All data
func GetAllBonCommandeLines(c *fiber.Ctx) error {
	db := database.DB
	var data []models.BonCommandeLine
	db.Find(&data)
	return c.JSON(fiber.Map{
		"status":  "success",
		"message": "All bonCommandeLines",
		"data":    data,
	})
}

// Get one data
func GetBonCommandeLine(c *fiber.Ctx) error {
	id := c.Params("id")
	db := database.DB
	var bonCommandeLine models.BonCommandeLine
	db.Find(&bonCommandeLine, id)
	if bonCommandeLine.ProductID == 0 {
		return c.Status(404).JSON(
			fiber.Map{
				"status":  "error",
				"message": "No bonCommandeLine name found",
				"data":    nil,
			},
		)
	}
	return c.JSON(
		fiber.Map{
			"status":  "success",
			"message": "bonCommandeLine found",
			"data":    bonCommandeLine,
		},
	)
}

// Create data
func CreateBonCommandeLine(c *fiber.Ctx) error {
	p := &models.BonCommandeLine{}

	if err := c.BodyParser(&p); err != nil {
		return err
	}

	database.DB.Create(p)

	return c.JSON(
		fiber.Map{
			"status":  "success",
			"message": "commande created success",
			"data":    p,
		},
	)
}

// Update data
func UpdateBonCommandeLine(c *fiber.Ctx) error {
	id := c.Params("id")
	db := database.DB

	type UpdateData struct {
		BonCommandeID uint
		ProductID     uint
		Quantity      uint64  `gorm:"not null" json:"quantity"`
		PriceUnit     float64 `gorm:"not null" json:"price_unit"`
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

	bonCommandeLine := new(models.BonCommandeLine)

	db.First(&bonCommandeLine, id)
	bonCommandeLine.BonCommandeID = updateData.BonCommandeID
	bonCommandeLine.ProductID = updateData.ProductID
	bonCommandeLine.Quantity = updateData.Quantity
	bonCommandeLine.PriceUnit = updateData.PriceUnit

	db.Save(&bonCommandeLine)

	return c.JSON(
		fiber.Map{
			"status":  "success",
			"message": "bonCommandeLine updated success",
			"data":    bonCommandeLine,
		},
	)

}

// Delete data
func DeleteBonCommandeLine(c *fiber.Ctx) error {
	id := c.Params("id")

	db := database.DB

	var bonCommandeLine models.BonCommandeLine
	db.First(&bonCommandeLine, id)
	if bonCommandeLine.ProductID == 0 {
		return c.Status(404).JSON(
			fiber.Map{
				"status":  "error",
				"message": "No bonCommandeLine name found",
				"data":    nil,
			},
		)
	}

	db.Delete(&bonCommandeLine)

	return c.JSON(
		fiber.Map{
			"status":  "success",
			"message": "bonCommandeLine deleted success",
			"data":    nil,
		},
	)
}
