package main

import (
	"gestion-de-depense/backend/internal/handlers"
	"gestion-de-depense/backend/internal/models"
	"gestion-de-depense/backend/internal/repositories"
	"gestion-de-depense/backend/internal/routes"
	"gestion-de-depense/backend/internal/services"
	"log"

	"github.com/gofiber/fiber/v2"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"github.com/joho/godotenv"
)

func init() {
	// Chargement du fichier .env
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Erreur lors du chargement du fichier .env")
	}
}

func main() {
	// Charger .env
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Connexion DB
	db, err := gorm.Open(sqlite.Open("expenses.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database")
	}

	// Migration des modèles
	db.AutoMigrate(&models.Utilisateur{}, &models.Transaction{})

	// Initialisation des repositories
	userRepo := repositories.NewUserRepository(db)
	transactionRepo := repositories.NewTransactionRepository(db)

	// Initialisation des services
	userService := services.NewUserService(userRepo)
	transactionService := services.NewTransactionService(transactionRepo)

	// Initialisation des handlers
	userHandler := handlers.NewUserHandler(userService)
	transactionHandler := handlers.NewTransactionHandler(transactionService)

	// Création de l'application Fiber
	app := fiber.New()

	// Définition des routes
	routes.SetupRoutes(app, userHandler, transactionHandler)

	// Lancement du serveur
	log.Fatal(app.Listen(":8080"))
}
