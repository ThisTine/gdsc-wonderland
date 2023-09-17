package model

import (
	"encoding/json"
	"os"

	"github.com/kamva/mgm/v3"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"

	mod "backend/modules"
	mh "backend/modules/mng/helper"
)

type PairPicture struct {
	mh.ModelBase `bson:"_,inline"`
	Filename     *string `bson:"filename,omitempty" json:"filename"`
	Visible      *bool   `bson:"visible,omitempty" json:"visible"`
	Title        *string `bson:"title,omitempty" json:"title"`
	Description  *string `bson:"description,omitempty" json:"description"`
}

const PairPictureFilename = "filename"
const PairPictureVisible = "visible"
const PairPictureTitle = "title"
const PairPictureDescription = "description"

func (r *PairPicture) Collection() *mgm.Collection {
	coll, _ := mh.CreateCollection(mod.Database, "pair_pictures")

	// * Query current data
	var pairPictures []*PairPicture
	if err := coll.SimpleFind(&pairPictures, bson.M{}); err != nil {
		logrus.WithError(err).Fatal("UNABLE TO QUERY PAIR PICTURES FILE")
	}

	// * Construct map
	pairPictureMap := make(map[string]*PairPicture)
	for _, pairPicture := range pairPictures {
		pairPictureMap[*pairPicture.Filename] = pairPicture
	}

	// * Read migration file
	if bytes, err := os.ReadFile("./resources/database/pair_pictures.json"); err != nil {
		logrus.WithError(err).Fatal("UNABLE TO READ PAIR PICTURES FILE")
	} else {
		pairPictures := make([]*PairPicture, 0)
		if err := json.Unmarshal(bytes, &pairPictures); err != nil {
			logrus.WithError(err).Fatal("UNABLE TO PARSE PAIR PICTURES FILE")
		}
		for _, pairPicture := range pairPictures {
			if picture, exist := pairPictureMap[*pairPicture.Filename]; !exist {
				if err := coll.Create(pairPicture); err != nil {
					logrus.WithError(err).WithField("filename", pairPicture.Filename).Fatal("UNABLE TO CREATE DEFAULT DATA FOR PAIR PICTURES FILE")
				}
			} else {
				if pairPicture.Title != picture.Title || pairPicture.Description != picture.Description {
					picture.Title = pairPicture.Title
					picture.Description = pairPicture.Description
					if err := coll.Update(picture); err != nil {
						logrus.WithError(err).WithField("filename", pairPicture.Filename).Fatal("UNABLE TO UPDATE DEFAULT DATA FOR PAIR PICTURES FILE")
					}
				}
			}
		}
	}
	return coll
}
