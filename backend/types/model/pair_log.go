package model

import (
	"github.com/kamva/mgm/v3"

	"backend/modules"
	"backend/modules/mng/helper"
)

type PairLog struct {
	mh.ModelBase `bson:"_,inline"`
	SessionNo    *string        `bson:"sessionNo,omitempty"`
	Action       *string        `bson:"action,omitempty"`
	Attribute    map[string]any `bson:"attribute,omitempty"`
}

func (r *PairLog) Collection() *mgm.Collection {
	coll, _ := mh.CreateCollection(mod.Database, "pair_logs")

	return coll
}
