package mng

import (
	"github.com/kamva/mgm/v3"

	"backend/types/model"
)

var Config *mgm.Collection
var PairCommit *mgm.Collection
var PairLog *mgm.Collection
var PairMatch *mgm.Collection
var PairPicture *mgm.Collection
var Session *mgm.Collection

func initCollection() {
	Config = mgm.Coll(new(model.Config))
	PairCommit = mgm.Coll(new(model.PairCommit))
	PairLog = mgm.Coll(new(model.PairLog))
	PairMatch = mgm.Coll(new(model.PairMatch))
	PairPicture = mgm.Coll(new(model.PairPicture))
	Session = mgm.Coll(new(model.Session))
}
