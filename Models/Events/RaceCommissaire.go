package events
import (
	model "toyger/models"
	"fmt"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)
type Position struct{
	Stage	string	`json:"stage"`
	Name 	string	`json:"name"`
}
type RaceCommissaire struct{
	Commissaire 	model.Commissaire	`json:"commissaire"`
	Event			string 				`json:"event"`
	Position		[]Position          `json:"position"`
}

type RaceCommissaires []RaceCommissaire	
var RACECOMMISSAIRES = "racecommissaires"

func CreateRaceCommissaire(eventId string, c RaceCommissaire, session *mgo.Session,
						 dbName string, tableName string) RaceCommissaire {
	var raceCom RaceCommissaire
	defer session.Close() 

	existingCommi := GetRaceCommissaireInside(eventId, c.Commissaire.UCIID, session, dbName, tableName)
	if existingCommi != nil {
		raceCom.Commissaire = c.Commissaire
		raceCom.Event = eventId
		raceCom.Position = c.Position
		collection := session.DB(dbName).C(tableName)
		if err := collection.Insert(raceCom); err != nil {
			panic(err)
		}
	} else {
		fmt.Println("User with the same UCIID exists")
	}
	return raceCom
}

func GetRaceCommissaireInside(eventId string, id string, session *mgo.Session,
							 dbName string, tableName string) error {

	var c RaceCommissaire
	collection := session.DB(dbName).C(tableName)
	err := collection.Find(bson.M{"commissaire.uciid": id, "event": eventId}).One(&c)
	// if err != nil {
	// 	fmt.Println(err)
	// }

	return err
}

func GetRaceCommissairesList(eventId string, session *mgo.Session,
							dbName string, tableName string) RaceCommissaires {
	
	raceComms := RaceCommissaires{}
	defer session.Close()

	c := session.DB(dbName).C(tableName)
	err := c.Find(bson.M{"event": eventId}).All(&raceComms)
	if err != nil {
		panic(err)
	}
	return raceComms
}

func GetRaceCommissaire(eventId string, id string, session *mgo.Session,
	dbName string, tableName string) RaceCommissaire {

	var c RaceCommissaire
	defer session.Close()
	collection := session.DB(dbName).C(tableName)
	err := collection.Find(bson.M{"commissaire.uciid": id, "event": eventId}).One(&c)
	if err != nil {
		fmt.Println(err)
	}
	return c
}

func UpdateRaceCommissaire(eventId string, id string, raceCom RaceCommissaire, session *mgo.Session,
	dbName string, tableName string) RaceCommissaire {
	
	var updateRaceComm RaceCommissaire

	defer session.Close()
	
	collection := session.DB(dbName).C(tableName)

	err := collection.Find(bson.M{"commissaire.uciid": id, "event": eventId}).One(&updateRaceComm)
	if err != nil {
		panic(err)
	}
	updateRaceComm.Commissaire = raceCom.Commissaire
	updateRaceComm.Position = raceCom.Position

	if err := collection.Update(bson.M{"commissaire.uciid": id}, updateRaceComm); err != nil {
		panic(err)
	}
	
	return updateRaceComm
}

func DeleteRaceCommissaire(eventId string, id string, session *mgo.Session,
	dbName string, tableName string) error {
	
	defer session.Close()
	
	collection := session.DB(dbName).C(tableName)
	err := collection.Remove(bson.M{"commissaire.uciid": id, "event": eventId})
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Object was deleted")	
	}

	return err
}