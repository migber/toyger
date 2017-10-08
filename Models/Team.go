package models

import "time"

type Team struct {
	Id		string
	Name	string
	Manager Manager
}

type Manager struct {
	Id 			int
	Name 		string
	Surname 	string
	Birthdate	time.Date
	Phone		string
	Nationality	string
}

type Teams []Team