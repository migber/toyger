package events

import (
	"time"
	model "toyger/models"
)

type Participant struct {
	No				int			`json:"no"`
	Rider			model.Cyclist	    `json:"rider"`
	TotalTime		time.Time	`json:"total_time"`
	TotalPoints		int			`json:"total_points"`			
	MountainPoints	int			`json:"mountain_points"`
	SprintPoints	int			`json:"sprint_points"`
	U23				bool		`json:"u23"`
	bk				bool		`json:"bk"`
	State			string		`json:"state"`
}

type Participants []Participant