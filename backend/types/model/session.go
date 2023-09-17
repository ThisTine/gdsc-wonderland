package model

import (
	"github.com/kamva/mgm/v3"

	mod "backend/modules"
	mh "backend/modules/mng/helper"
)

type Session struct {
	mh.ModelBase `bson:"_,inline"`
	Email        *string `bson:"email,omitempty"`
	Hash         *string `bson:"hash,omitempty"`
	MatchedCount *int64  `bson:"matchedCount,omitempty"`
}

const SessionEmail = "email"
const SessionHash = "hash"
const SessionMatchedCount = "matchedCount"

func (r *Session) Collection() *mgm.Collection {
	coll, _ := mh.CreateCollection(mod.Database, "sessions")

	return coll
}
