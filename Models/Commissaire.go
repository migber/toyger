package models

type Commissaire struct{
	UCIID		string	`json:"uciid"`
	Name		string	`json:"name"`
	Birthdate	string	`json:"birth_date"`
	UCICategory string	`json:"uci_category"`
	Nationality	string	`json:"nationality"`
}

type Commissaires []Commissaire 