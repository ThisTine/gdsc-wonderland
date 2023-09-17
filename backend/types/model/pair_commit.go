package model

import (
	"github.com/kamva/mgm/v3"

	mod "backend/modules"
	"backend/modules/mng/helper"
)

type PairCommit struct {
	mh.ModelBase `bson:"_,inline"`
	SessionId    *string `bson:"sessionId,omitempty"`
	ItemNo       *string `bson:"itemNo,omitempty"`
}

func (r *PairCommit) Collection() *mgm.Collection {
	coll, _ := mh.CreateCollection(mod.Database, "pair_commits")

	return coll
}
