package main

import (
	"testing"
	"fmt"
	m "toyger/models"
)

var deleteTeamId string

func TestTeamCreateMethod(t *testing.T){

	fmt.Println("TESTING: TEAM CREATE")
	var team m.Team
	var manager m.Manager
	manager.Name = "TEAM SKY MANAGER I"
	manager.Nationality = "German"
	manager.Phone = "+45633452235"
	manager.Surname = "SURNAME Manager I"
	manager.Birthdate = "1865-04-22"
	team.Manager = manager
	team.Name = "TEAM SKY"

	createTeam := m.CreateTeam(team, ConnectionTesting(), TESTING, TEAMS)
	found := m.GetTeam(createTeam.Id,
					   ConnectionTesting(),
					   TESTING, TEAMS)
	deleteTeamId = found.Id
	if (createTeam.Id != found.Id &&
		createTeam.Name != found.Name && 
		createTeam.Manager.Name != found.Manager.Name &&
		createTeam.Manager.Nationality != found.Manager.Nationality){
		t.Error("Should create Team")
	}
}

func TestTeamsGetList(t *testing.T) {
	fmt.Println("TESTING: TEAMS LIST")

	teams := m.GetTeamsList(ConnectionTesting(),
							TESTING, TEAMS)
	if (len(teams) == 0){
		t.Error("Should be not empty teams list")
	}
}

func TestTeamUpdate(t *testing.T){
	
	fmt.Println("TESTING: TEAM UPDATE")
	foundTeam := m.GetTeam(deleteTeamId,
				  ConnectionTesting(), TESTING, TEAMS)
	beforeName := foundTeam.Name
	foundTeam.Name = "Team Name Updated"
	updateTeam := m.UpdateTeam(foundTeam.Id,
					foundTeam,
					ConnectionTesting(), TESTING, TEAMS)
	if(updateTeam.Name == beforeName){
		t.Error("Should not the same team name ")
	}
}

func TestTeamsRidersInsertion(t *testing.T){
	fmt.Println("TESTING: TEAMS RIDERS INSERT INSIDE")
	riderId:= "12345678912"
	err := m.InsertRider(deleteTeamId, riderId, 
			 ConnectionTesting(), TESTING, TEAMS)
	if err != nil{
		t.Error("Should update rider in teams riders` list")
	}
	getTeam := m.GetTeam(deleteTeamId, 
				ConnectionTesting(), TESTING, TEAMS)
	index := m.FindCyclist(getTeam.Riders, riderId)
	if(getTeam.Riders[index] != riderId){
		t.Error("Should found inserted rider")
	}
}

func TestTeamsRidersDeletion(t *testing.T){
	fmt.Println("TESTING: TEAMS RIDERS DELETE INSIDE")
	riderId:= "12345678912"
	err := m.DeleteRider(deleteTeamId, riderId,
			 ConnectionTesting(), TESTING, TEAMS)
	if err != nil{
		t.Error("Should update riders list in team")
	}
	getTeam := m.GetTeam(deleteTeamId, 
				ConnectionTesting(), TESTING, TEAMS)
	index := m.FindCyclist(getTeam.Riders, riderId)
	if(index!= -1){
		t.Error("Should not found rider")
	}
}

func TestTeamDelete(t *testing.T){

	fmt.Println("TESTING: TEAM DELETE")
	err := m.DeleteTeam(deleteTeamId, 
						ConnectionTesting(), TESTING, TEAMS)
	if (err != nil){
		t.Error("Should delete specified team")
	}
}