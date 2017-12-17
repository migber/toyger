package events

import (
	"fmt"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"errors"
)

type Stage struct {
	ID				int		    `json:"id"`
	Event			string		`json:"event"`
	Name 			string		`json:"name	"`
	Km				int			`json:"km"`
	Sprints			[]Sprint	`json:"sprints"`
	Abandoned		int			`json:"abandoned"`
	Disqualified	int			`json:"disqualfied"`
	Starters		int			`json:"starters"`
	Remaining		int			`json:"remaining"`
	AfterTimeLimit	int			`json:"after_time_limit"`
	TimeLimit 		int			`json:"time_limit"`
	NotStarted		int			`json:"not_started"`
}

type Stages []Stage
var STAGES = "stages"

func CreateStage(eventID string, s Stage, session *mgo.Session,
	dbName string, tableName string) (Stage, error) {
	
	var stage Stage
	defer session.Close() 
	var err error
	existsStage := GetStageInside(eventID, s.ID, session, dbName, tableName)
	if existsStage != nil {
		stage.ID = s.ID
		stage.Event = eventID
		stage.Name = s.Name
		stage.Km = s.Km
		stage.Sprints = s.Sprints
		stage.Abandoned = s.Abandoned
		stage.Disqualified = s.Disqualified
		stage.Starters = s.Starters
		stage.Remaining = s.Remaining
		stage.TimeLimit = s.TimeLimit
		stage.AfterTimeLimit = s.AfterTimeLimit
		stage.NotStarted = s.NotStarted	
		
		collection := session.DB(dbName).C(tableName)
		if err := collection.Insert(stage); err != nil {
			fmt.Println(err)
			return stage, err
		}
	} else {
		fmt.Println("User with the same UCIID exists")
		err = errors.New("User with the same UCIID exists")
	}
	return stage, err
}

func GetStageInside(eventID string, id int, session *mgo.Session,
	dbName string, tableName string) error {
	
	var s Stage
	collection := session.DB(dbName).C(tableName)
	err := collection.Find(bson.M{"id": id, "event": eventID}).One(&s)
	if err != nil {
		fmt.Println(err)
	}

	return err
}

func GetStageList(eventID string, session *mgo.Session,
	dbName string, tableName string) (Stages, error) {
	
	stages := Stages{}
	defer session.Close()

	c := session.DB(dbName).C(tableName)
	err := c.Find(bson.M{"event": eventID}).All(&stages)
	if err != nil {
		fmt.Println(err)
	}
	return stages, err
}

func GetStage(eventID string, id int, session *mgo.Session,
	dbName string, tableName string) (Stage, error) {

	var s Stage
	defer session.Close()
	collection := session.DB(dbName).C(tableName)
	err := collection.Find(bson.M{"id": id, "event": eventID}).One(&s)
	if err != nil {
		fmt.Println(err)
	}

	return s, err
}

func UpdateStage(eventID string, id int, s Stage, session *mgo.Session,
	dbName string, tableName string) (Stage, error) {
	
	var updateStage Stage

	defer session.Close()
	
	collection := session.DB(dbName).C(tableName)

	err := collection.Find(bson.M{"id": id, "event": eventID}).One(&updateStage)
	if err != nil {
		fmt.Println(err)
		return updateStage, err
	}
	updateStage.Event = s.Event
	updateStage.Name = s.Name
	updateStage.Km = s.Km
	updateStage.Sprints = s.Sprints
	updateStage.Abandoned = s.Abandoned
	updateStage.Disqualified = s.Disqualified
	updateStage.Starters = s.Starters
	updateStage.Remaining = s.Remaining
	updateStage.TimeLimit = s.TimeLimit
	updateStage.AfterTimeLimit = s.AfterTimeLimit
	updateStage.NotStarted = s.NotStarted	

	if err2 := collection.Update(bson.M{"id": id}, updateStage); err2 != nil {
		fmt.Println(err2)
		err = err2
	}
	
	return updateStage, err
}

func DeleteStage(eventID string, id int, session *mgo.Session,
	dbName string, tableName string) error {
	
	defer session.Close()
	
	collection := session.DB(dbName).C(tableName)
	err := collection.Remove(bson.M{"id": id, "event": eventID})
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Object was deleted")

	return err
}

func AddSprints(eventID string, stageID int, s Sprint, session *mgo.Session,
	dbName string, tableName string) error {
	
	var updateStages Stage
	defer session.Close()
	collection := session.DB(dbName).C(tableName)
	err := collection.Find(bson.M{"id": stageID}).One(&updateStages)
	if err != nil {
		return err
	}
	sprints := updateStages.Sprints
	newSprints := append(sprints, s)
	updateStages.Sprints = newSprints
	if err := collection.Update(bson.M{"id": stageID}, updateStages); err != nil {
		return err
	}
	return nil
}

func DeleteStageSprint(id string, stageID int, sprintID string, session *mgo.Session,
	dbName string, tableName string) error{
	
	var updateStage Stage
	var emptySprints []Sprint
	defer session.Close()
	
	collection := session.DB(dbName).C(tableName)
	err := collection.Find(bson.M{"id": id}).One(&updateStage)
	if err != nil {
		fmt.Println(err)
	}
	sprints := updateStage.Sprints
	index := FindStageSprint(sprints, sprintID)
	if (len(sprints) == 1 && index == 0){
		updateStage.Sprints = emptySprints
	} else {
		newSprint := append(sprints[:index], sprints[index+1:]...)
		updateStage.Sprints = newSprint
	}
	newS := append(sprints[:index], sprints[index+1:]...)
	updateStage.Sprints = newS
	if err := collection.Update(bson.M{"id": id}, updateStage); err != nil {
		fmt.Println(err)
		return err
	} else {
		return nil
	}
}

func FindStageSprint(sprints []Sprint, id string) int {
	
	var index int
	for ind, sprint := range sprints {
		if sprint.Id == id {
			index = ind
		}
	}
	return index
}

func UpdateStageSprint(stageID int, sprintID string, sprintNew Sprint, session *mgo.Session,
	dbName string, tableName string) error{
	
	var updateStage Stage
	defer session.Close()
	
	collection := session.DB(dbName).C(tableName)
	err := collection.Find(bson.M{"id": stageID}).One(&updateStage)
	if err != nil {
		fmt.Println(err)
	}
	sprints := updateStage.Sprints
	index := FindStageSprint(sprints, sprintID)
	updateStage.Sprints[index] = sprintNew
	if err := collection.Update(bson.M{"id": stageID}, updateStage); err != nil {
		fmt.Println(err)
		return err
	} else {
		return nil
	}
}
	