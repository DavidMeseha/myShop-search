package handlers

import (
	"context"
	"encoding/json"
	"net/http"
	"shop-search/internal/database"
	"shop-search/internal/utils"
	"time"
)

type SearchOptionsBody struct {
	SearchText string `json:"searchText"`
	Categories bool   `json:"categories"`
	Vendors    bool   `json:"vendors"`
	Tags       bool   `json:"tags"`
	Products   bool   `json:"products"`
}

func (h *Handler) FindInAll(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), 10*time.Second)
	defer cancel()

	var body SearchOptionsBody
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil || body.SearchText == "" {
		http.Error(w, "Search text is required", http.StatusBadRequest)
		return
	}

	query := body.SearchText
	toleranceCount := 1
	if len(query) > 0 {
		toleranceCount = int(float64(len(query))*0.2 + 0.5)
	}
	var queryRegex string
	if len(query) >= 4 {
		queryRegex = utils.GenerateSearchTextVariants(query, toleranceCount)
	} else {
		queryRegex = query
	}
	queryRegex += "|" + query + ".."

	enabledOptions := 0
	if body.Products {
		enabledOptions++
	}
	if body.Vendors {
		enabledOptions++
	}
	if body.Categories {
		enabledOptions++
	}
	if body.Tags {
		enabledOptions++
	}
	if enabledOptions == 0 {
		http.Error(w, "No search options enabled", http.StatusBadRequest)
		return
	}
	limit := int64(8 / enabledOptions) // Assume Handler has a *mongo.Database field

	type Item struct {
		Item interface{} `json:"item"`
		Type string      `json:"type"`
	}
	var items []Item

	if body.Products {
		products, _ := database.FindProductsByName(ctx, queryRegex, limit)
		for _, p := range products {
			items = append(items, Item{Item: p, Type: "product"})
		}
	}
	if body.Vendors {
		vendors, _ := database.FindVendorsByName(ctx, queryRegex, limit)
		for _, v := range vendors {
			items = append(items, Item{Item: v, Type: "vendor"})
		}
	}
	if body.Categories {
		categories, _ := database.FindCategoriesByName(ctx, queryRegex, limit)
		for _, c := range categories {
			items = append(items, Item{Item: c, Type: "category"})
		}
	}
	if body.Tags {
		tags, _ := database.FindTagsByName(ctx, queryRegex, limit)
		for _, t := range tags {
			items = append(items, Item{Item: t, Type: "tag"})
		}
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(items)
}
