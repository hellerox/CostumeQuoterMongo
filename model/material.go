package model

import "gopkg.in/mgo.v2/bson"

type Material struct {
	Id               bson.ObjectId `bson:"_id" json:"id"`
	Nombrematerial   string        `bson:"nombrematerial" json:"nombrematerial"`
	Idtipomaterial   int           `bson:"idtipomaterial" json:"idtipomaterial"`
	Idmedidamaterial int           `bson:"idmedidamaterial" json:"idmedidamaterial"`
	Costo            int           `bson:"costo" json:"costo"`
}
