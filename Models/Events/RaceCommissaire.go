package events

type RaceCommissaire struct{
	Commissaire 	Commissaire
	Possition		struct{
		Stage	Stage
		Name 	string
	}
}

type RaceCommissaires []RaceCommissaire