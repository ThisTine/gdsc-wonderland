package mng

import (
	"github.com/kamva/mgm/v3"

	"backend/types/model"
)

var PairCommit *mgm.Collection
var PairLog *mgm.Collection

func initCollection() {
	PairCommit = mgm.Coll(new(model.PairCommit))
	PairLog = mgm.Coll(new(model.PairLog))
}
