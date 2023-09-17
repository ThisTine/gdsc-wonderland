package mh

import (
	"github.com/kamva/mgm/v3"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
)

func CreateCollection(database *mongo.Database, name string) (*mgm.Collection, bool) {
	exist := IsCollectionExist(database, name)
	_, _, database, err := mgm.DefaultConfigs()
	if err != nil {
		logrus.WithField("colName", name).Fatal("UNABLE TO LOAD COLLECTION")
	}

	if !exist {
		if err := database.CreateCollection(
			mgm.Ctx(),
			name,
		); err != nil {
			logrus.WithField("colName", name).WithField("e", err.Error()).Fatal("UNABLE TO CREATE COLLECTION")
		}
	}

	return mgm.CollectionByName(name), exist
}
