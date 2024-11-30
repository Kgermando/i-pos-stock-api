package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/kgermando/i-pos-stock/controllers/auth"
	"github.com/kgermando/i-pos-stock/controllers/boncommande"
	"github.com/kgermando/i-pos-stock/controllers/commande"
	"github.com/kgermando/i-pos-stock/controllers/contact"
	"github.com/kgermando/i-pos-stock/controllers/entreprise"
	"github.com/kgermando/i-pos-stock/controllers/fournisseurclient"
	"github.com/kgermando/i-pos-stock/controllers/pos"
	"github.com/kgermando/i-pos-stock/controllers/product"
	"github.com/kgermando/i-pos-stock/controllers/stock"
	"github.com/kgermando/i-pos-stock/controllers/users"
	"github.com/kgermando/i-pos-stock/middlewares"
)

func Setup(app *fiber.App) {

	api := app.Group("/api", logger.New())

	// Authentification controller
	au := api.Group("/auth")
	au.Post("/register", auth.Register)
	au.Post("/login", auth.Login)
	au.Post("/forgot-password", auth.Forgot)
	au.Post("/reset/:token", auth.ResetPassword)

	app.Use(middlewares.IsAuthenticated)

	au.Get("/user", auth.AuthUser)
	au.Put("/profil/info", auth.UpdateInfo)
	au.Put("/change-password", auth.ChangePassword)
	au.Post("/logout", auth.Logout)

	// User controller
	u := api.Group("/users")
	u.Get("/all", users.GetAllUsers)
	u.Get("/all/paginate", users.GetPaginatedUsers) 
	u.Get("/all/paginate/:user_id", users.GetPaginatedUserByID)
	u.Get("/all/:id", users.GetUserByID)
	u.Get("/all/count/:entreprise_id", users.GetUserByIDCount)
	u.Get("/get/:id", users.GetUser)
	u.Post("/create", users.CreateUser) 
	u.Put("/update/:id", users.UpdateUser)
	u.Delete("/delete/:id", users.DeleteUser) 

	// Entreprise controller
	e := api.Group("/entreprises")
	e.Get("/all", entreprise.GetAllEntreprises)
	e.Get("/all/paginate", entreprise.GetPaginatedEntreprise) 
	e.Get("/get/:id", entreprise.GetEntreprise)
	e.Post("/create", entreprise.CreateEntreprise) 
	e.Put("/update/:id", entreprise.UpdateEntreprise)
	e.Delete("/delete/:id", entreprise.DeleteEntreprise)

	// POS controller
	p := api.Group("/pos")
	p.Get("/all", pos.GetAllPoss)
	p.Get("/all/paginate", pos.GetPaginatedPos) 
	p.Get("/get/:id", pos.GetPos)
	p.Post("/create", pos.CreatePos) 
	p.Put("/update/:id", pos.UpdatePos)
	p.Delete("/delete/:id", pos.DeletePos)

	// Category controller
	c := api.Group("/categories")
	c.Get("/all", product.GetAllCategorys)
	c.Get("/all/paginate", product.GetPaginatedCategory) 
	c.Get("/get/:id", product.GetCategory)
	c.Post("/create", product.CreateCategory) 
	c.Put("/update/:id", product.UpdateCategory)
	c.Delete("/delete/:id", product.DeleteCategory)

	// Product controller
	pr := api.Group("/products")
	pr.Get("/all", product.GetAllProducts)
	pr.Get("/all/paginate", product.GetPaginatedProduct) 
	pr.Get("/get/:id", product.GetProduct)
	pr.Post("/create", product.CreateProduct) 
	pr.Put("/update/:id", product.UpdateProduct)
	pr.Delete("/delete/:id", product.DeleteProduct)

	// Stock controller
	s := api.Group("/stocks")
	s.Get("/all", stock.GetAllStocks)
	s.Get("/all/paginate", stock.GetPaginatedStock) 
	s.Get("/get/:id", stock.GetStock)
	s.Post("/create", stock.CreateStock) 
	s.Put("/update/:id", stock.UpdateStock)
	s.Delete("/delete/:id", stock.DeleteStock)

	// Unite of sale controller
	uv := api.Group("/unite-ventes")
	uv.Get("/all", stock.GetAllUniteVentes)
	uv.Get("/all/paginate", stock.GetPaginatedUniteVente) 
	uv.Get("/get/:id", stock.GetUniteVente)
	uv.Post("/create", stock.CreateUniteVente) 
	uv.Put("/update/:id", stock.UpdateUniteVente)
	uv.Delete("/delete/:id", stock.DeleteUniteVente)

	// Bon of commande controller
	bc := api.Group("/bon-commandes")
	bc.Get("/all", boncommande.GetAllBonCommandes)
	bc.Get("/all/paginate", boncommande.GetPaginatedBonCommande) 
	bc.Get("/get/:id", boncommande.GetBonCommande)
	bc.Post("/create", boncommande.CreateBonCommande) 
	bc.Put("/update/:id", boncommande.UpdateBonCommande)
	bc.Delete("/delete/:id", boncommande.DeleteBonCommande)

	// Bon of commande line controller
	bcl := api.Group("/bon-commandes-lines")
	bcl.Get("/all", boncommande.GetAllBonCommandeLines)
	bcl.Get("/all/paginate", boncommande.GetPaginatedBonCommandeLine)
	bcl.Get("/all/paginate/:bon_commande_id", boncommande.GetPaginatedBonCommandeLineByID)
	bcl.Get("/get/:id", boncommande.GetBonCommandeLine)
	bcl.Post("/create", boncommande.CreateBonCommandeLine) 
	bcl.Put("/update/:id", boncommande.UpdateBonCommandeLine)
	bcl.Delete("/delete/:id", boncommande.DeleteBonCommandeLine)

	// Commande controller
	cmd := api.Group("/commandes")
	cmd.Get("/all", commande.GetAllCommandes)
	cmd.Get("/all/paginate", commande.GetPaginatedCommande) 
	cmd.Get("/get/:id", commande.GetCommande)
	cmd.Post("/create", commande.CreateCommande) 
	cmd.Put("/update/:id", commande.UpdateCommande)
	cmd.Delete("/delete/:id", commande.DeleteCommande)

	// Commande line controller
	cmdl := api.Group("/commandes-lines")
	cmdl.Get("/all", commande.GetAllCommandeLines)
	cmdl.Get("/all/paginate", commande.GetPaginatedCommandeLine)
	cmdl.Get("/all/paginate/:commande_id", commande.GetPaginatedCommandeLineByID)
	cmdl.Get("/get/:id", commande.GetCommandeLine)
	cmdl.Post("/create", commande.CreateCommandeLine) 
	cmdl.Put("/update/:id", commande.UpdateCommandeLine)
	cmdl.Delete("/delete/:id", commande.DeleteCommandeLine)

	// Client controller
	cl := api.Group("/clients")
	cl.Get("/all", fournisseurclient.GetAllClients)
	cl.Get("/all/paginate", fournisseurclient.GetPaginatedClient) 
	cl.Get("/get/:id", fournisseurclient.GetClient)
	cl.Post("/create", fournisseurclient.CreateClient) 
	cl.Put("/update/:id", fournisseurclient.UpdateClient)
	cl.Delete("/delete/:id", fournisseurclient.DeleteClient)

	// Fournisseur controller
	fs := api.Group("/fournisseurs")
	fs.Get("/all", fournisseurclient.GetAllFournisseurs)
	fs.Get("/all/paginate", fournisseurclient.GetPaginatedFournisseur) 
	fs.Get("/get/:id", fournisseurclient.GetFournisseur)
	fs.Post("/create", fournisseurclient.CreateFournisseur) 
	fs.Put("/update/:id", fournisseurclient.UpdateFournisseur)
	fs.Delete("/delete/:id", fournisseurclient.DeleteFournisseur)

	// Contact controller
	ctc := api.Group("/contacts")
	ctc.Get("/all", contact.GetAllContacts)
	ctc.Get("/all/paginate", contact.GetPaginatedContact) 
	ctc.Get("/get/:id", contact.GetContact)
	ctc.Post("/create", contact.CreateContact) 
	ctc.Put("/update/:id", contact.UpdateContact)
	ctc.Delete("/delete/:id", contact.DeleteContact)
	
}
