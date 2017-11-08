package models


import "errors"

import (
	"fmt"
	"github.com/satori/go.uuid"
)

type Team struct {
	Id		uuid.UUID	`json:"id"`
	Name	string		`json:"name"`
	Manager Manager		`json:"manager"`
}

type Manager struct {
	Id 			uuid.UUID	`json:"id"`
	Name 		string		`json:"name"`
	Surname 	string		`json:"surname"`
	Birthdate	string		`json:"birth_date"`
	Phone		string		`json:"phone"`
	Nationality	string		`json:"nationality"`
}

type Teams []Team

var team Team
var teams Teams

func CreateTeam(t Team) Team {
	t.Id = uuid.NewV4()
	t.Manager.Id = uuid.NewV4()
	team.Id = t.Id
	team.Name = t.Name
	team.Manager = t.Manager
	teams = append(teams, team)
	return t
}

func GetTeamsList() Teams {
	return teams
}

func GetTeam(id string) Team {
	var t Team
	u2, err :=  uuid.FromString(id)
	if err != nil{
		fmt.Errorf("Error occured while parsing uuid %v", err)	
	}

	for _, team := range teams {
		if team.Id == u2 {
			t = team
		}
	}
	return t
}

func UpdateTeam(uid string, t Team) Team {

	updateTeam, ind := FindTeam(uid)

	updateTeam.Name = t.Name
	updateTeam.Manager.Name = t.Manager.Name
	updateTeam.Manager.Surname = t.Manager.Surname
	updateTeam.Manager.Birthdate = t.Manager.Birthdate
	updateTeam.Manager.Phone = t.Manager.Phone
	updateTeam.Manager.Nationality = t.Manager.Nationality

	teams[ind] = updateTeam

	return updateTeam
}

func DeleteTeam(uid string) error {

	delete := -1
	u2, err :=  uuid.FromString(uid)
	if err != nil{
		fmt.Errorf("Error occured while parsing uuid %v", err)	
	}

	for ind, team := range teams {
		if team.Id == u2 {
			delete = ind
		}
	}

	if delete != -1{
		teams = append(teams[:delete], teams[delete + 1:]...)
		return nil
	}

	return errors.New("Could not find team")
}

func FindTeam(id string) (Team, int) {

	var t Team
	var index int
	u2, err :=  uuid.FromString(id)
	if err != nil{
		fmt.Errorf("Error occured while parsing uuid %v", err)	
	}

	for ind, team := range teams {
		if team.Id == u2 {
			t = team
			index = ind
		}
	}
	return t, index
}