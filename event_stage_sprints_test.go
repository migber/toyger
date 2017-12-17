package main

import (
	"testing"
	"fmt"
	m "toyger/models/events"
)

var deleteEvenStagetsprintID string

func TestEventStageSprintCreateMethod(t *testing.T){

	fmt.Println("TESTING: EVENT STAGE SPRINT CREATE")
	var sprint m.Sprint
	sprint.Bonuses = []int{3,2,1}
	sprint.Category = "Intermediate sprint I"
	sprint.Event = globaleventID
	sprint.Id = "064212fd-d1cd-4fbb-8e04-0538ec48ac80"
	sprint.Name = "Points I"
	sprint.Second = 23
	sprint.Stage = 1
	sprint.Third = 43
	sprint.Winner = 55

	createSprint := m.CreateSprint(globaleventID, globalstageID, sprint,
					  ConnectionTesting(), TESTING, SPRINTS)
	found := m.GetSprint(globaleventID, globalstageID,
						createSprint.Id, ConnectionTesting(),
						TESTING, SPRINTS)
	deleteEvenStagetsprintID = found.Id
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

	sprints := m.GetSprintList(globaleventID, globalstageID,
				 ConnectionTesting(),
				 TESTING, SPRINTS)
	if (len(sprints) == 0){
		t.Error("Should be not empty sprint`s list")
	}
}

func TestEventStageSprintUpdate(t *testing.T){
	
	fmt.Println("TESTING: EVENT STAGE SPRINT UPDATE")
	foundSprint := m.GetSprint(globaleventID, globalstageID, 
					 deleteEvenStagetsprintID,
					 ConnectionTesting(), TESTING, SPRINTS)
	beforeCategory := foundSprint.Category
	foundSprint.Category = "Category 3"
	updateSprint := m.UpdateSprint(globaleventID, globalstageID,
					  foundSprint.Id,
					  foundSprint,
					  ConnectionTesting(), TESTING, SPRINTS)
	if(updateSprint.Category == beforeCategory){
		t.Error("Should not the same sprint`s category ")
	}
}

func TestEventStageSprintDelete(t *testing.T){

	fmt.Println("TESTING: EVENT STAGE SPRINT`S DELETE")
	err := m.DeleteSprint(globaleventID, globalstageID, 
						 deleteEvenStagetsprintID, 
						 ConnectionTesting(), TESTING, SPRINTS)
	if (err != nil){
		t.Error("Should delete specified sprint")
	}
}