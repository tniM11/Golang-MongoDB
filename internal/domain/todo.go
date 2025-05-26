package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

// Todo represents a task in the system.
type Todo struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Title       string             `bson:"title" json:"title"`
	Description string             `bson:"description" json:"description"`
}
