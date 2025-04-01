package main

import (
	"fmt"
	"gestion-de-depense/backend/internal/handlers"
	"gestion-de-depense/backend/internal/models"
	"gestion-de-depense/backend/internal/repositories"
	"gestion-de-depense/backend/internal/routes"
	"gestion-de-depense/backend/internal/services"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"github.com/joho/godotenv"
)

func main() {
	// Charger .env
	err := godotenv.Load("cmd/.env")
	if err != nil {
		log.Fatal("Erreur du Chargement du fichier .env")
	}

	fmt.Println("DB_PATH:", os.Getenv("DB_PATH"))

	// Connexion DB
	db, err := gorm.Open(sqlite.Open("expenses.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("Connexion à la base de données échouée")
	}

	// Migration des modèles
	db.AutoMigrate(&models.Utilisateur{}, &models.Transaction{})

	// Initialisation des repositories
	userRepo := repositories.Initialisation_Depot_Utilisateur(db)
	transactionRepo := repositories.Initialisation_Gestionnaire_Transaction(db)

	// Initialisation des services
	userService := services.Initilisation_Service_Utilisteur(userRepo)
	transactionService := services.Initialisation_Transaction_Service(transactionRepo)

	// Initialisation des handlers
	userHandler := handlers.NewGestion_Utilisateur(userService)
	Controleur_Transaction := handlers.Initialisation_Gest_Transaction(transactionService)

	// Création de l'application Fiber
	app := fiber.New()

	// Définition des routes
	routes.SetupRoutes(app, userHandler, Controleur_Transaction)

	// Lancement du serveur
	log.Fatal(app.Listen(":8080"))
}
