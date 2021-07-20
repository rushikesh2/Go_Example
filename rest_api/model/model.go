package model

import "gopkg.in/mgo.v2/bson"

// Employee ...
type Employee struct {
	ID       bson.ObjectId `bson:"_id,omitempty" json:"id,omitempty"`
	EmpID    string        `bson:"empid" json:"empid"`
	Name     string        `bson:"name" json:"name"`
	Email    string        `bson:"mail" json:"mail"`
	Phone    string        `bson:"phone" json:"phone"`
	Position string        `bson:"position" json:"position"`
	Practice string        `bson:"practice" json:"practice"`
	Status   string        `bson:"status" json:"status"`
}

type MoviesDAO struct {
	Server   string
	Database string
}
