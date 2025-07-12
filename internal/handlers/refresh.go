package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"shop-search/internal/cron"
	"shop-search/internal/database"
)

func (h *Handler) RefreshHandler(w http.ResponseWriter, r *http.Request) {
	err := cron.SyncCollections(database.GetSrcDB(), database.GetMainDB(), []string{"products", "vendors", "categories", "tags"})
	if err != nil {
		log.Println("Sync error:", err)
	} else {
		log.Println("Sync completed.")
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]any{
		"message": "database sync successfully",
	})
}
