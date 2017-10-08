package events

import "time"

type Event struct {
	ID 				string
	Name 			string
	NoParticipants	int
	NoStages 		int
	Date 			time.Date
	Location 		string
	NoCommissaires 	int 
	TotalKm 		int
	Stages			[]Stage
	Participants	[]Participants
	Commissaires	[]RaceCommissaires
}

type Events []Event