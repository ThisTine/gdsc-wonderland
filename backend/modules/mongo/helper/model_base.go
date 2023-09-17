package mh

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type ModelBase struct {
	ID        *primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	CreatedAt *time.Time          `json:"createdAt,omitempty" bson:"created_at,omitempty"`
	UpdatedAt *time.Time          `json:"updatedAt,omitempty" bson:"updated_at,omitempty"`
}

func (r *ModelBase) PrepareID(id any) (any, error) {
	if idStr, ok := id.(string); ok {
		return primitive.ObjectIDFromHex(idStr)
	}

	return id, nil
}

// GetID method returns a model's ID
func (r *ModelBase) GetID() interface{} {
	return r.ID
}

// SetID sets the value of a model's ID field.
func (r *ModelBase) SetID(id interface{}) {
	objectID := id.(primitive.ObjectID)
	r.ID = &objectID
}

// Creating hook is used here to set the `created_at` field
// value when inserting a new model into the database.
func (r *ModelBase) Creating() error {
	utc := time.Now().UTC()
	r.CreatedAt = &utc
	return nil
}

// Saving hook is used here to set the `updated_at` field
// value when creating or updating a model.
func (r *ModelBase) Saving() error {
	utc := time.Now().UTC()
	r.UpdatedAt = &utc
	return nil
}
