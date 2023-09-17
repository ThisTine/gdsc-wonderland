package model

import (
	"time"

	"github.com/kamva/mgm/v3"
	"go.mongodb.org/mongo-driver/bson/primitive"

	mod "backend/modules"
	mh "backend/modules/mng/helper"
)

type PairMatch struct {
	mh.ModelBase `bson:"_,inline"`
	CommitA      *primitive.ObjectID `bson:"commitA,omitempty"`
	CommitB      *primitive.ObjectID `bson:"commitB,omitempty"`
	Diff         *time.Duration      `bson:"diff,omitempty"`
}

const PairMatchCommitA = "commitA"
const PairMatchCommitB = "commitB"
const PairMatchDiff = "diff"

func (r *PairMatch) Collection() *mgm.Collection {
	coll, _ := mh.CreateCollection(mod.Database, "pair_matches")

	return coll
}
