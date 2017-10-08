package models

import "time"

type Cyclist struct{
	UCIID	string
	Name	string
	Surname	string
	Team 	Team
	Coach 	string
	Age		int
	Birthdate	time.Date
	Gender	string
	UCICategory	string
	Nationality string
}

type Cyclists []Cyclist