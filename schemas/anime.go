package schemas

import "time"

type Data struct {
	Name        string    `json:"name" bjson:"name"`
	Description string    `json:"description" bjson:"description"`
	Genre       string    `json:"genre" bjson:"genre"`
	CreatedAt   time.Time `json:"createdAt" bson:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt" bson:"updatedAt"`
}
