package events

import (
	model "toyger/models"
	"fmt"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"errors"
)

type Participant struct {
	No				int			`json:"no"`
	Event			string		`json:"event"`
	Rider			model.Cyclist	    `json:"rider"`
	TotalTime		int 		`json:"total_time"`
	TotalPoints		int			`json:"total_points"`			
	MountainPoints	int			`json:"mountain_points"`
	SprintPoints	int			`json:"sprint_points"`
	U23				bool		`json:"u23"`
	Bk				bool		`json:"bk"`
	State			string		`json:"state"`
}

type Participants []Participant
var PARTICIPANTS = "participants"

func CreateParticipant(eventID string, p Participant, session *mgo.Session,
					  dbName string, tableName string) (Participant, error) {
	
	var participant Participant
	defer session.Close()
	existsParticipant := GetParticipantInside(eventID, p.No, session,
											 dbName, tableName)
	var err error
	if existsParticipant != nil {
		participant.No = p.No
		participant.Event = eventID
		participant.Rider = p.Rider
		participant.TotalTime = p.TotalTime
		participant.TotalPoints = p.TotalPoints
		participant.MountainPoints = p.MountainPoints
		participant.SprintPoints = p.SprintPoints
		participant.U23 = p.U23
		participant.Bk = p.Bk
		participant.State = p.State
		
		collection := session.DB(dbName).C(tableName)
		if err := collection.Insert(participant); err != nil {
			fmt.Println(err)
			return participant, err
		}
	} else {
		fmt.Println("User with the same RACE NUMBER exists")
		err = errors.New("User with the same RACE NUMBER exists")
	}
	return participant, err
}

func GetParticipantInside(eventID string, id int, session *mgo.Session,
	dbName string, tableName string) error {
	
	var p Participant
	collection := session.DB(dbName).C(tableName)
	err := collection.Find(bson.M{"no": id, "event": eventID}).One(&p)
	if err != nil {
		fmt.Println(err)
	}

	return err
}

func GetParticipantsList(eventID string, session *mgo.Session,
	dbName string, tableName string) (Participants, error) {
	
	part := Participants{}
	defer session.Close()

	c := session.DB(dbName).C(tableName)
	err := c.Find(bson.M{"event": eventID}).All(&part)
	if err != nil {
		fmt.Println(err)
	}
	return part, err
}

func GetParticipant(eventID string, id int, session *mgo.Session,
					dbName string, tableName string) (Participant, error) {
	
	var p Participant
	defer session.Close()
	collection := session.DB(dbName).C(tableName)
	err := collection.Find(bson.M{"no": id, "event": eventID}).One(&p)
	if err != nil {
		fmt.Println(err)
	}

	return p, err
}

func UpdateParticipant(eventID string, id int, p Participant, session *mgo.Session,
	dbName string, tableName string) (Participant, error) {
	
	var updateParticipant Participant

	defer session.Close()
	
	collection := session.DB(dbName).C(tableName)
	
	err := collection.Find(bson.M{"no": id, "event": eventID}).One(&updateParticipant)
	if err != nil {
		fmt.Println(err)
		return updateParticipant, err
	}
	updateParticipant.No = p.No
	updateParticipant.Event = eventID
	updateParticipant.Rider = p.Rider
	updateParticipant.TotalTime = p.TotalTime
	updateParticipant.TotalPoints = p.TotalPoints
	updateParticipant.MountainPoints = p.MountainPoints
	updateParticipant.SprintPoints = p.SprintPoints
	updateParticipant.U23 = p.U23
	updateParticipant.Bk = p.Bk
	updateParticipant.State = p.State	

	if err2 := collection.Update(bson.M{"no": id}, updateParticipant); err2 != nil {
		fmt.Println(err2)
		err = err2
	}
	
	return updateParticipant, err
}

func DeleteParticipant(eventID string, id int, session *mgo.Session,
	dbName string, tableName string) error {
	
	defer session.Close()
	
	collection := session.DB(dbName).C(tableName)
	err := collection.Remove(bson.M{"no": id, "event": eventID})
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Object was deleted")

	return err
}