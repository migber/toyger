package events
import (
	model "toyger/models"
	"fmt"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"errors"
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

func CreateRaceCommissaire(eventID string, c RaceCommissaire, session *mgo.Session,
						 dbName string, tableName string) (RaceCommissaire, error) {
	var raceCom RaceCommissaire
	defer session.Close() 
	var err error
	existingCommi := GetRaceCommissaireInside(eventID, c.Commissaire.UCIID, session, dbName, tableName)
	if existingCommi != nil {
		raceCom.Commissaire = c.Commissaire
		raceCom.Event = eventID
		raceCom.Position = c.Position
		collection := session.DB(dbName).C(tableName)
		if err2 := collection.Insert(raceCom); err2 != nil {
			fmt.Println(err2)
			err = err2
		}
	} else {
		fmt.Println("User with the same UCIID exists")
		err = errors.New("User with the same UCIID exists")
	}
	return raceCom, err
}

func GetRaceCommissaireInside(eventID string, id string, session *mgo.Session,
							 dbName string, tableName string) error {

	var c RaceCommissaire
	collection := session.DB(dbName).C(tableName)
	err := collection.Find(bson.M{"commissaire.uciid": id, "event": eventID}).One(&c)
	// if err != nil {
	// 	fmt.Println(err)
	// }

	return err
}

func GetRaceCommissairesList(eventID string, session *mgo.Session,
							dbName string, tableName string) (RaceCommissaires, error) {
	
	raceComms := RaceCommissaires{}
	defer session.Close()

	c := session.DB(dbName).C(tableName)
	err := c.Find(bson.M{"event": eventID}).All(&raceComms)
	if err != nil {
		fmt.Println(err)
	}
	return raceComms, err
}

func GetRaceCommissaire(eventID string, id string, session *mgo.Session,
	dbName string, tableName string) (RaceCommissaire, error) {

	var c RaceCommissaire
	defer session.Close()
	collection := session.DB(dbName).C(tableName)
	err := collection.Find(bson.M{"commissaire.uciid": id, "event": eventID}).One(&c)
	if err != nil {
		fmt.Println(err)
	}
	return c, err
}

func UpdateRaceCommissaire(eventID string, id string, raceCom RaceCommissaire, session *mgo.Session,
	dbName string, tableName string) (RaceCommissaire, error) {
	
	var updateRaceComm RaceCommissaire

	defer session.Close()
	
	collection := session.DB(dbName).C(tableName)

	err := collection.Find(bson.M{"commissaire.uciid": id, "event": eventID}).One(&updateRaceComm)
	if err != nil {
		fmt.Println(err)
	}
	updateRaceComm.Commissaire = raceCom.Commissaire
	updateRaceComm.Position = raceCom.Position

	if err2 := collection.Update(bson.M{"commissaire.uciid": id}, updateRaceComm); err2 != nil {
		fmt.Println(err2)
		err = err2
	}
	
	return updateRaceComm, err
}

func DeleteRaceCommissaire(eventID string, id string, session *mgo.Session,
	dbName string, tableName string) error {
	
	defer session.Close()
	
	collection := session.DB(dbName).C(tableName)
	err := collection.Remove(bson.M{"commissaire.uciid": id, "event": eventID})
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Object was deleted")	
	}

	return err
}