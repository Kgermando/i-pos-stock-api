package commande

import (
	"fmt"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/kgermando/i-pos-stock/database"
	"github.com/kgermando/i-pos-stock/models"
)

// Paginate
func GetPaginatedCommandeEntreprise(c *fiber.Ctx) error {
	db := database.DB
	codeEntreprise := c.Params("code_entreprise")

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

	var dataList []models.Commande

	var length int64
	db.Model(dataList).Where("code_entreprise = ?", codeEntreprise).Count(&length)
	db.Where("code_entreprise = ?", codeEntreprise). 
		Where("ncommande::text ILIKE ? OR status ILIKE ?", "%"+search+"%", "%"+search+"%").
		Offset(offset).
		Limit(limit).
		Order("commandes.updated_at DESC").
		Preload("CommandeLines").
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
		"message":    "All commandes",
		"data":       dataList,
		"pagination": pagination,
	})
}

// Paginate
func GetPaginatedCommande(c *fiber.Ctx) error {
	db := database.DB
	codeEntreprise := c.Params("code_entreprise")
	posId := c.Params("pos_id")

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

	var dataList []models.Commande

	var length int64
	db.Model(dataList).Where("code_entreprise = ?", codeEntreprise).
	Where("pos_id = ?", posId).Count(&length)
	db.Where("code_entreprise = ?", codeEntreprise).
		Where("pos_id = ?", posId).
		Where("ncommande ILIKE ?", "%"+search+"%").
		Offset(offset).
		Limit(limit).
		Order("commandes.updated_at DESC").
		Preload("CommandeLines").
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
		"message":    "All commandes",
		"data":       dataList,
		"pagination": pagination,
	})
}

// Get All data
func GetAllCommandes(c *fiber.Ctx) error {
	db := database.DB
	codeEntreprise := c.Params("code_entreprise")
	posId := c.Params("pos_id")

	var data []models.Commande
	db.Where("code_entreprise = ?", codeEntreprise).
		Where("pos_id = ?", posId).
		Preload("CommandeLines").Find(&data)
	return c.JSON(fiber.Map{
		"status":  "success",
		"message": "All commandes",
		"data":    data,
	})
}

// Get one data
func GetCommande(c *fiber.Ctx) error {
	id := c.Params("id")
	db := database.DB
	var commande models.Commande
	db.Preload("CommandeLines").Find(&commande, id)
	if commande.Ncommande == 0 {
		return c.Status(404).JSON(
			fiber.Map{
				"status":  "error",
				"message": "No commande name found",
				"data":    nil,
			},
		)
	}
	return c.JSON(
		fiber.Map{
			"status":  "success",
			"message": "commande found",
			"data":    commande,
		},
	)
}

// Create data
func CreateCommande(c *fiber.Ctx) error {
	p := &models.Commande{}

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
func UpdateCommande(c *fiber.Ctx) error {
	id := c.Params("id")
	db := database.DB

	type UpdateData struct {
		PosID          uint   `json:"pos_id"`
		Ncommande      uint64 `json:"ncommande"` // Number Random
		Status         string `json:"status"`     // Ouverte et Fermée
		ClientID       uint   `json:"client_id"`
		Signature      string `json:"signature"`
		CodeEntreprise uint   `json:"code_entreprise"`
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

	commande := new(models.Commande)

	db.First(&commande, id)
	commande.PosID = updateData.PosID
	commande.Ncommande = updateData.Ncommande
	commande.Status = updateData.Status
	commande.ClientID = updateData.ClientID
	commande.Signature = updateData.Signature
	commande.CodeEntreprise = updateData.CodeEntreprise

	db.Save(&commande)

	return c.JSON(
		fiber.Map{
			"status":  "success",
			"message": "commande updated success",
			"data":    commande,
		},
	)

}

// Delete data
func DeleteCommande(c *fiber.Ctx) error {
	id := c.Params("id")

	db := database.DB

	var commande models.Commande
	db.First(&commande, id)
	if commande.Ncommande == 0 {
		return c.Status(404).JSON(
			fiber.Map{
				"status":  "error",
				"message": "No commande name found",
				"data":    nil,
			},
		)
	}

	db.Delete(&commande)

	return c.JSON(
		fiber.Map{
			"status":  "success",
			"message": "commande deleted success",
			"data":    nil,
		},
	)
}
