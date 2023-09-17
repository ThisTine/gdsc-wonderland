package mh

import (
	"github.com/kamva/mgm/v3"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"

	"densomap-backend/utils/value"
)

var cachedCollectionList []string

func IsCollectionExist(database *mongo.Database, name string) bool {
	if cachedCollectionList == nil {
		if collections, err := database.ListCollectionNames(mgm.Ctx(), bson.M{}); err != nil {
			logrus.Fatal("UNABLE TO CHECK COLLECTION EXIST")
		} else {
			cachedCollectionList = collections
		}
	}

	return value.Contain[string](cachedCollectionList, name)
}
