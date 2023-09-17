package payload

import "go.mongodb.org/mongo-driver/bson/primitive"

type InitialQuery struct {
	SessionNo *string `json:"sessionNo" validate:"required"`
}

type InitialResponse struct {
	SessionId *primitive.ObjectID `json:"sessionId"`
	Email     *string             `json:"email"`
	Pictures  []*Picture          `json:"pictures"`
}

type Picture struct {
	Id          *primitive.ObjectID `json:"id"`
	Src         *string             `json:"src"`
	Title       *string             `json:"title"`
	Description *string             `json:"description"`
}
