package events

type Stage struct {
	ID				string
	Name 			string
	Km				int
	Sprints			[]Sprint
	Abandoned		int
	Disqualified	int
	Starters		int
	Remaining		int
	AfterTimeLimit	int
	NotStarted		int
}

type Stages []Stage