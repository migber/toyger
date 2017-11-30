package events

import (
	"fmt"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
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

func CreateStage(eventId string, s Stage, session *mgo.Session) Stage {
	
	var stage Stage
	defer session.Close() 

	existsStage := GetStageInside(eventId, s.ID, session)
	fmt.Println(existsStage)
	if existsStage != nil {
		stage.ID = s.ID
		stage.Event = s.Event
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
		
		collection := session.DB(DATABASE).C(STAGES)
		if err := collection.Insert(stage); err != nil {
			panic(err)
		}
	} else {
		fmt.Println("User with the same UCIID exists")
	}
	return stage
}

func GetStageInside(eventId string, id int, session *mgo.Session) error {
	
	var s Stage
	collection := session.DB(DATABASE).C(STAGES)
	err := collection.Find(bson.M{"id": id, "event": eventId}).One(&s)
	if err != nil {
		fmt.Println(err)
	}

	return err
}

func GetStageList(eventId string, session *mgo.Session) Stages {
	
	stages := Stages{}
	defer session.Close()

	c := session.DB(DATABASE).C(STAGES)
	err := c.Find(nil).All(&stages)
	if err != nil {
		panic(err)
	}
	return stages
}

func GetStage(eventId string, id int, session *mgo.Session) Stage {

	var s Stage
	defer session.Close()
	collection := session.DB(DATABASE).C(STAGES)
	err := collection.Find(bson.M{"id": id, "event": eventId}).One(&s)
	if err != nil {
		fmt.Println(err)
	}

	return s
}

func UpdateStage(eventId string, id int, s Stage, session *mgo.Session) Stage {
	
	var updateStage Stage

	defer session.Close()
	
	collection := session.DB(DATABASE).C(STAGES)

	err := collection.Find(bson.M{"id": id, "event": eventId}).One(&updateStage)
	if err != nil {
		panic(err)
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

	if err := collection.Update(bson.M{"id": id}, updateStage); err != nil {
		panic(err)
	}
	
	return updateStage
}

func DeleteStage(eventId string, id int, session *mgo.Session) error {
	
	defer session.Close()
	
	collection := session.DB(DATABASE).C(STAGES)
	err := collection.Remove(bson.M{"id": id, "event": eventId})
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Object was deleted")

	return err
}