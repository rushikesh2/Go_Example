package model

import "gopkg.in/mgo.v2/bson"

// Avengers ...
type MissionDetails struct {
	ID             bson.ObjectId `bson:"_id,omitempty" json:"id,omitempty"`
	AvengersName   string      `bson:"avengername" json:"avengername"`
	Mission_Name   string        `bson:"mission_name" json:"mission_name"`
	Mission_Detail string        `bson:"mission_detail" json:"mission_detail"`
	Status         string        `bson:"status" json:"status"`
}

// Avengers_Details stores Avengers information
type Avengers_Details struct {
	ID                bson.ObjectId `bson:"_id,omitempty" json:"id,omitempty"`
	AvengersName       string        `bson:"avengersName" json:"avengersName"`
	AssignedMission  string        `bson:"assignedMission" json:"assignedMission"`
	Status         string        `bson:"status" json:"status"`
}

// Mission ...
type MissionStatus struct {
	Mission_Name string   `bson:"mission_name" json:"mission_name"`
	Status       string   `bson:"status" json:"status"`
	AvengersName string `bson:"avengername" json:"avengername"`
}


