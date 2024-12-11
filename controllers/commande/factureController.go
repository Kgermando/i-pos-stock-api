package commande

import (
	"fmt"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/kgermando/i-pos-stock/database"
	"github.com/kgermando/i-pos-stock/models"
)

// Paginate
func GetPaginatedFacture(c *fiber.Ctx) error {
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

	var dataList []models.Facture

	var length int64
	db.Model(dataList).Count(&length)
	db.
		Where("n_facture ILIKE ? status ILIKE ?", "%"+search+"%", "%"+search+"%").
		Offset(offset).
		Limit(limit).
		Order("factures.updated_at DESC").
		Preload("Commande"). 
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
		"message":    "All factures",
		"data":       dataList,
		"pagination": pagination,
	})
}

// Query all data ID
func GetPaginatedFactureByID(c *fiber.Ctx) error {
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

	var dataList []models.Facture

	var length int64
	var data []models.Facture
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
		"message":    "All factures",
		"data":       dataList,
		"pagination": pagination,
	})
}

// Get All data
func GetAllFactures(c *fiber.Ctx) error {
	db := database.DB
	var data []models.Facture
	db.Find(&data)
	return c.JSON(fiber.Map{
		"status":  "success",
		"message": "All factures",
		"data":    data,
	})
}

// Get one data
func GetFacture(c *fiber.Ctx) error {
	id := c.Params("id")
	db := database.DB
	var facture models.Facture
	db.Find(&facture, id)
	if facture.NFacture == 0 {
		return c.Status(404).JSON(
			fiber.Map{
				"status":  "error",
				"message": "No facture name found",
				"data":    nil,
			},
		)
	}
	return c.JSON(
		fiber.Map{
			"status":  "success",
			"message": "facture found",
			"data":    facture,
		},
	)
}

// Create data
func CreateFacture(c *fiber.Ctx) error {
	p := &models.Facture{}

	if err := c.BodyParser(&p); err != nil {
		return err
	}

	database.DB.Create(p)

	return c.JSON(
		fiber.Map{
			"status":  "success",
			"message": "facture created success",
			"data":    p,
		},
	)
}

// Update data
func UpdateFacture(c *fiber.Ctx) error {
	id := c.Params("id")
	db := database.DB

	type UpdateData struct {
		NFacture        float64  `json:"n_facture"`
	CommandeID      uint     `json:"commande_id"`
	Status          string   `json:"status"` // Cash ou Creance
	DelaiPaiement   string   `json:"delai_paiement"`
	Signature       string   `json:"signature"`
	PosID           uint     `json:"pos_id"`
	CodeEentreprise float64  `json:"code_entreprise"`
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

	facture := new(models.Facture)

	db.First(&facture, id)
	facture.NFacture = updateData.NFacture
	facture.CommandeID = updateData.CommandeID
	facture.Status = updateData.Status
	facture.DelaiPaiement = updateData.DelaiPaiement
	facture.Signature = updateData.Signature
	facture.PosID = updateData.PosID
	facture.CodeEentreprise = updateData.CodeEentreprise

	db.Save(&facture)

	return c.JSON(
		fiber.Map{
			"status":  "success",
			"message": "facture updated success",
			"data":    facture,
		},
	)

}

// Delete data
func DeleteFacture(c *fiber.Ctx) error {
	id := c.Params("id")

	db := database.DB

	var facture models.Facture
	db.First(&facture, id)
	if facture.NFacture == 0 {
		return c.Status(404).JSON(
			fiber.Map{
				"status":  "error",
				"message": "No facture name found",
				"data":    nil,
			},
		)
	}

	db.Delete(&facture)

	return c.JSON(
		fiber.Map{
			"status":  "success",
			"message": "facture deleted success",
			"data":    nil,
		},
	)
}
