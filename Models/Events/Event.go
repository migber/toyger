package events

import (
	"fmt"
	"github.com/satori/go.uuid"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"	
)

type Event struct {	
	ID 				string				`json:"id"`
	Name 			string				`json:"name"`
	NoParticipants	int					`json:"no_participants"`
	NoStages 		int					`json:"no_stages"`
	Date 			string			    `json:"date"`
	Location 		string				`json:"location"`
	NoCommissaires 	int 				`json:"no_commissaires"`
	TotalKm 		int					`json:"total_km"`
	Stages			[]int		        `json:"stages"`
	Participants	[]int		    `json:"participants"`
	Commissaires	[]string	        `json:"commissaires"`
}

type Events []Event 

var DATABASE = "toyger"
var EVENTS = "events"

func CreateEvent(e Event, session *mgo.Session, dbName string, dbTable string) Event {

	var event Event
	defer session.Close() 

	event.ID = uuid.NewV4().String()
	event.Name = e.Name
	event.NoParticipants = e.NoParticipants
	event.NoStages = e.NoStages
	event.Date = e.Date
	event.Location = e.Location
	event.NoCommissaires = e.NoCommissaires
	event.TotalKm = e.TotalKm
	event.Stages = e.Stages
	event.Participants = e.Participants
	event.Commissaires = e.Commissaires
	
	
	collection := session.DB(dbName).C(dbTable)
	if err := collection.Insert(event); err != nil {
		fmt.Println(err)
	}
	return event
}

func GetEventsList(session *mgo.Session, dbName string, tableName string) Events {
	
	events := Events{}
	defer session.Close()
	c := session.DB(dbName).C(tableName)
	err := c.Find(nil).All(&events)
	if err != nil {
		fmt.Println(err)
	}
	return events
}

func GetEvent(id string, session *mgo.Session, dbName string, tableName string) Event {
	
	var e Event
	defer session.Close()
	collection := session.DB(dbName).C(tableName)
	err := collection.Find(bson.M{"id": id}).One(&e)
	if err != nil {
		fmt.Println(err)
	}
	return e
}

func UpdateEvent(id string, e Event, session *mgo.Session, dbName string, tableName string) Event {
	
	var updateEvent Event

	defer session.Close()
	
	collection := session.DB(dbName).C(tableName)

	err := collection.Find(bson.M{"id": id}).One(&updateEvent)
	if err != nil {
		fmt.Println(err)
	}

	updateEvent.Name = e.Name
	updateEvent.NoParticipants = e.NoParticipants
	updateEvent.NoStages = e.NoStages
	updateEvent.Date = e.Date
	updateEvent.Location = e.Location
	updateEvent.NoCommissaires = e.NoCommissaires
	updateEvent.TotalKm = e.TotalKm
	updateEvent.Stages = e.Stages
	updateEvent.Participants = e.Participants
	updateEvent.Commissaires = e.Commissaires

	if err := collection.Update(bson.M{"id": id}, updateEvent); err != nil {
		fmt.Println(err)
	}
	
	return updateEvent
}

func DeleteEvent(id string, session *mgo.Session, dbName string, tableName string) error {
	
	defer session.Close()
	collection := session.DB(dbName).C(tableName)
	err := collection.Remove(bson.M{"id": id})
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Object was deleted")

	return err
}

func FindEventsParticipant(riders []int, id int) int {
	
	index:= -1
	for ind, cyclist := range riders {
		if cyclist == id {
			index = ind
		}
	}
	return index
}

func FindEventStage(stages []int, id int) int {
	
	index:= -1
	for ind, stage := range stages {
		if stage == id {
			index = ind
		}
	}
	return index
}

func FindEventCommissaires(commissaires []string, id string) int {
	
	index := -1
	for ind, com := range commissaires {
		if com == id {
			index = ind
		}
	}
	return index
}

func InsertEventParticipants(id string, riderId int, session *mgo.Session,
							 dbName string, tableName string) error {
	
	var updatedEvent Event
	defer session.Close()
	
	collection := session.DB(dbName).C(tableName)
	err := collection.Find(bson.M{"id": id}).One(&updatedEvent)
	if err != nil {
		return err
	}
	riders := updatedEvent.Participants
	newRiders := append(riders, riderId)
	updatedEvent.Participants = newRiders
	if err := collection.Update(bson.M{"id": id}, updatedEvent); err != nil {
		return err
	}
	return nil
}

