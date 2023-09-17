package model

import (
	"github.com/kamva/mgm/v3"
	"go.mongodb.org/mongo-driver/bson/primitive"

	mod "backend/modules"
	"backend/modules/mng/helper"
)

type PairCommit struct {
	mh.ModelBase `bson:"_,inline"`
	SessionNo    *string             `bson:"sessionNo,omitempty"`
	ItemNo       *string             `bson:"itemNo,omitempty"`
	PairedWith   *primitive.ObjectID `bson:"pairedWith,omitempty"`
}

func (r *PairCommit) Collection() *mgm.Collection {
	coll, _ := mh.CreateCollection(mod.Database, "pair_commits")

	return coll
}
