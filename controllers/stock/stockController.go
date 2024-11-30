package stock

import (
	"fmt"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/kgermando/i-pos-stock/database"
	"github.com/kgermando/i-pos-stock/models"
)

// Paginate
func GetPaginatedStock(c *fiber.Ctx) error {
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

	var dataList []models.Stock

	var length int64
	db.Model(dataList).Count(&length)
	db.
		Joins("JOIN products ON stocks.product_id=products.id").
		Joins("JOIN fournisseurs ON stocks.fournisseur_id=fournisseurs.id").
		Where("products.name ILIKE ? OR products.reference ILIKE ?", "%"+search+"%", "%"+search+"%").
		Select(`
			stocks.id AS id,
			products.reference AS reference,
			products.name AS name,
			stocks.quantity AS quantity,
			stocks.prix_achat AS prix_achat,
			stocks.prix_vente AS prix_vente,
			stocks.date_expiration AS date_expiration,
			fournisseurs.name AS fournisseur,
			stocks.signature AS signature,
		`).
		Offset(offset).
		Limit(limit).
		Order("stocks.updated_at DESC").
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
		"message":    "All stocks",
		"data":       dataList,
		"pagination": pagination,
	})
}

// Get All data
func GetAllStocks(c *fiber.Ctx) error {
	db := database.DB
	var data []models.Stock
	db.Find(&data)
	return c.JSON(fiber.Map{
		"status":  "success",
		"message": "All stocks",
		"data":    data,
	})
}

// Get one data
func GetStock(c *fiber.Ctx) error {
	id := c.Params("id")
	db := database.DB
	var stock models.Stock
	db.Find(&stock, id)
	if stock.ProductID == 0 {
		return c.Status(404).JSON(
			fiber.Map{
				"status":  "error",
				"message": "No stock name found",
				"data":    nil,
			},
		)
	}
	return c.JSON(
		fiber.Map{
			"status":  "success",
			"message": "stock found",
			"data":    stock,
		},
	)
}

// Create data
func CreateStock(c *fiber.Ctx) error {
	p := &models.Stock{}

	if err := c.BodyParser(&p); err != nil {
		return err
	}

	database.DB.Create(p)

	return c.JSON(
		fiber.Map{
			"status":  "success",
			"message": "stock created success",
			"data":    p,
		},
	)
}

// Update data
func UpdateStock(c *fiber.Ctx) error {
	id := c.Params("id")
	db := database.DB

	type UpdateData struct {
		PosID          uint
		ProductID      uint
		Description    string `json:"description"`
		FournisseurID  uint
		Quantity       uint64    `json:"quantity"`
		PrixAchat      float64   `json:"prix_achat"`
		PrixVente      float64   `json:"prix_vente"`
		DateExpiration time.Time `json:"date_expiration"`
		Signature      string    `json:"signature"`
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

	stock := new(models.Stock)

	db.First(&stock, id)
	stock.PosID = updateData.PosID
	stock.ProductID = updateData.ProductID
	stock.Description = updateData.Description
	stock.FournisseurID = updateData.FournisseurID
	stock.Quantity = updateData.Quantity
	stock.PrixAchat = updateData.PrixAchat
	stock.PrixVente = updateData.PrixVente
	stock.DateExpiration = updateData.DateExpiration
	stock.Signature = updateData.Signature

	db.Save(&stock)

	return c.JSON(
		fiber.Map{
			"status":  "success",
			"message": "stock updated success",
			"data":    stock,
		},
	)

}

// Delete data
func DeleteStock(c *fiber.Ctx) error {
	id := c.Params("id")

	db := database.DB

	var stock models.Stock
	db.First(&stock, id)
	if stock.ProductID == 0 {
		return c.Status(404).JSON(
			fiber.Map{
				"status":  "error",
				"message": "No stock name found",
				"data":    nil,
			},
		)
	}

	db.Delete(&stock)

	return c.JSON(
		fiber.Map{
			"status":  "success",
			"message": "stock deleted success",
			"data":    nil,
		},
	)
}
