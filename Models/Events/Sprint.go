package events

type Sprint struct {
	Id string
	Category 	string
	Winner		Participant
	Second		Participant
	Third		Participant
	Bonuses 	[]int
}

type Sprints []Sprint