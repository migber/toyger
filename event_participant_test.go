package main

import (
	"testing"
	"fmt"
	m "toyger/models/events"
	Models "toyger/models"
)

var deleteEvenParticipanttId int

func TestEventParticipantCreateMethod(t *testing.T){

	fmt.Println("TESTING: EVENT PARTICIPANT CREATE")
	var participant m.Participant
	var rider Models.Cyclist
	rider.Name = "Alexander"
	rider.Gender = "male"
	rider.Birthdate = "1998-02-03"
	rider.Nationality = "Italian"
	rider.Surname = "Amdrolevic"
	rider.Team = "fda56266-831f-4a09-947d-f63e7ca419c5"
	rider.UCICategory = "Men Elite"
	rider.UCIID = "98765432189"
	rider.Coaches = []string {"Manager I", "Manager II"}

	participant.Event = globaleventID
	participant.No = 34
	participant.Bk = false
	participant.MountainPoints = 3
	participant.SprintPoints = 15
	participant.State = ""
	participant.TotalPoints = 18
	participant.TotalTime = 1449900
	participant.U23 = false
	participant.Rider = rider

	createdParticipant := m.CreateParticipant(globaleventID, participant,
							ConnectionTesting(), TESTING, PARTICIPANTS)
	found := m.GetParticipant(globaleventID, createdParticipant.No,
							  ConnectionTesting(),
						      TESTING, PARTICIPANTS)
	deleteEvenParticipanttId = createdParticipant.No
	if (createdParticipant.No != found.No &&
		createdParticipant.Event!= found.Event && 
		createdParticipant.Rider.Name != found.Rider.Name &&
		createdParticipant.Rider.Surname != found.Rider.Surname){
		t.Error("Should create race participant")
	}
}

func TestEventParticipantGetList(t *testing.T) {
	fmt.Println("TESTING: EVENT PARTICIPANTS LIST")

	participants := m.GetParticipantsList(globaleventID, ConnectionTesting(),
					TESTING, PARTICIPANTS)
	if (len(participants) == 0){
		t.Error("Should be not empty event participant`s list")
	}
}

func TestEventParticipantUpdate(t *testing.T){
	
	fmt.Println("TESTING: EVENT PARTICIPANT UPDATE")
	foundParticipant := m.GetParticipant(globaleventID, deleteEvenParticipanttId,
						  ConnectionTesting(), TESTING, PARTICIPANTS)
	beforeTotalPoints := foundParticipant.TotalPoints
	foundParticipant.TotalPoints = 23
	updateParticipant := m.UpdateParticipant(globaleventID, 
						   foundParticipant.No,
						   foundParticipant,
						   ConnectionTesting(), TESTING, PARTICIPANTS)
	if(updateParticipant.TotalPoints == beforeTotalPoints){
		t.Error("Should not the same total points ")
	}
}

func TestEventParticipantDelete(t *testing.T){

	fmt.Println("TESTING: EVENT PARTICIPANT DELETE")
	err := m.DeleteParticipant(globaleventID, deleteEvenParticipanttId, 
							ConnectionTesting(), TESTING, PARTICIPANTS)
	if (err != nil){
		t.Error("Should delete specified participant")
	}
}