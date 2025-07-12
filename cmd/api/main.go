package main

import (
	"log"
	"net/http"
	"os"
	"shop-search/internal/api"
	"shop-search/internal/cron"
	"shop-search/internal/database"

	"github.com/joho/godotenv"
	c "github.com/robfig/cron/v3"
	"github.com/rs/cors"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Printf("Warning: .env file not found")
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	MAIN_DB_URI := os.Getenv("MONGODB_URI")
	SRC_DB_URI := os.Getenv("MONGODB_URI_SRC")

	mainClient, err := database.Connect(MAIN_DB_URI)
	if err != nil {
		log.Fatal("Could not connect to main MongoDB:", err)
	}

	srcClient, err := database.Connect(SRC_DB_URI)
	if err != nil {
		log.Fatal("Could not connect to src MongoDB:", err)
	}

	database.SetMainClient(mainClient)
	database.SetSrcClient(srcClient)

	cronJob := c.New()
	cronJob.AddFunc("0 2 * * *", func() { // Every day at 2am
		log.Println("Starting sync job...")
		err := cron.SyncCollections(database.GetSrcDB(), database.GetMainDB(), []string{"products", "vendors", "categories", "tags"})
		if err != nil {
			log.Println("Sync error:", err)
		} else {
			log.Println("Sync completed.")
		}
	})
	cronJob.Start()

	router := api.NewRouter()

	// Create a CORS middleware
	c := cors.New(cors.Options{
		AllowedOrigins: []string{os.Getenv("ORIGIN")},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders: []string{"*"},
	})

	handler := c.Handler(router)

	log.Printf("Server starting on port %s...", port)
	log.Fatal(http.ListenAndServe(":"+port, handler))
}
