package mng

import (
	"github.com/kamva/mgm/v3"

	"backend/types/model"
)

var Config *mgm.Collection
var PairCommit *mgm.Collection
var PairLog *mgm.Collection

func initCollection() {
	Config = mgm.Coll(new(model.Config))
	PairCommit = mgm.Coll(new(model.PairCommit))
	PairLog = mgm.Coll(new(model.PairLog))
}
