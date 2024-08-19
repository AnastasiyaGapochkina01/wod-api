package models

import "gopkg.in/mgo.v2/bson"

// Represents a movie, we uses bson keyword to tell the mgo driver how to name
// the properties in mongodb document
type Wod struct {
	ID          bson.ObjectId `bson:"_id" json:"id"`
	Title        string        `bson:"title" json:"title"`
	Desc  string        `bson:"title" json:"desc"`
	Content string        `bson:"content" json:"content"`
}
