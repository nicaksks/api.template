package database

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func FindAll() ([]primitive.M, error) {
	client, ctx := StartConnect()
	defer client.Disconnect(ctx)

	collection := client.Database("Anime").Collection("Catalog")

	filter := options.Find().SetProjection(bson.M{"_id": false, "createdAt": false, "updatedAt": false}) //Hide _id, created_at, update_at
	cur, err := collection.Find(context.Background(), bson.D{}, filter)
	if err != nil {
		return nil, err
	}
	defer cur.Close(context.Background())

	var all []primitive.M
	for cur.Next(context.Background()) {
		var a primitive.M
		if err := cur.Decode(&a); err != nil {
			return nil, err
		}
		all = append(all, a)
	}

	return all, nil
}