func InsertEventCommissaire(id string, comId string, session *mgo.Session,
							dbName string, tableName string) error {
	
	var updatedEvent Event
	defer session.Close()
	
	collection := session.DB(dbName).C(tableName)
	err := collection.Find(bson.M{"id": id}).One(&updatedEvent)
	if err != nil {
		return err
	}
	commissaires := updatedEvent.Commissaires
	newComm := append(commissaires, comId)
	updatedEvent.Commissaires = newComm
	if err := collection.Update(bson.M{"id": id}, updatedEvent); err != nil {
		return err
	}
	return nil
}

func InsertEventStages(id string, stageId int, session *mgo.Session,
						dbName string, tableName string) error {
	
	var updatedEvent Event
	defer session.Close()
	
	collection := session.DB(dbName).C(tableName)
	err := collection.Find(bson.M{"id": id}).One(&updatedEvent)
	if err != nil {
		return err
	}
	stages := updatedEvent.Stages
	newStages := append(stages, stageId)
	updatedEvent.Stages = newStages
	if err := collection.Update(bson.M{"id": id}, updatedEvent); err != nil {
		return err
	}
	return nil
}

func DeleteEventParticipant(id string, riderId int, session *mgo.Session,
	dbName string, tableName string) error{
	
	var updateEvent Event
	var emptyRiders []int
	defer session.Close()
	
	collection := session.DB(dbName).C(tableName)
	err := collection.Find(bson.M{"id": id}).One(&updateEvent)
	if err != nil {
		fmt.Println(err)
	}
	riders := updateEvent.Participants
	index := FindEventsParticipant(riders, riderId)
	if (len(riders) == 1 && index == 0){
		updateEvent.Participants = emptyRiders
	} else {
		newRiders := append(riders[:index], riders[index+1:]...)
		updateEvent.Participants = newRiders
	}
	newRiders := append(riders[:index], riders[index+1:]...)
	updateEvent.Participants = newRiders
	if err := collection.Update(bson.M{"id": id}, updateEvent); err != nil {
		fmt.Println(err)
		return err
	} else {
		return nil
	}
}

func DeleteEventCommissaire(id string, comId string, session *mgo.Session,
		dbName string, tableName string) error{
	
	var updateEvent Event
	var emptyCom []string
	defer session.Close()
	
	collection := session.DB(dbName).C(tableName)
	err := collection.Find(bson.M{"id": id}).One(&updateEvent)
	if err != nil {
		fmt.Println(err)
	}
	commissaires := updateEvent.Commissaires
	index := FindEventCommissaires(commissaires, comId)
	if (len(commissaires) == 1 && index == 0){
		updateEvent.Commissaires = emptyCom
	} else {
		newComm := append(commissaires[:index], commissaires[index+1:]...)
		updateEvent.Commissaires = newComm
	}
	newCommi := append(commissaires[:index], commissaires[index+1:]...)
	updateEvent.Commissaires = newCommi
	if err := collection.Update(bson.M{"id": id}, updateEvent); err != nil {
		fmt.Println(err)
		return err
	} else {
		return nil
	}
}

func DeleteEventStage(id string, stageId int, session *mgo.Session,
						dbName string, tableName string) error{
	
	var updateEvent Event
	var emptyStage []int
	defer session.Close()
	
	collection := session.DB(dbName).C(tableName)
	err := collection.Find(bson.M{"id": id}).One(&updateEvent)
	if err != nil {
		fmt.Println(err)
	}
	stages := updateEvent.Stages
	index := FindEventStage(stages, stageId)
	if (len(stages) == 1 && index == 0){
		updateEvent.Stages = emptyStage
	} else {
		newStage := append(stages[:index], stages[index+1:]...)
		updateEvent.Stages = newStage
	}
	newS := append(stages[:index], stages[index+1:]...)
	updateEvent.Stages = newS
	if err := collection.Update(bson.M{"id": id}, updateEvent); err != nil {
		fmt.Println(err)
		return err
	} else {
		return nil
	}
}