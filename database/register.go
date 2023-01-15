package database

import (
	"api/schemas"
	"context"
	"time"
)

func Register(name string, description string, genre string, createdAt time.Time, updatedAt time.Time) error {
	client, ctx := StartConnect()
	defer client.Disconnect(ctx)

	collection := client.Database("Anime").Collection("Catalog")
	filter := schemas.Data{Name: name, Description: description, Genre: genre, CreatedAt: createdAt, UpdatedAt: updatedAt}

	_, err := collection.InsertOne(context.Background(), filter)
	if err != nil {
		panic(err)
	}

	return err
}
