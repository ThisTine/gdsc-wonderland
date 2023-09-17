package mongo

import (
	"github.com/kamva/mgm/v3"

	"backend/types/model"
)

var PairCommit *mgm.Collection

func initCollection() {
	PairCommit = mgm.Coll(new(model.PairCommit))
}
