package models

type Cyclist struct{
	UCIID		string		`json:"uci_id"`
	Name		string		`json:"name"`
	Surname		string		`json:"surname"`
	Team 		Team		`json:"team"`
	Coaches 	[]string	`json:"coaches"`
	Age			int			`json:"age"`
	Birthdate	string		`json:"birth_date"`
	Gender		string		`json:"gender"`
	UCICategory	string		`json:"uci_category"`
	Nationality string		`json:"nationality"`
}

type Cyclists []Cyclist	