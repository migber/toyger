package main

import (
	"testing"
	"fmt"
	m "toyger/models/events"
)

const globalStageId = 1
var deleteEvenStagetId int

func TestEventStageCreateMethod(t *testing.T){

	fmt.Println("TESTING: EVENT STAGE CREATE")
	var stage m.Stage
	stage.Abandoned = 0
	stage.AfterTimeLimit = 0
	stage.Disqualified = 1
	stage.Event = globalEventId
	stage.ID = 2
	stage.Km = 102
	stage.Name = "Stage I"
	stage.NotStarted = 0
	stage.Remaining = 99
	stage.Sprints = []m.Sprint{}
	stage.Starters = 100
	stage.TimeLimit = 8

	createStage := m.CreateStage(globalEventId, stage,
					 ConnectionTesting(), TESTING, STAGES)
	found := m.GetStage(globalEventId, stage.ID,
						ConnectionTesting(),
						TESTING, STAGES)
	deleteEvenStagetId = found.ID
	if (createStage.ID != found.ID &&
		createStage.Event!= found.Event && 
		createStage.Disqualified != found.Disqualified &&
		createStage.Starters != found.Starters){
		t.Error("Should create stage")
	}
}

func TestEventStageGetList(t *testing.T) {
	fmt.Println("TESTING: EVENT STAGE LIST")

	stages := m.GetStageList(globalEventId, ConnectionTesting(),
							TESTING, STAGES)
	if (len(stages) == 0){
		t.Error("Should be not empty stage list")
	}
}

func TestEventStageUpdate(t *testing.T){
	
	fmt.Println("TESTING: EVENT STAGE UPDATE")
	foundStage := m.GetStage(globalEventId, deleteEvenStagetId,
							 ConnectionTesting(), TESTING, STAGES)
	beforeRemaining := foundStage.Remaining
	foundStage.Remaining = 97
	updateStage := m.UpdateStage(globalEventId, foundStage.ID,
								foundStage,
								ConnectionTesting(), TESTING, STAGES)
	if(updateStage.Remaining == beforeRemaining){
		t.Error("Should not the same remaining riders number ")
	}
}

func TestEventStageDelete(t *testing.T){

	fmt.Println("TESTING: EVENT STAGE DELETE")
	err := m.DeleteStage(globalEventId, deleteEvenStagetId, 
						 ConnectionTesting(), TESTING, STAGES)
	if (err != nil){
		t.Error("Should delete specified stage")
	}
}