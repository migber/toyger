package events

import (
	"fmt"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"github.com/satori/go.uuid"
	"errors"
)

type Sprint struct {
	Id 			string		`json:"id"`
	Name		string		`json:"name"`
	Event		string		`json:"event"`
	Stage		int			`json:"stage"`
	Category 	string		`json:"category"`
	Winner		int			`json:"winner"`
	Second		int			`json:"second"`
	Third		int			`json:"third"`
	Bonuses 	[]int		`json:"bonuses"`
}

type Sprints []Sprint
var SPRINTS = "sprints"

func CreateSprint(eventID string, stageID int, s Sprint, session *mgo.Session,
	dbName string, tableName string) (Sprint, error) {
	
	var sprint Sprint
	defer session.Close() 
	var err error
	existingSprint := GetSprintInside(eventID, stageID, s.Id, 
									  session, dbName, tableName)
	if existingSprint != nil {
		sprint.Id = uuid.NewV4().String()
		sprint.Name = s.Name
		sprint.Event = eventID
		sprint.Stage = stageID
		sprint.Category = s.Category
		sprint.Winner = s.Winner
		sprint.Second = s.Second
		sprint.Third = s.Third
		sprint.Bonuses = s.Bonuses

		collection := session.DB(dbName).C(tableName)
		if err := collection.Insert(sprint); err != nil {
			fmt.Println(err)
			return sprint, err
		}
	} else {
		fmt.Println("User with the same UCIID exists")
		err = errors.New("User with the same UCIID exists")
	}
	return sprint, err
}

func GetSprintInside(eventID string, stageID int, id string, session *mgo.Session,
	dbName string, tableName string) error {
	
	var s Sprint
	collection := session.DB(dbName).C(tableName)
	err := collection.Find(bson.M{"id": id, "event": eventID, "stage": stageID}).One(&s)
	if err != nil {
		fmt.Println(err)
	}

	return err
}

func GetSprintList(eventID string, stageID int, session *mgo.Session,
	dbName string, tableName string) (Sprints, error) {
	
	sprints := Sprints{}
	defer session.Close()

	c := session.DB(dbName).C(tableName)
	err := c.Find(bson.M{"event": eventID, "stage": stageID}).All(&sprints)
	if err != nil {
		fmt.Println(err)
	}
	return sprints, err
}

func GetSprint(eventID string, stageID int, id string, session *mgo.Session,
	dbName string, tableName string) (Sprint, error) {
	
	var s Sprint
	defer session.Close()
	collection := session.DB(dbName).C(tableName)
	err := collection.Find(bson.M{"id": id, "event": eventID, "stage": stageID}).One(&s)
	if err != nil {
		fmt.Println(err)
	}

	return s, err
}

func UpdateSprint(eventID string, stageID int, id string, s Sprint, session *mgo.Session,
	dbName string, tableName string) (Sprint, error) {
	
	var updateSprint Sprint

	defer session.Close()
	
	collection := session.DB(dbName).C(tableName)
	err := collection.Find(bson.M{"id": id, "event": eventID, "stage": stageID}).One(&updateSprint)
	if err != nil {
		fmt.Println(err)
		return updateSprint, err
	}
	updateSprint.Name = s.Name
	updateSprint.Category = s.Category
	updateSprint.Winner = s.Winner
	updateSprint.Second = s.Second
	updateSprint.Third = s.Third
	updateSprint.Bonuses = s.Bonuses

	if err2 := collection.Update(bson.M{"id": id}, updateSprint); err2 != nil {
		fmt.Println(err2)
		err = err2
	}
	
	return updateSprint, err
}

func DeleteSprint(eventID string, stageID int, id string, session *mgo.Session,
	dbName string, tableName string) error {
	
	defer session.Close()
	
	collection := session.DB(dbName).C(tableName)
	err := collection.Remove(bson.M{"id": id, "event": eventID, "stage": stageID})
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Object was deleted")

	return err
}
