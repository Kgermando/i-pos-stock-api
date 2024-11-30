package main

import ( 
	"log" 
	"os"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/encryptcookie"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/kgermando/i-pos-stock/database"
	"github.com/kgermando/i-pos-stock/routes"
)

func GetPort() string {
	port := os.Getenv("PORT")
	if port == "" {
		port = ":3000"
	} else {
		port = ":" + port
	}

	return port
}

func main() {

	database.Connect()

	app := fiber.New()

	// Initialize default config
	app.Use(logger.New())

	// Configuration du middleware EncryptCookie
	app.Use(encryptcookie.New(encryptcookie.Config{
		Key: "3BX/f5BIiRqs7zClDx0ODmMKX3+6sV33L21vUhCTg+8=",
	}))

	// Middleware
	app.Use(cors.New(cors.Config{
		AllowOrigins:     "https://i-stock-pos.onrender.com, http://localhost:4200, http://192.168.1.81:4200, http://192.168.100.143:4200",
		AllowHeaders:     "Origin, Content-Type, Accept",
		AllowCredentials: true,
		AllowMethods: strings.Join([]string{
			fiber.MethodGet,
			fiber.MethodPost,
			fiber.MethodHead,
			fiber.MethodPut,
			fiber.MethodDelete,
			fiber.MethodPatch,
		}, ","),
	}))

	routes.Setup(app)  

	// Chargement des certificats SSL
	// cert, err := tls.LoadX509KeyPair("192.168.1.81.pem", "192.168.1.81-key.pem")
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// cert, err := tls.LoadX509KeyPair("192.168.100.143.pem", "192.168.100.143-key.pem")
	// if err != nil {
	// 	log.Fatal(err)
	// }
 

	// // Configuration TLS
	// tlsConfig := &tls.Config{
	// 	Certificates: []tls.Certificate{cert},
	// }

	// // linkIp := fmt.Sprintf("https://192.168.1.81:%s", getPort())

	// // Démarrage du serveur HTTPS
	// listener, err := tls.Listen("tcp", "192.168.100.143:3000", tlsConfig)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// log.Fatal(app.Listener(listener))

	log.Fatal(app.Listen(GetPort()))

}
