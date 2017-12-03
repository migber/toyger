package models

import (
	"fmt"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type Cyclist struct{
	UCIID		string		`json:"uci_id"`
	Name		string		`json:"name"`
	Surname		string		`json:"surname"`
	Team 		string		`json:"team"`
	Coaches 	[]string	`json:"coaches"`
	Birthdate	string		`json:"birth_date"`
	Gender		string		`json:"gender"`
	UCICategory	string		`json:"uci_category"`
	Nationality string		`json:"nationality"`
}

type Cyclists []Cyclist	

var DATABASE = "toyger"
var CYCLISTS = "cyclists"

func CreateCyclist(teamId string, c Cyclist, session *mgo.Session) Cyclist {
	
	var cyclist Cyclist
	defer session.Close() 

	existsCyclist := GetCyclistInside(teamId, c.UCIID, session)
	fmt.Println(existsCyclist)
	if existsCyclist != nil {
		cyclist.UCIID = c.UCIID
		cyclist.Name = c.Name
		cyclist.Surname = c.Surname
		cyclist.Team = teamId
		cyclist.Coaches = c.Coaches
		cyclist.Birthdate = c.Birthdate
		cyclist.Gender = c.Gender
		cyclist.UCICategory = c.UCICategory
		cyclist.Nationality = c.Nationality
		
		collection := session.DB(DATABASE).C(CYCLISTS)
		if err := collection.Insert(cyclist); err != nil {
			panic(err)
		}
	} else {
		fmt.Println("User with the same UCIID exists")
	}
	return cyclist
}

func GetCyclistsList(teamId string, session *mgo.Session) Cyclists {
	
		cyclists := Cyclists{}
		defer session.Close()

		c := session.DB(DATABASE).C(CYCLISTS)
		err := c.Find(bson.M{"team": teamId}).All(&cyclists)
		if err != nil {
			panic(err)
		}
		return cyclists
}

func GetCyclist(teamId string, id string, session *mgo.Session) Cyclist {
	
		var c Cyclist
		defer session.Close()
		collection := session.DB(DATABASE).C(CYCLISTS)
		err := collection.Find(bson.M{"uciid": id, "team": teamId}).One(&c)
		if err != nil {
			fmt.Println(err)
		}
	
		return c
}

func GetCyclistInside(teamId string, id string, session *mgo.Session) error {
	
		var c Cyclist
		collection := session.DB(DATABASE).C(CYCLISTS)
		err := collection.Find(bson.M{"uciid": id, "team": teamId}).One(&c)
		if err != nil {
			fmt.Println(err)
		}
	
		return err
}

func UpdateCyclist(teamId string, uid string, c Cyclist, session *mgo.Session) Cyclist {
	
		var updateC Cyclist
		defer session.Close()
		
		collection := session.DB(DATABASE).C(CYCLISTS)
	
		err := collection.Find(bson.M{"uciid": uid, "team": teamId}).One(&updateC)
		if err != nil {
			panic(err)
		}
	
		updateC.UCIID = c.UCIID
		updateC.Name = c.Name
		updateC.Surname = c.Surname
		updateC.Team = c.Team
		updateC.Coaches = c.Coaches
		updateC.Birthdate = c.Birthdate
		updateC.Gender = c.Gender
		updateC.UCICategory = c.UCICategory
		updateC.Nationality = c.Nationality
	
		if err := collection.Update(bson.M{"uciid": uid}, updateC); err != nil {
			panic(err)
		}
		
		return updateC
}


func DeleteCyclist(teamId string, uid string, session *mgo.Session) error {

	defer session.Close()
	
	collection := session.DB(DATABASE).C(CYCLISTS)
	err := collection.Remove(bson.M{"uciid": uid, "team": teamId})
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Object was deleted")

	return err
}
