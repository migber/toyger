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

func CreateTeam(t Team, session *mgo.Session) Team {

	var team Team
	defer session.Close() 

	t.Id = uuid.NewV4().String()
	t.Manager.Id = uuid.NewV4().String()
	team.Id = t.Id
	team.Name = t.Name
	team.Manager = t.Manager
	team.Riders = t.Riders
	
	collection := session.DB(DATABASE).C(TEAMS)
	if err := collection.Insert(team); err != nil {
		panic(err)
	}
	return t
}

func GetTeamsList(session *mgo.Session) Teams {

	teams := Teams{}
	defer session.Close()
	c := session.DB(DATABASE).C(TEAMS)
	err := c.Find(nil).All(&teams)
	if err != nil {
		panic(err)
	}
	return teams
}

func GetTeam(id string, session *mgo.Session) Team {

	var t Team
	defer session.Close()
	collection := session.DB(DATABASE).C(TEAMS)
	err := collection.Find(bson.M{"id": id}).One(&t)
	if err != nil {
		panic(err)
	}
	fmt.Println(t)

	return t
}

func UpdateTeam(uid string, t Team, session *mgo.Session) Team {

	var updateT Team

	defer session.Close()
	
	collection := session.DB(DATABASE).C(TEAMS)

	err := collection.Find(bson.M{"id": uid}).One(&updateT)
	if err != nil {
		panic(err)
	}

	updateT.Name = t.Name
	updateT.Manager.Name = t.Manager.Name
	updateT.Manager.Surname = t.Manager.Surname
	updateT.Manager.Birthdate = t.Manager.Birthdate
	updateT.Manager.Phone = t.Manager.Phone
	updateT.Manager.Nationality = t.Manager.Nationality
	updateT.Riders = t.Riders

	if err := collection.Update(bson.M{"id": uid}, updateT); err != nil {
		panic(err)
	}
	
	return updateT
}

func DeleteTeam(uid string, session *mgo.Session) error {

	defer session.Close()
	
	collection := session.DB(DATABASE).C(TEAMS)
	err := collection.Remove(bson.M{"id": uid})
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Object was deleted")

	return err
}

 func FindCyclist(riders []string, id string) int {

	var index int
	for ind, cyclist := range riders {
		if cyclist == id {
			index = ind
		}
	}
	return index
}

func InsertRider(id string, riderId string, session *mgo.Session) error {

	var updatedTeam Team
	defer session.Close()
	
	collection := session.DB(DATABASE).C(TEAMS)
	err := collection.Find(bson.M{"id": id}).One(&updatedTeam)
	if err != nil {
		return err
	}
	riders := updatedTeam.Riders
	newRiders := append(riders, riderId)
	updatedTeam.Riders = newRiders
	if err := collection.Update(bson.M{"id": id}, updatedTeam); err != nil {
		return err
	}
	return nil
}

func DeleteRider(id string, riderId string, session *mgo.Session) error{

	var updatedTeam Team
	var emptyRiders []string
	defer session.Close()
	
	collection := session.DB(DATABASE).C(TEAMS)
	err := collection.Find(bson.M{"id": id}).One(&updatedTeam)
	if err != nil {
		fmt.Println(err)
	}
	riders := updatedTeam.Riders
	index := FindCyclist(riders, riderId)
	fmt.Println(index)
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
	} else {
		return nil
	}
}