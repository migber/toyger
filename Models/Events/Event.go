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
	Stages			[]string		    `json:"stages"`
	Participants	[]string		    `json:"participants"`
	Commissaires	[]string	        `json:"commissaires"`
}

type Events []Event 

var DATABASE = "toyger"
var EVENTS = "events"

func CreateEvent(e Event, session *mgo.Session) Event {

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
	
	
	collection := session.DB(DATABASE).C(EVENTS)
	if err := collection.Insert(event); err != nil {
		panic(err)
	}
	return event
}

func GetEventsList(session *mgo.Session) Events {
	
	events := Events{}
	defer session.Close()
	c := session.DB(DATABASE).C(EVENTS)
	err := c.Find(nil).All(&events)
	if err != nil {
		panic(err)
	}
	return events
}

func GetEvent(id string, session *mgo.Session) Event {
	
	var e Event
	defer session.Close()
	collection := session.DB(DATABASE).C(EVENTS)
	err := collection.Find(bson.M{"id": id}).One(&e)
	if err != nil {
		panic(err)
	}
	fmt.Println(e)

	return e
}

func UpdateEvent(id string, e Event, session *mgo.Session) Event {
	
	var updateEvent Event

	defer session.Close()
	
	collection := session.DB(DATABASE).C(EVENTS)

	err := collection.Find(bson.M{"id": id}).One(&updateEvent)
	if err != nil {
		panic(err)
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
		panic(err)
	}
	
	return updateEvent
}

func DeleteEvent(id string, session *mgo.Session) error {
	
	defer session.Close()
	
	collection := session.DB(DATABASE).C(EVENTS)
	err := collection.Remove(bson.M{"id": id})
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Object was deleted")

	return err
}

func FindEventsParticipant(riders []string, id string) int {
	
	var index int
	for ind, cyclist := range riders {
		if cyclist == id {
			index = ind
		}
	}
	return index
}

func FindEventStage(stages []string, id string) int {
	
	var index int
	for ind, stage := range stages {
		if stage == id {
			index = ind
		}
	}
	return index
}

func FindEventCommissaires(commissaires []string, id string) int {
	
	var index int
	for ind, com := range commissaires {
		if com == id {
			index = ind
		}
	}
	return index
}

func InsertEventParticipants(id string, riderId string, session *mgo.Session) error {
	
	var updatedEvent Event
	defer session.Close()
	
	collection := session.DB(DATABASE).C(EVENTS)
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

func InsertEventCommissaire(id string, comId string, session *mgo.Session) error {
	
	var updatedEvent Event
	defer session.Close()
	
	collection := session.DB(DATABASE).C(EVENTS)
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

func InsertEventStages(id string, stageId string, session *mgo.Session) error {
	
	var updatedEvent Event
	defer session.Close()
	
	collection := session.DB(DATABASE).C(EVENTS)
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

func DeleteEventParticipant(id string, riderId string, session *mgo.Session) error{
	
	var updateEvent Event
	var emptyRiders []string
	defer session.Close()
	
	collection := session.DB(DATABASE).C(EVENTS)
	err := collection.Find(bson.M{"id": id}).One(&updateEvent)
	if err != nil {
		fmt.Println(err)
	}
	riders := updateEvent.Participants
	index := FindEventsParticipant(riders, riderId)
	fmt.Println(index)
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

func DeleteEventCommissaire(id string, comId string, session *mgo.Session) error{
	
	var updateEvent Event
	var emptyCom []string
	defer session.Close()
	
	collection := session.DB(DATABASE).C(EVENTS)
	err := collection.Find(bson.M{"id": id}).One(&updateEvent)
	if err != nil {
		fmt.Println(err)
	}
	commissaires := updateEvent.Commissaires
	index := FindEventCommissaires(commissaires, comId)
	fmt.Println(index)
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

func DeleteEventStages(id string, stageId string, session *mgo.Session) error{
	
	var updateEvent Event
	var emptyStage []string
	defer session.Close()
	
	collection := session.DB(DATABASE).C(EVENTS)
	err := collection.Find(bson.M{"id": id}).One(&updateEvent)
	if err != nil {
		fmt.Println(err)
	}
	stages := updateEvent.Stages
	index := FindEventStage(stages, stageId)
	fmt.Println(index)
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