package database

import (
	"context"
	"shop-search/internal/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func FindProductsByName(ctx context.Context, regex string, limit int64) ([]models.Product, error) {
	var products []models.Product
	cursor, err := GetCollection("products").Find(ctx, bson.M{"name": bson.M{"$regex": regex, "$options": "i"}}, options.Find().SetLimit(limit))
	if err != nil {
		return nil, err
	}
	if err := cursor.All(ctx, &products); err != nil {
		return nil, err
	}
	return products, nil
}

func FindVendorsByName(ctx context.Context, regex string, limit int64) ([]models.Vendor, error) {
	var vendors []models.Vendor
	cursor, err := GetCollection("vendors").Find(ctx, bson.M{"name": bson.M{"$regex": regex, "$options": "i"}}, options.Find().SetLimit(limit))
	if err != nil {
		return nil, err
	}
	if err := cursor.All(ctx, &vendors); err != nil {
		return nil, err
	}
	return vendors, nil
}

func FindCategoriesByName(ctx context.Context, regex string, limit int64) ([]models.Category, error) {
	var categories []models.Category
	cursor, err := GetCollection("categories").Find(ctx, bson.M{"name": bson.M{"$regex": regex, "$options": "i"}}, options.Find().SetLimit(limit))
	if err != nil {
		return nil, err
	}
	if err := cursor.All(ctx, &categories); err != nil {
		return nil, err
	}
	return categories, nil
}

func FindTagsByName(ctx context.Context, regex string, limit int64) ([]models.Tag, error) {
	var tags []models.Tag
	cursor, err := GetCollection("tags").Find(ctx, bson.M{"name": bson.M{"$regex": regex, "$options": "i"}}, options.Find().SetLimit(limit))
	if err != nil {
		return nil, err
	}
	if err := cursor.All(ctx, &tags); err != nil {
		return nil, err
	}
	return tags, nil
}
