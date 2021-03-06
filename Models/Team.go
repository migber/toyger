package models

import (
	"fmt"
	"github.com/satori/go.uuid"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"	
)

type Team struct {
	Id		string	`json:"id"`
	Name	string		`json:"name"`
	Manager Manager		`json:"manager"`
	Riders  []string	`json:"riders"`
}

type Manager struct {
	Id 			string	`json:"id"`
	Name 		string		`json:"name"`
	Surname 	string		`json:"surname"`
	Birthdate	string		`json:"birth_date"`
	Phone		string		`json:"phone"`
	Nationality	string		`json:"nationality"`
}

type Teams []Team

var TEAMS = "teams"

func CreateTeam(t Team, session *mgo.Session,
	dbName string, tableName string) (Team, error){

	var team Team
	defer session.Close() 

	t.Id = uuid.NewV4().String()
	t.Manager.Id = uuid.NewV4().String()
	team.Id = t.Id
	team.Name = t.Name
	team.Manager = t.Manager
	team.Riders = t.Riders
	
	collection := session.DB(dbName).C(tableName)
	var err error
	if err = collection.Insert(team); err != nil {
		fmt.Println(err)
	}
	return t, err
}

func GetTeamsList(session *mgo.Session, dbName string, tableName string) (Teams, error) {

	teams := Teams{}
	defer session.Close()
	c := session.DB(dbName).C(tableName)
	err := c.Find(nil).All(&teams)
	if err != nil {
		fmt.Println(err)
	}
	return teams, err
}

func GetTeam(id string, session *mgo.Session,
	dbName string, tableName string) (Team, error) {

	var t Team
	defer session.Close()
	collection := session.DB(dbName).C(tableName)
	err := collection.Find(bson.M{"id": id}).One(&t)
	if err != nil {
		fmt.Println(err)
	}

	return t, err
}

func UpdateTeam(uid string, t Team, session *mgo.Session,
	dbName string, tableName string) (Team, error) {

	var updateT Team

	defer session.Close()
	
	collection := session.DB(dbName).C(tableName)

	err := collection.Find(bson.M{"id": uid}).One(&updateT)
	if err != nil {
		fmt.Println(err)
		return updateT, err
	}

	updateT.Name = t.Name
	updateT.Manager.Name = t.Manager.Name
	updateT.Manager.Surname = t.Manager.Surname
	updateT.Manager.Birthdate = t.Manager.Birthdate
	updateT.Manager.Phone = t.Manager.Phone
	updateT.Manager.Nationality = t.Manager.Nationality
	updateT.Riders = t.Riders
	var err2 error
	if err2 = collection.Update(bson.M{"id": uid}, updateT); err != nil {
		fmt.Println(err)
	}
	
	return updateT, err2
}

func DeleteTeam(uid string, session *mgo.Session,
	dbName string, tableName string) error {

	defer session.Close()
	
	collection := session.DB(dbName).C(tableName)
	err := collection.Remove(bson.M{"id": uid})
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Object was deleted")	
	}

	return err
}

 func FindCyclist(riders []string, id string) int {

	index:= -1
	for ind, cyclist := range riders {
		if cyclist == id {
			index = ind
		}
	}
	return index
}

func InsertRider(id string, riderId string, session *mgo.Session,
	dbName string, tableName string) error {

	var updatedTeam Team
	defer session.Close()
	
	collection := session.DB(dbName).C(tableName)
	err := collection.Find(bson.M{"id": id}).One(&updatedTeam)
	if err != nil {
		return err
	}
	riders := updatedTeam.Riders
	newRiders := append(riders, riderId)
	updatedTeam.Riders = newRiders
	err = collection.Update(bson.M{"id": id}, updatedTeam)
	if err != nil {
		return err
	}
	return nil
}

func DeleteRider(id string, riderId string, session *mgo.Session,
	dbName string, tableName string) error{

	var updatedTeam Team
	var emptyRiders []string
	defer session.Close()
	
	collection := session.DB(dbName).C(tableName)
	err := collection.Find(bson.M{"id": id}).One(&updatedTeam)
	if err != nil {
		fmt.Println(err)
	}
	riders := updatedTeam.Riders
	index := FindCyclist(riders, riderId)
	if (len(riders) == 1 && index == 0){
		updatedTeam.Riders = emptyRiders
	} else {
		newRiders := append(riders[:index], riders[index+1:]...)
		updatedTeam.Riders = newRiders
	}
	newRiders := append(riders[:index], riders[index+1:]...)
	updatedTeam.Riders = newRiders
	if err := collection.Update(bson.M{"id": id}, updatedTeam); err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}