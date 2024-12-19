package product

import (
	"github.com/gofiber/fiber/v2"
	"github.com/kgermando/i-pos-stock/database"
	"github.com/kgermando/i-pos-stock/models"
)

// Paginate

// Get All data
func GetAllStockDispo(c *fiber.Ctx) error {
	db := database.DB
	codeEntreprise := c.Params("code_entreprise")
	posId := c.Params("pos_id")

	var data []models.StockDispo
	db.Where("code_entreprise = ?", codeEntreprise).
		Where("pos_id = ?", posId).
		Find(&data)
	return c.JSON(fiber.Map{
		"status":  "success",
		"message": "All StockDispo",
		"data":    data,
	})
}

// Get one data
func GetStockDispoSUM(c *fiber.Ctx) error {
	codeEntreprise := c.Params("code_entreprise")
	posId := c.Params("pos_id")
	productId := c.Params("product_id")
	db := database.DB

	var stockDispo models.StockDispo
	db.Where("code_entreprise = ?", codeEntreprise).
		Where("pos_id = ?", posId).
		Where("product_id = ?", productId).
		Select(`
			id AS id,
			price_stock AS price_stock,
			SUM(qty_stock::FLOAT) AS qty_stock,
			SUM(qty_cmdline::FLOAT) AS qty_cmdline,
			SUM(qty_stock::FLOAT - qty_cmdline::FLOAT) AS dispo,
			SUM(qty_cmdline::FLOAT * qty_stock::FLOAT / 100) AS pourcent
		`).
		Preload("Product").
		Find(&stockDispo)
	if stockDispo.ProductID == 0 {
		return c.Status(404).JSON(
			fiber.Map{
				"status":  "error",
				"message": "No StockDispo name found",
				"data":    nil,
			},
		)
	}
	return c.JSON(
		fiber.Map{
			"status":  "success",
			"message": "stockDispo found",
			"data":    stockDispo,
		},
	)
}

// Get One data
func GetStockDispo(c *fiber.Ctx) error {
	db := database.DB
	codeEntreprise := c.Params("code_entreprise")
	posId := c.Params("pos_id")
	productId := c.Params("product_id")

	var data models.StockDispo
	db.Where("code_entreprise = ?", codeEntreprise).
		Where("pos_id = ?", posId).
		Where("product_id = ?", productId).
		First(&data)
	return c.JSON(fiber.Map{
		"status":  "success",
		"message": "StockDispo",
		"data":    data,
	})
}

// Create data
func CreateStockDispo(c *fiber.Ctx) error {
	p := &models.StockDispo{}

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
func UpdateStockDispo(c *fiber.Ctx) error {
	id := c.Params("id")
	db := database.DB

	type UpdateData struct {
		PosID          uint    `json:"pos_id"`
		ProductID      uint    `json:"product_id"`
		QtyStock       uint64  `json:"qty_stock"`
		QtyCmdline     uint64  `json:"qty_cmdline"`
		PriceStock     float64 `json:"price_stock"`
		CodeEntreprise uint    `json:"code_entreprise"`
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

	stockDispo := new(models.StockDispo)

	db.First(&stockDispo, id)
	stockDispo.PosID = updateData.PosID
	stockDispo.ProductID = updateData.ProductID
	stockDispo.QtyStock = updateData.QtyStock
	stockDispo.QtyCmdline = updateData.QtyCmdline
	stockDispo.PriceStock = updateData.PriceStock
	stockDispo.CodeEntreprise = updateData.CodeEntreprise

	db.Save(&stockDispo)

	return c.JSON(
		fiber.Map{
			"status":  "success",
			"message": "stockDispo updated success",
			"data":    stockDispo,
		},
	)
}

// Delete data
func DeleteStockDispo(c *fiber.Ctx) error {
	id := c.Params("id")

	db := database.DB

	var stockDispo models.StockDispo
	db.First(&stockDispo, id)
	if stockDispo.ProductID == 0 {
		return c.Status(404).JSON(
			fiber.Map{
				"status":  "error",
				"message": "No stockDispo name found",
				"data":    nil,
			},
		)
	}

	db.Delete(&stockDispo)

	return c.JSON(
		fiber.Map{
			"status":  "success",
			"message": "stockDispo deleted success",
			"data":    nil,
		},
	)
}
