package models

import (
	"fmt"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"	
	"errors"
)

type Commissaire struct{
	UCIID		string	`json:"uciid"`
	Name		string	`json:"name"`
	Birthdate	string	`json:"birth_date"`
	Gender		string 	`json:"gender"`
	UCICategory string	`json:"uci_category"`
	Nationality	string	`json:"nationality"`
}

type Commissaires []Commissaire
var COMMISSAIRES = "commissaires"

func CreateCommissaire(c Commissaire, session *mgo.Session) (Commissaire, error) {
	fmt.Println("Inside commissaire")
	var com Commissaire
	defer session.Close()
	
	existsCom := GetCommissaireInside(c.UCIID, session)
	var err error
	if existsCom != nil {
		com.UCIID = c.UCIID
		com.Name = c.Name
		com.Birthdate = c.Birthdate
		com.Gender = c.Gender
		com.UCICategory = c.UCICategory
		com.Nationality = c.Nationality
		collection := session.DB(DATABASE).C(COMMISSAIRES)	
		if err = collection.Insert(com); err != nil {
			fmt.Println(err)
		}
	} else {
		fmt.Println("Commissaire with UCIID already exists")
		err = errors.New("Commissaire with UCIID already exists")
	}
	fmt.Println(com)
	return com, err
}
func GetCommissaireInside(id string, session *mgo.Session) error {
	
		var c Commissaire
		collection := session.DB(DATABASE).C(COMMISSAIRES)
		err := collection.Find(bson.M{"uciid": id}).One(&c)
		if err != nil {
			fmt.Println(err)
		}
		return err
}

func GetCommissairesList(session *mgo.Session) (Commissaires, error) {
	
	commissaires := Commissaires{}
	defer session.Close()
	c := session.DB(DATABASE).C(COMMISSAIRES)
	err := c.Find(nil).All(&commissaires)
	if err != nil {
		fmt.Println(err)
	}
	return commissaires, err
}

func GetCommissaire(id string, session *mgo.Session) (Commissaire, error) {
	
	var c Commissaire
	defer session.Close()
	collection := session.DB(DATABASE).C(COMMISSAIRES)
	err := collection.Find(bson.M{"uciid": id}).One(&c)
	if err != nil {
		fmt.Println(err)
	}
	return c, err
}


func UpdateCommissaire(uid string, c Commissaire, session *mgo.Session) (Commissaire, error) {
	
	var updateC Commissaire

	defer session.Close()
	
	collection := session.DB(DATABASE).C(COMMISSAIRES)

	err := collection.Find(bson.M{"uciid": uid}).One(&updateC)
	if err != nil {
		fmt.Println(err)
		return updateC, err
	}

	updateC.UCIID = c.UCIID
	updateC.Name = c.Name
	updateC.Birthdate = c.Birthdate
	updateC.Gender = c.Gender
	updateC.UCICategory = c.UCICategory
	updateC.Nationality = c.Nationality

	if err := collection.Update(bson.M{"uciid": uid}, updateC); err != nil {
		fmt.Println(err)
	}
	return updateC, err
}

func DeleteCommissaire(uid string, session *mgo.Session) error {
	
	defer session.Close()
	
	collection := session.DB(DATABASE).C(COMMISSAIRES)
	err := collection.Remove(bson.M{"uciid": uid})
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Object was deleted")

	return err
}
