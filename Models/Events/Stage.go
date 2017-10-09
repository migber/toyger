package events

type Stage struct {
	ID				string		`json:"id"`
	Name 			string		`json:"name	"`
	Km				int			`json:"km"`
	Sprints			[]Sprint	`json:"sprints"`
	Abandoned		int			`json:"abandoned"`
	Disqualified	int			`json:"disqualfied"`
	Starters		int			`json:"starters"`
	Remaining		int			`json:"remaining"`
	AfterTimeLimit	int			`json:"after_time_limit"`
	NotStarted		int			`json:"not_started"`
}

type Stages []Stage