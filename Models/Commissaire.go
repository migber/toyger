package models

import "time"

type Commissaire struct{
	UCIID		string
	Name		string
	Birthdate	time.Date
	UCICategory string
	Nationality	string
}

type Commissaires []Commissaire