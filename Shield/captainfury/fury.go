package captainfury

import (
	"Shield/dao"
	"Shield/model"
	"bufio"
	//"encoding/json"
	"fmt"
	"os"
	// "strings"
	"errors"

	"github.com/tatsushid/go-prettytable"
	mgo "gopkg.in/mgo.v2"

)

var da = dao.MongoDao{}
var db *mgo.Database


type CapMission interface {
	CheckMisionDetails() 
	AssignMissionToAvengers()
	CheckMissionDetails()
	UpdateMissionStatus()
	ListAvengers()
}

// CheckMisionDetails get all mission's  Details
func CheckMisionDetails() {
	var mission []model.MissionStatus
	mission, err := da.FindAll()
	if err != nil {
		fmt.Print("Error while querying database for reading all mission \n:", err)
		return 
	}

	if len(mission) < 1 {
		fmt.Print("No Mission has been assigned to the Avengeres \n")
		return
	}
	tbl, err := prettytable.NewTable([]prettytable.Column{
		{Header: "Mission Name", MaxWidth: 100},
		{Header: "Status", MinWidth: 10},
		{Header: "Avengers"},
	}...)
	if err != nil {
		panic(err)
	}
	tbl.Separator = " | "
	for _, v := range mission {
		tbl.AddRow(v.Mission_Name, v.Status, v.AvengersName)
	}
	tbl.Print()
	return
}

// AssignMissionToAvengers Assign missin to avengers
func AssignMissionToAvengers() {
	var avenger model.MissionDetails
	var details model.Avengers_Details
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Printf("Enter Mission Name: ")
	scanner.Scan()
	avenger.Mission_Name = scanner.Text()

	fmt.Printf("Enter Avenger: ")
	scanner.Scan()
	avenger.AvengersName = scanner.Text()

	status := da.CheckCountOfAvenger(avenger.AvengersName)
	if status == false {
		fmt.Printf("Sorry, %s has already been working on two missions. \n", avenger.AvengersName)
		return
	}
	// // var arr = make([]string, 2)
	// for i:=0; i<len(a); i++ {
	//    fmt.Scanf("%s", &a[i])
	//    avenger.AvengersName = a
	//    //checkCount(avenger.AvengersName)
	// }

	fmt.Printf("Enter Mission Details: ")
	scanner.Scan()
	avenger.Mission_Detail = scanner.Text()
	avenger.Status = "Assigned"

	err := da.Insert(avenger)
	if err != nil {
		fmt.Print("Unable to insert record into database \n")
		return
	}
	details.AssignedMission = avenger.Mission_Name
	details.Status = "On Mission"
	details.AvengersName = avenger.AvengersName

	err = da.InsertRecord(details)
	if err != nil {
		fmt.Print("Unable to insert record into database \n")
		return
	}

	fmt.Printf("Mission has been assigned to the Avengeres %s \n", avenger.AvengersName)
	fmt.Println("Email notification sent and/or SMS notification sent. \n")
}

// CheckMissionDetails get specific mission details
func CheckMissionDetails() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Printf("Enter Mission Name: ")
	scanner.Scan()
	name := scanner.Text()

	mission, err := da.FindById(name)
	if err != nil {
		fmt.Print("unable to find record \n")
		return
	}
	// jsonF, err := json.Marshal(mission)
	// if err != nil {
	// 	fmt.Print("unable to find record \n")
	// 	return
	// }

	// s := string(jsonF)
	// t := strings.Replace(s, "{", " ", -1)
	fmt.Printf("%+v\n", mission)
}

// UpdateMissionStatus Update mission status complete or Assigned
func UpdateMissionStatus() {

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Printf("Enter Mission Name: ")
	scanner.Scan()
	name := scanner.Text()

	fmt.Printf("Enter New Status: ")
	scanner.Scan()
	status := scanner.Text()
	if status != "Completed" && status != "Assigned" {
		fmt.Print("Please provide valid status... \n")
		return
	}

	err := da.UpdateMission(name, status)
	if err != nil {
		fmt.Print("unable to find record \n")
		return
	}
	fmt.Printf("Email has been sent to the avenger")
}

// ListAvengers will list all the avengers with their assinged mission
func ListAvengers() {
	avengers, err := da.FindAvengers()
	if err != nil {
		fmt.Print("Error while querying database for reading all mission \n:", err)
		return
	}
	tbl, err := prettytable.NewTable([]prettytable.Column{
		{Header: "Avenger Name", MaxWidth: 100},
		{Header: "Status", MinWidth: 10},
		{Header: "Assigned Mission"},
	}...)
	if err != nil {
		panic(err)
	}
	tbl.Separator = " | "
	for _, v := range avengers {
		tbl.AddRow(v.AvengersName, v.Status, v.AssignedMission)
	}
	tbl.Print()

}

// func checkCount(avenger []string){
// 	n, err := da.FindById(avenger)
// 	if err != nil {
// 		fmt.Print("Unable to find record into database")
// 		return
// 	}
// 	if len(n) > 2 {
// 		fmt.Print("Sorry, %s has already been working on two missions.", avenger[0])
// 		return
// 	}
// }
