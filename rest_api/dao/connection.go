package dao

import (
	"fmt"
	"log"
	"rest_api/model"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type MongoDao struct {
	Database string
	Server   string
}

var db *mgo.Database

const (
	COLLECTION = "employee"
)

func (m *MongoDao) Connect() {
	session, err := mgo.Dial(m.Server)
	if err != nil {
		log.Fatal("Error while connecting to the mongo server:", err)
	}
	db = session.DB(m.Database)
}

func (m *MongoDao) Insert(usr model.Employee) error {
	err := db.C(COLLECTION).Insert(&usr)
	return err
}

func (m *MongoDao) FindAll() ([]model.Employee, error) {
	var usr []model.Employee
	err := db.C(COLLECTION).Find(bson.M{}).All(&usr)
	return usr, err
}

func (m *MongoDao) FindById(id string) (usr model.Employee, err error) {
	err = db.C(COLLECTION).Find(bson.M{"empid": id}).One(&usr)
	return usr, err
}

func (m *MongoDao) Delete(id string) error {
	err := db.C(COLLECTION).Remove(bson.M{"empid": id})
	return err
}

func (m *MongoDao) Update(employee model.Employee) error {
	fmt.Printf("%s", employee)
	err := db.C(COLLECTION).Update(bson.M{"empid": employee.EmpID}, employee)
	return err
}

func (m *MongoDao) UpdateEmp(employee model.Employee) error {
	fmt.Printf("%s", employee)
	_, err := db.C(COLLECTION).Upsert(bson.M{"empid": employee.EmpID}, employee)
	return err
}
