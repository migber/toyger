package main

import (
	"testing"
	"github.com/koding/multiconfig"
	"gopkg.in/mgo.v2"
	"time"
	"log"
	"fmt"
	m "toyger/models/events"
)

const TESTING = "toygertesting"
const globaleventID = "33baff49-23a3-40bd-ad5a-f2913865e505"
var deleteeventID string

func ConnectionTesting() *mgo.Session {

	m := multiconfig.NewWithPath("./config/config.test.json")
	db := new(DB)
	m.MustLoad(db)

	mongoDBDialInfo := &mgo.DialInfo{
		Addrs:    []string{db.Host},
		Timeout:  60 * time.Second,
		Database: db.Database,
		Username: db.Username,
		Password: db.Password,
	}	

	mongoSession, err := mgo.DialWithInfo(mongoDBDialInfo)
	if err != nil {
		log.Fatalf("CreateSession: %s\n", err)
	}

	return mongoSession.Clone()
}

func TestEventCreateMethod(t *testing.T){

	fmt.Println("TESTING: EVENT CREATE")

	var event m.Event
	event.Name = "Plento varzybos 5554"
	event.NoParticipants = 10
	event.NoStages = 2
	event.Date = "2017-04-12"
	event.Location = "Panevezys"
	event.NoCommissaires = 10
	event.TotalKm = 180
	event.Stages = []int{1,2,3}
	event.Participants = []int{1000234567, 11102223332}
	event.Commissaires = []string{}

	createdEvent := m.CreateEvent(event, ConnectionTesting(), TESTING, EVENTS)
	deleteeventID = createdEvent.ID
	found := m.GetEvent(createdEvent.ID, ConnectionTesting(), TESTING, EVENTS)
	if (found.Name != event.Name){
		t.Error("Should be the same event name")
	}
}

func TestEventGetList(t *testing.T) {
	fmt.Println("TESTING: EVENT GET LIST")

	eventList := m.GetEventsList(ConnectionTesting(), TESTING, EVENTS)
	if (len(eventList) == 0){
		t.Error("Should be not empty event list")
	}
}

func TestEventUpdate(t *testing.T){
	
	fmt.Println("TESTING: EVENT UPDATE")
	foundEvent := m.GetEvent(deleteeventID, ConnectionTesting(), TESTING, EVENTS)
	beforeLocation := foundEvent.Location
	foundEvent.Location = "New event location"
	updatedEvent := m.UpdateEvent(foundEvent.ID, foundEvent, ConnectionTesting(), TESTING, EVENTS)
	if(updatedEvent.Location == beforeLocation){
		t.Error("Should not the same location name")
	}
}

func TestEventParticipantInsertion(t *testing.T){
	fmt.Println("TESTING: EVENT PARTICIPANT INSERT INSIDE")
	riderNo:= 100
	err := m.InsertEventParticipants(deleteeventID, riderNo, ConnectionTesting(), TESTING, EVENTS)
	if err != nil{
		t.Error("Should update Participants list in event")
	}
	getEvent := m.GetEvent(deleteeventID, ConnectionTesting(), TESTING, EVENTS)
	index := m.FindEventsParticipant(getEvent.Participants, riderNo)
	if(getEvent.Participants[index] != riderNo){
		t.Error("Should found inserted rider")
	}
}

func TestEventParticipantDeletion(t *testing.T){
	fmt.Println("TESTING: EVENT PARTICIPANT DELETE INSIDE")
	riderNo:= 100
	err := m.DeleteEventParticipant(deleteeventID, riderNo, ConnectionTesting(), TESTING, EVENTS)
	if err != nil{
		t.Error("Should update Participants list in event")
	}
	getEvent := m.GetEvent(deleteeventID, ConnectionTesting(), TESTING, EVENTS)
	index := m.FindEventsParticipant(getEvent.Participants, riderNo)
	if(index != -1){
		t.Error("Should not found rider")
	}
}

func TestEventCommissaireInsertion(t *testing.T){
	fmt.Println("TESTING: EVENT COMMISSAIRE INSERT INSIDE")
	commissaireID:= "12345678912"
	err := m.InsertEventCommissaire(deleteeventID, commissaireID, ConnectionTesting(), TESTING, EVENTS)
	if err != nil{
		t.Error("Should update Commissaires list in event")
	}
	getEvent := m.GetEvent(deleteeventID, ConnectionTesting(), TESTING, EVENTS)
	index := m.FindEventCommissaires(getEvent.Commissaires, commissaireID)
	if(getEvent.Commissaires[index] != commissaireID){
		t.Error("Should found inserted commissaire")
	}
}

func TestEventCommissaireDeletion(t *testing.T){
	fmt.Println("TESTING: EVENT COMMISSAIRE DELETE INSIDE")
	commissaireID:= "12345678912"
	err := m.DeleteEventCommissaire(deleteeventID, commissaireID, ConnectionTesting(), TESTING, EVENTS)
	if err != nil{
		t.Error("Should update Commissaires list in event")
	}
	getEvent := m.GetEvent(deleteeventID, ConnectionTesting(), TESTING, EVENTS)
	index := m.FindEventCommissaires(getEvent.Commissaires, commissaireID)
	if(index!= -1){
		t.Error("Should not found commissaire")
	}
}

func TestEventStagesInsertion(t *testing.T){
	fmt.Println("TESTING: EVENT STAGE INSERT INSIDE")
	stageNr:= 3
	err := m.InsertEventStages(deleteeventID, stageNr, ConnectionTesting(), TESTING, EVENTS)
	if err != nil{
		t.Error("Should update Stages list in event")
	}
	getEvent := m.GetEvent(deleteeventID, ConnectionTesting(), TESTING, EVENTS)
	index := m.FindEventStage(getEvent.Stages, stageNr)
	if(getEvent.Stages[index] != stageNr ){
		t.Error("Should found inserted stage")
	}
}

func TestEventStagesDeletion(t *testing.T){
	fmt.Println("TESTING: EVENT STAGE DELETE INSIDE")
	stageNr:= 2
	err := m.DeleteEventStage(deleteeventID, stageNr, ConnectionTesting(), TESTING, EVENTS)
	if err != nil{
		t.Error("Should update Stages list in event")
	}
	getEvent := m.GetEvent(deleteeventID, ConnectionTesting(), TESTING, EVENTS)
	index := m.FindEventStage(getEvent.Stages, stageNr)
	if(index != -1){
		t.Error("Should not found stage")
	}
}

func TestEventDelete(t *testing.T){

	fmt.Println("TESTING: EVENT DELETE INSIDE")
	err := m.DeleteEvent(deleteeventID, ConnectionTesting(), TESTING, EVENTS)
	
	if (err != nil){
		t.Error("Should delete specified event")
	}
}