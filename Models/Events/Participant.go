package events

import "time"

type Participant struct {
	No				int
	Rider			Participant
	TotalTime		time.Time
	TotalPoints		int
	MountainPoints	int
	SprintPoints	int
	U23				bool
	bk				bool
	State			string
}

type Participants []Participant