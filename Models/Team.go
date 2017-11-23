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

var team Team
var teams Teams
var DATABASE = "toyger"
var NAME = "teams"

func CreateTeam(t Team, session *mgo.Session) Team {

	defer session.Close() 

	t.Id = uuid.NewV4().String()
	t.Manager.Id = uuid.NewV4().String()
	team.Id = t.Id
	team.Name = t.Name
	team.Manager = t.Manager
	team.Riders = t.Riders
	
	collection := session.DB(DATABASE).C(NAME)
	if err := collection.Insert(team); err != nil {
		panic(err)
	}
	//teams = append(teams, team)
	return t
}

func GetTeamsList(session *mgo.Session) Teams {

	teams := Teams{}
	defer session.Close()
	c := session.DB(DATABASE).C(NAME)
	err := c.Find(nil).All(&teams)
	if err != nil {
		panic(err)
	}
	return teams
}

func GetTeam(id string, session *mgo.Session) Team {

	var t Team
	defer session.Close()
	collection := session.DB(DATABASE).C(NAME)
	err := collection.Find(bson.M{"id": id}).One(&t)
	if err != nil {
		panic(err)
	}

	return t
}

func UpdateTeam(uid string, t Team, session *mgo.Session) Team {

	// updateTeam, ind := FindTeam(uid)

	// updateTeam.Name = t.Name
	// updateTeam.Manager.Name = t.Manager.Name
	// updateTeam.Manager.Surname = t.Manager.Surname
	// updateTeam.Manager.Birthdate = t.Manager.Birthdate
	// updateTeam.Manager.Phone = t.Manager.Phone
	// updateTeam.Manager.Nationality = t.Manager.Nationality
	defer session.Close()
	
	collection := session.DB(DATABASE).C(NAME)
	var updateT Team
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
	//teams[ind] = updateTeam
	
	return updateT
}

func DeleteTeam(uid string, session *mgo.Session) error {

	// delete := -1
	//u2, err :=  uuid.FromString(uid)
	// if err != nil{
	// 	fmt.Errorf("Error occured while parsing uuid %v", err)	
	// // }

	// for ind, team := range teams {
	// 	if team.Id == uid {
	// 		delete = ind
	// 	}
	// }

	// if delete != -1{
	// 	teams = append(teams[:delete], teams[delete + 1:]...)
	// 	return nil
	// }
	collection := session.DB(DATABASE).C(NAME)
	err := collection.Remove(bson.M{"id": uid})
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Object was deleted")

	return err
}

func FindTeam(id string) (Team, int) {

	var t Team
	var index int
	//u2, err :=  uuid.FromString(id)
	// if err != nil{
	// 	fmt.Errorf("Error occured while parsing uuid %v", err)	
	// }

	for ind, team := range teams {
		if team.Id == id {
			t = team
			index = ind
		}
	}
	return t, index
}