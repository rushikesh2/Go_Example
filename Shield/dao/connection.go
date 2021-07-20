package dao

import (
	"Shield/model"
	"fmt"
	"log"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type MongoDao struct {
	Database string
	Server   string
}

var db *mgo.Database

const (
	COLLECTION     = "employee"
	MISSION        = "missiondetails"
	AVENGERDETAILS = "AvengersDetails"
)

func (m *MongoDao) Connect() {
	session, err := mgo.Dial(m.Server)
	if err != nil {
		log.Fatal("Error while connecting to the mongo server: \n", err)
	}
	db = session.DB(m.Database)
	log.Print("Mongo Connection established successfully. \n")
}

func (m *MongoDao) Insert(avn model.MissionDetails) error {
	err := db.C(MISSION).Insert(&avn)
	return err
}

func (m *MongoDao) InsertRecord(avn model.Avengers_Details) error {
	err := db.C(AVENGERDETAILS).Insert(&avn)
	return err
}

func (m *MongoDao) FindAll() (mission []model.MissionStatus, err error) {
	err = db.C(MISSION).Find(bson.M{}).All(&mission)
	return mission, err
}

func (m *MongoDao) FindAvengers() (avengers []model.Avengers_Details, err error) {
	err = db.C(AVENGERDETAILS).Find(bson.M{}).All(&avengers)
	return avengers, err
}

func (m *MongoDao) CheckCountOfAvenger(name string) bool {
	var avenger []model.MissionDetails
	err := db.C(MISSION).Find(bson.M{"avengername": name}).All(&avenger)
	if err != nil {
		fmt.Print("No record found in database \n")
		return false
	}

	if len(avenger) >= 2 {
		return false
	}
	return true
}

// func (m *MongoDao) Count(avn []string) (n int,err error) {
// 	// var mission []model.Mission
// 	for _, v := range avn {
// 		selector := bson.M{"avengername": v}
// 		n, err = db.C(MISSION).Count(selector)
// 	}

// 	return n, err
// }

func (m *MongoDao) FindById(name string) (usr model.MissionStatus, err error) {
	err = db.C(MISSION).Find(bson.M{"mission_name": name}).One(&usr)
	return usr, err
}

// func (m *MongoDao) Delete(id string) error {
// 	err := db.C(COLLECTION).Remove(bson.M{"empid": id})
// 	return err
// }

// func (m *MongoDao) Update(employee model.Employee) error {
// 	fmt.Printf("%s", employee)
// 	err := db.C(COLLECTION).Update(bson.M{"empid": employee.EmpID}, employee)
// 	return err
// }

func (m *MongoDao) UpdateMission(name, status string) error {
	avenger_status := "On Mission"

	if status == "Completed" {
		avenger_status = "Available"
	}

	change := bson.M{"$set": bson.M{"status": status}}
	err := db.C(MISSION).Update(bson.M{"mission_name": name}, change)
	if err == nil {
		change := bson.M{"$set": bson.M{"status": avenger_status}}
		err = db.C(AVENGERDETAILS).Update(bson.M{"assignedMission": name}, change)
		if err != nil {
			return err
		}
	}
	return nil
}
