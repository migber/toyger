package events

import "time"

type Event struct {	
	ID 				string				`json:"id"`
	Name 			string				`json:"name"`
	NoParticipants	int					`json:"no_participants"`
	NoStages 		int					`json:"no_stages"`
	Date 			time.Time			`json:"date"`
	Location 		string				`json:"location"`
	NoCommissaires 	int 				`json:"no_commissaires"`
	TotalKm 		int					`json:"total_km"`
	Stages			[]Stage				`json:"stages"`
	Participants	[]Participants		`json:"participants"`
	Commissaires	[]RaceCommissaires	`json:"commissaires"`
}

type Events []Event 