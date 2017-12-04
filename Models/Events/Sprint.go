package events

import (
	"fmt"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"github.com/satori/go.uuid"
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

func CreateSprint(eventId string, stageId int, s Sprint, session *mgo.Session,
	dbName string, tableName string) Sprint {
	
	var sprint Sprint
	defer session.Close() 

	existingSprint := GetSprintInside(eventId, stageId, s.Id, 
									  session, dbName, tableName)
	if existingSprint != nil {
		sprint.Id = uuid.NewV4().String()
		sprint.Name = s.Name
		sprint.Event = eventId
		sprint.Stage = stageId
		sprint.Category = s.Category
		sprint.Winner = s.Winner
		sprint.Second = s.Second
		sprint.Third = s.Third
		sprint.Bonuses = s.Bonuses

		collection := session.DB(dbName).C(tableName)
		if err := collection.Insert(sprint); err != nil {
			panic(err)
		}
	} else {
		fmt.Println("User with the same UCIID exists")
	}
	return sprint
}

func GetSprintInside(eventId string, stageId int, id string, session *mgo.Session,
	dbName string, tableName string) error {
	
	var s Sprint
	collection := session.DB(dbName).C(tableName)
	err := collection.Find(bson.M{"id": id, "event": eventId, "stage": stageId}).One(&s)
	if err != nil {
		fmt.Println(err)
	}

	return err
}

func GetSprintList(eventId string, stageId int, session *mgo.Session,
	dbName string, tableName string) Sprints {
	
	sprints := Sprints{}
	defer session.Close()

	c := session.DB(dbName).C(tableName)
	err := c.Find(bson.M{"event": eventId, "stage": stageId}).All(&sprints)
	if err != nil {
		panic(err)
	}
	return sprints
}

func GetSprint(eventId string, stageId int, id string, session *mgo.Session,
	dbName string, tableName string) Sprint {
	
	var s Sprint
	defer session.Close()
	collection := session.DB(dbName).C(tableName)
	err := collection.Find(bson.M{"id": id, "event": eventId, "stage": stageId}).One(&s)
	if err != nil {
		fmt.Println(err)
	}

	return s
}

func UpdateSprint(eventId string, stageId int, id string, s Sprint, session *mgo.Session,
	dbName string, tableName string) Sprint {
	
	var updateSprint Sprint

	defer session.Close()
	
	collection := session.DB(dbName).C(tableName)
	err := collection.Find(bson.M{"id": id, "event": eventId, "stage": stageId}).One(&updateSprint)
	if err != nil {
		panic(err)
	}
	updateSprint.Name = s.Name
	updateSprint.Category = s.Category
	updateSprint.Winner = s.Winner
	updateSprint.Second = s.Second
	updateSprint.Third = s.Third
	updateSprint.Bonuses = s.Bonuses

	if err := collection.Update(bson.M{"id": id}, updateSprint); err != nil {
		panic(err)
	}
	
	return updateSprint
}

func DeleteSprint(eventId string, stageId int, id string, session *mgo.Session,
	dbName string, tableName string) error {
	
	defer session.Close()
	
	collection := session.DB(dbName).C(tableName)
	err := collection.Remove(bson.M{"id": id, "event": eventId, "stage": stageId})
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Object was deleted")

	return err
}
