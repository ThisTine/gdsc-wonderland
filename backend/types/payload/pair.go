package payload

import "go.mongodb.org/mongo-driver/bson/primitive"

type ParingCommit struct {
	SessionId *primitive.ObjectID `json:"sessionId" validate:"required"`
	ItemNo    *string             `json:"itemNo" validate:"required"`
}

type ParingCommitResponse struct {
	Matched     *bool   `json:"matched"`
	ForwardLink *string `json:"forwardLink"`
	PairedWith  *string `json:"pairedWith"`
}
