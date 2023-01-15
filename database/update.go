package database

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

func Update(name string, description string, genre string, updatedat time.Time) error {
	client, ctx := StartConnect()
	defer client.Disconnect(ctx)

	collection := client.Database("Anime").Collection("Catalog")
	filter := bson.M{"name": name}
	update := bson.M{"$set": bson.M{"description": description, "genre": genre, "updatedAt": updatedat}}

	_, err := collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		panic(err)
	}

	return nil
}
