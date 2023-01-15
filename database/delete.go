package database

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
)

func Delete(name string) error {
	client, ctx := StartConnect()
	defer client.Disconnect(ctx)

	collection := client.Database("Anime").Collection("Catalog")
	filter := bson.M{"name": name}

	_, err := collection.DeleteOne(context.Background(), filter)

	if err != nil {
		panic(err)
	}

	return nil
}
