package commande

import (
	"fmt"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/kgermando/i-pos-stock/database"
	"github.com/kgermando/i-pos-stock/models"
)

// Paginate
func GetPaginatedCommandeLine(c *fiber.Ctx) error {
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

	var dataList []models.CommandeLine

	var length int64
	db.Model(dataList).Count(&length)
	db.Joins("JOIN commandes ON commande_lines.commande_id=commandes.id").
		Joins("JOIN products ON commande_lines.product_id=products.id").
		Where("products.name ILIKE ? OR products.reference ILIKE ?", "%"+search+"%", "%"+search+"%").
		Select(`
			commande_lines.id AS id,
			products.reference AS reference,
			products.name AS name,
			products.description AS description,
			products.unite_ventes.name AS unite,
			commande_lines.quantity AS quantity,
			commande_lines.prix_vente AS prix_vente
		`).
		Offset(offset).
		Limit(limit).
		Order("commande_lines.updated_at DESC"). 
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
		"message":    "All commandeLines",
		"data":       dataList,
		"pagination": pagination,
	})
}

// Query all data ID
func GetPaginatedCommandeLineByID(c *fiber.Ctx) error {
	db := database.DB
	commandeID := c.Params("commande_id")

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

	var dataList []models.CommandeLine

	var length int64
	var data []models.CommandeLine
	db.Model(data).Where("commande_id = ?", commandeID).Count(&length)
	db.Joins("JOIN commandes ON commande_lines.commande_id=commandes.id").
		Joins("JOIN products ON commande_lines.product_id=products.id").
		Where("commande_lines.commande_id = ?", commandeID).
		Where("products.name ILIKE ? OR products.reference ILIKE ?", "%"+search+"%", "%"+search+"%").
		Select(`
			commande_lines.id AS id,
			products.reference AS reference,
			products.name AS name,
			products.description AS description,
			products.unite_ventes.name AS unite,
			commande_lines.quantity AS quantity,
			commande_lines.prix_vente AS prix_vente
		`).
		Offset(offset).
		Limit(limit).
		Order("commande_lines.updated_at DESC").
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
		"message":    "All commandeLine by commande",
		"data":       dataList,
		"pagination": pagination,
	})
}

// Get All data
func GetAllCommandeLines(c *fiber.Ctx) error {
	db := database.DB
	var data []models.CommandeLine
	db.Find(&data)
	return c.JSON(fiber.Map{
		"status":  "success",
		"message": "All commandes",
		"data":    data,
	})
}

// Get one data
func GetCommandeLine(c *fiber.Ctx) error {
	id := c.Params("id")
	db := database.DB
	var commandeLine models.CommandeLine
	db.Find(&commandeLine, id)
	if commandeLine.ProductID == 0 {
		return c.Status(404).JSON(
			fiber.Map{
				"status":  "error",
				"message": "No commandeLine name found",
				"data":    nil,
			},
		)
	}
	return c.JSON(
		fiber.Map{
			"status":  "success",
			"message": "commandeLine found",
			"data":    commandeLine,
		},
	)
}

// Create data
func CreateCommandeLine(c *fiber.Ctx) error {
	p := &models.CommandeLine{}

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
func UpdateCommandeLine(c *fiber.Ctx) error {
	id := c.Params("id")
	db := database.DB

	type UpdateData struct {
		CommandeID uint
		ProductID  uint
		Quantity   uint64  `json:"quantity"`
		PrixVente  float64 `json:"prix_vente"`
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

	commandeLine := new(models.CommandeLine)

	db.First(&commandeLine, id)
	commandeLine.CommandeID = updateData.CommandeID
	commandeLine.ProductID = updateData.ProductID
	commandeLine.Quantity = updateData.Quantity
	commandeLine.PrixVente = updateData.PrixVente

	db.Save(&commandeLine)

	return c.JSON(
		fiber.Map{
			"status":  "success",
			"message": "commandeLine updated success",
			"data":    commandeLine,
		},
	)

}

// Delete data
func DeleteCommandeLine(c *fiber.Ctx) error {
	id := c.Params("id")

	db := database.DB

	var commandeLine models.CommandeLine
	db.First(&commandeLine, id)
	if commandeLine.ProductID == 0 {
		return c.Status(404).JSON(
			fiber.Map{
				"status":  "error",
				"message": "No commandeLine name found",
				"data":    nil,
			},
		)
	}

	db.Delete(&commandeLine)

	return c.JSON(
		fiber.Map{
			"status":  "success",
			"message": "bonCommandeLine deleted success",
			"data":    nil,
		},
	)
}
