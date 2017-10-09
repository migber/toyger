package events

type RaceCommissaire struct{
	Commissaire 	Commissaire	`json:"commissaire"`
	Position		struct{		`json:"position"`
		Stage	Stage	`json:"stage"`
		Name 	string	`json:"name"`
	}
}

type RaceCommissaires []RaceCommissaire	