package main

import (
	"testing"
	"fmt"
	m "toyger/models/events"
	Models "toyger/models"
)

var deleteEvenCommissairetId string

func TestEventCommissaireCreateMethod(t *testing.T){

	fmt.Println("TESTING: EVENT COMMISSAIRE CREATE")
	
	var raceCommissaire m.RaceCommissaire
	var commissaire Models.Commissaire
	commissaire.UCIID = "12345432123"
	commissaire.Birthdate = "1996-04-30"
	commissaire.Gender = "male"
	commissaire.Name = "Alex"
	commissaire.Nationality = "German"
	commissaire.UCICategory = "Nationl Commissaire"
	raceCommissaire.Commissaire = commissaire
	raceCommissaire.Event = globaleventID
	raceCommissaire.Position = []m.Position{
		{ Stage: "1", Name: "Timekeeper"},
		{ Stage: "2", Name: "Finish judge"},
	}

	createdEventCommissaire := m.CreateRaceCommissaire(globaleventID, raceCommissaire,
								 ConnectionTesting(), TESTING, RACECOMMISSAIRE)
	found := m.GetRaceCommissaire(globaleventID, createdEventCommissaire.Commissaire.UCIID,
								 ConnectionTesting(),
								 TESTING, RACECOMMISSAIRE)
	deleteEvenCommissairetId = createdEventCommissaire.Commissaire.UCIID
	if (found.Commissaire.UCIID != createdEventCommissaire.Commissaire.UCIID &&
		found.Event != createdEventCommissaire.Event && 
		found.Position[0].Name != createdEventCommissaire.Position[0].Name &&
		len(found.Position) != 0){
		t.Error("Should create race commissaire")
	}
}

func TestRaceCommissaireGetList(t *testing.T) {
	fmt.Println("TESTING: EVENT RACE COMMISSAIRE LIST")

	raceCommissairesList := m.GetRaceCommissairesList(globaleventID, ConnectionTesting(),
							  TESTING, RACECOMMISSAIRE)
	if (len(raceCommissairesList) == 0){
		t.Error("Should be not empty race commissaire list")
	}
}

func TestEventCommissaireUpdate(t *testing.T){
	
	fmt.Println("TESTING: EVENT COMMISSAIRE UPDATE")
	foundRaceCommissaire := m.GetRaceCommissaire(globaleventID, deleteEvenCommissairetId,
							  ConnectionTesting(), TESTING, RACECOMMISSAIRE)
	beforePositionName := foundRaceCommissaire.Position[0].Name
	foundRaceCommissaire.Position[0].Name = "Finish judge"
	updateRaceCommissaire := m.UpdateRaceCommissaire(globaleventID, 
							   foundRaceCommissaire.Commissaire.UCIID,
							   foundRaceCommissaire,
							   ConnectionTesting(), TESTING, RACECOMMISSAIRE)
	if(updateRaceCommissaire.Position[0].Name == beforePositionName){
		t.Error("Should not the same location name")
	}
}

func TestEventRaceCommissaireDelete(t *testing.T){

	fmt.Println("TESTING: EVENT COMMISSAIRE DELETE")
	err := m.DeleteRaceCommissaire(globaleventID, deleteEvenCommissairetId, 
								   ConnectionTesting(), TESTING, RACECOMMISSAIRE)
	if (err != nil){
		t.Error("Should delete specified race commissaire")
	}
}