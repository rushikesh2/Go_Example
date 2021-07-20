package main

import (
	fury "Shield/captainfury"
	"Shield/config"
	"Shield/dao"
	"fmt"

)

var confi = config.Config{}
var da = dao.MongoDao{}
var f fury.CapMission

var (
	GetMissionDetails = fury.CheckMisionDetails
)

// Parse the configuration file 'config.toml', and establish a connection to DB
func init() {
	confi.Read()

	da.Server = confi.Server
	da.Database = confi.Database
	da.Connect()
}

func captainActions(number int) {
	switch number {
	case 1:
		fury.CheckMisionDetails()
		break
	case 2:
		fury.AssignMissionToAvengers()
		break
	case 3:
		fury.CheckMissionDetails()
		break
	case 4:
		fury.UpdateMissionStatus()
		break
	case 5:
		fury.ListAvengers()
	default:
		fmt.Println("Invalid Operation")
	}

}
func main() {
	var number int
	for  {
		fmt.Print("Enter the Option: ")
		fmt.Scanln(&number)
		captainActions(number)
	}
	
}
