package cron

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func SyncCollections(srcDB, dstDB *mongo.Database, collections []string) error {
	ctx := context.Background()
	for _, col := range collections {
		srcCol := srcDB.Collection(col)
		dstCol := dstDB.Collection(col)

		cursor, err := srcCol.Find(ctx, bson.M{})
		if err != nil {
			return err
		}
		var docs []bson.M
		if err := cursor.All(ctx, &docs); err != nil {
			return err
		}

		for _, doc := range docs {
			id := doc["_id"]
			_, err := dstCol.ReplaceOne(ctx, bson.M{"_id": id}, doc, options.Replace().SetUpsert(true))
			if err != nil {
				log.Printf("Upsert error for %s: %v", col, err)
			}
		}
	}
	return nil
}
