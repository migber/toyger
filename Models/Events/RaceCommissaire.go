package events
import (
	model "toyger/models"
)

type RaceCommissaire struct{
	Commissaire 	model.Commissaire	`json:"commissaire"`
	Position		struct{		
		Stage	Stage	`json:"stage"`
		Name 	string	`json:"name"`
	} `json:"position"`
}

type RaceCommissaires []RaceCommissaire	