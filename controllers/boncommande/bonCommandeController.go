package boncommande

import (
	"fmt"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/kgermando/i-pos-stock/database"
	"github.com/kgermando/i-pos-stock/models"
)

// Paginate
func GetPaginatedBonCommande(c *fiber.Ctx) error {
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

	var dataList []models.BonCommande

	var length int64
	db.Model(dataList).Count(&length)
	db.Where("n_commande ILIKE ?", "%"+search+"%").
		Offset(offset).
		Limit(limit).
		Order("bon_commandes.updated_at DESC").
		Preload("BonCommadeLines").
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
		"message":    "All bonCommandes",
		"data":       dataList,
		"pagination": pagination,
	})
}

// Get All data
func GetAllBonCommandes(c *fiber.Ctx) error {
	db := database.DB
	var data []models.BonCommande
	db.Preload("Users").Preload("Pos").Find(&data)
	return c.JSON(fiber.Map{
		"status":  "success",
		"message": "All bonCommandes",
		"data":    data,
	})
}

// Get one data
func GetBonCommande(c *fiber.Ctx) error {
	id := c.Params("id")
	db := database.DB
	var bonCommande models.BonCommande
	db.Preload("Users").Preload("Pos").Find(&bonCommande, id)
	if bonCommande.NCommande == 0 {
		return c.Status(404).JSON(
			fiber.Map{
				"status":  "error",
				"message": "No bonCommande NCommande found",
				"data":    nil,
			},
		)
	}
	return c.JSON(
		fiber.Map{
			"status":  "success",
			"message": "bonCommande found",
			"data":    bonCommande,
		},
	)
}

// Create data
func CreateBonCommande(c *fiber.Ctx) error {
	p := &models.BonCommande{}

	if err := c.BodyParser(&p); err != nil {
		return err
	}

	database.DB.Create(p)

	return c.JSON(
		fiber.Map{
			"status":  "success",
			"message": "bonCommande created success",
			"data":    p,
		},
	)
}

// Update data
func UpdateBonCommande(c *fiber.Ctx) error {
	id := c.Params("id")
	db := database.DB

	type UpdateData struct {
		PosID         uint      `gorm:"index"`
		NCommande     uint64    `gorm:"not null" json:"n_commande"` // Number Random
		DateCommande  time.Time `gorm:"not null" json:"date_commande"`
		DateLivraison time.Time `gorm:"not null" json:"date_livraison"`
		FournisseurID uint
		Status        string  `gorm:"not null" json:"status"` // brouillon, valide, annule
		MontantTotal  float64 `gorm:"not null" json:"montant_total"`
		Notes         string  `json:"notes"`
		Signature     string  `json:"signature"`
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

	bonCommande := new(models.BonCommande)

	db.First(&bonCommande, id)
	bonCommande.PosID = updateData.PosID
	bonCommande.NCommande = updateData.NCommande
	bonCommande.DateCommande = updateData.DateCommande
	bonCommande.DateLivraison = updateData.DateLivraison
	bonCommande.FournisseurID = updateData.FournisseurID
	bonCommande.Status = updateData.Status
	bonCommande.MontantTotal = updateData.MontantTotal
	bonCommande.Notes = updateData.Notes
	bonCommande.Signature = updateData.Signature 

	db.Save(&bonCommande)

	return c.JSON(
		fiber.Map{
			"status":  "success",
			"message": "bonCommande updated success",
			"data":    bonCommande,
		},
	)

}

// Delete data
func DeleteBonCommande(c *fiber.Ctx) error {
	id := c.Params("id")

	db := database.DB

	var bonCommande models.BonCommande
	db.First(&bonCommande, id)
	if bonCommande.NCommande == 0 {
		return c.Status(404).JSON(
			fiber.Map{
				"status":  "error",
				"message": "No bonCommande NCommande found",
				"data":    nil,
			},
		)
	}

	db.Delete(&bonCommande)

	return c.JSON(
		fiber.Map{
			"status":  "success",
			"message": "BonCommande deleted success",
			"data":    nil,
		},
	)
}
