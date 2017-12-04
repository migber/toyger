package main

import (
	"testing"
	"fmt"
	m "toyger/models/events"
)

var deleteEvenStagetSprintId string

func TestEventStageSprintCreateMethod(t *testing.T){

	fmt.Println("TESTING: EVENT STAGE SPRINT CREATE")
	var sprint m.Sprint
	sprint.Bonuses = []int{3,2,1}
	sprint.Category = "Intermediate sprint I"
	sprint.Event = globalEventId
	sprint.Id = "064212fd-d1cd-4fbb-8e04-0538ec48ac80"
	sprint.Name = "Points I"
	sprint.Second = 23
	sprint.Stage = 1
	sprint.Third = 43
	sprint.Winner = 55

	createSprint := m.CreateSprint(globalEventId, globalStageId, sprint,
								ConnectionTesting(), TESTING, SPRINTS)
	found := m.GetSprint(globalEventId, globalStageId,
						createSprint.Id, ConnectionTesting(),
						TESTING, SPRINTS)
	deleteEvenStagetSprintId = found.Id
	if (createSprint.Id != found.Id &&
		createSprint.Event!= found.Event &&
		createSprint.Stage != found.Stage && 
		createSprint.Category != found.Category &&
		createSprint.Name != found.Name){
		t.Error("Should create sprint")
	}
}

func TestEventStageSprintsGetList(t *testing.T) {
	fmt.Println("TESTING: EVENT STAGE SPRINT`S LIST")

	sprints := m.GetSprintList(globalEventId, globalStageId, ConnectionTesting(),
							TESTING, SPRINTS)
	if (len(sprints) == 0){
		t.Error("Should be not empty sprint`s list")
	}
}

func TestEventStageSprintUpdate(t *testing.T){
	
	fmt.Println("TESTING: EVENT STAGE SPRINT UPDATE")
	foundSprint := m.GetSprint(globalEventId, globalStageId, deleteEvenStagetSprintId,
							 ConnectionTesting(), TESTING, SPRINTS)
	beforeCategory := foundSprint.Category
	foundSprint.Category = "Category 3"
	updateSprint := m.UpdateSprint(globalEventId, globalStageId, foundSprint.Id,
								   foundSprint,
								   ConnectionTesting(), TESTING, SPRINTS)
	if(updateSprint.Category == beforeCategory){
		t.Error("Should not the same sprint`s category ")
	}
}

func TestEventStageSprintDelete(t *testing.T){

	fmt.Println("TESTING: EVENT STAGE SPRINT`S DELETE")
	err := m.DeleteSprint(globalEventId, globalStageId, deleteEvenStagetSprintId, 
						 ConnectionTesting(), TESTING, SPRINTS)
	if (err != nil){
		t.Error("Should delete specified sprint")
	}
}