package payload

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ParingCommit struct {
	SessionId *primitive.ObjectID `json:"sessionId" validate:"required"`
	ItemId    *string             `json:"itemId" validate:"required"`
}

type ParingCommitResponse struct {
	Matched     *bool   `json:"matched"`
	ForwardLink *string `json:"forwardLink"`
	PairedWith  *string `json:"pairedWith"`
	PairedDiff  *int64  `json:"pairedDiff"`
}
