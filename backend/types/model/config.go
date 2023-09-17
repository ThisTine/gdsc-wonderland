package model

import (
	"github.com/kamva/mgm/v3"
	"github.com/sirupsen/logrus"

	mod "backend/modules"
	mh "backend/modules/mng/helper"
)

type Config struct {
	mh.ModelBase `bson:"_,inline"`
	Secret       *string `bson:"secret,omitempty"`
}

func (r *Config) Collection() *mgm.Collection {
	coll, exist := mh.CreateCollection(mod.Database, "configs")

	if !exist {
		config := &Config{
			Secret: &mod.Conf.Secret,
		}
		if err := coll.Create(config); err != nil {
			logrus.WithError(err).Fatal("UNABLE TO CREATE DEFAULT DATA FOR CONFIG")
		}
	}

	return coll
}
