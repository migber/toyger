package events

type Sprint struct {
	Id 			string		`json:"id"`
	Category 	string		`json:"category"`
	Winner		Participant	`json:"winner"`
	Second		Participant	`json:"second"`
	Third		Participant	`json:"third"`
	Bonuses 	[]int		`json:"bonuses"`
}

type Sprints []Sprint
var SPRINTS = "sprints"

