package events

import (
	"time"
	model "toyger/models"
	"fmt"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type Participant struct {
	No				int			`json:"no"`
	Event			string		`json:"event"`
	Rider			model.Cyclist	    `json:"rider"`
	TotalTime		time.Time	`json:"total_time"`
	TotalPoints		int			`json:"total_points"`			
	MountainPoints	int			`json:"mountain_points"`
	SprintPoints	int			`json:"sprint_points"`
	U23				bool		`json:"u23"`
	Bk				bool		`json:"bk"`
	State			string		`json:"state"`
}

type Participants []Participant
var PARTICIPANTS = "participants"

func CreateParticipant(eventId string, p Participant, session *mgo.Session) Participant {
	
	var participant Participant
	defer session.Close() 

	existsParticipant := GetParticipantInside(eventId, p.No, session)
	if existsParticipant != nil {
		participant.No = p.No
		participant.Event = eventId
		participant.Rider = p.Rider
		participant.TotalTime = p.TotalTime
		participant.TotalPoints = p.TotalPoints
		participant.MountainPoints = p.MountainPoints
		participant.SprintPoints = p.SprintPoints
		participant.U23 = p.U23
		participant.Bk = p.Bk
		participant.State = p.State
		
		collection := session.DB(DATABASE).C(PARTICIPANTS)
		if err := collection.Insert(participant); err != nil {
			panic(err)
		}
	} else {
		fmt.Println("User with the same RACE NUMBER exists")
	}
	return participant
}

func GetParticipantInside(eventId string, id int, session *mgo.Session) error {
	
	var p Participant
	collection := session.DB(DATABASE).C(PARTICIPANTS)
	err := collection.Find(bson.M{"no": id, "event": eventId}).One(&p)
	if err != nil {
		fmt.Println(err)
	}

	return err
}

func GetParticipantsList(eventId string, session *mgo.Session) Participants {
	
	part := Participants{}
	defer session.Close()

	c := session.DB(DATABASE).C(PARTICIPANTS)
	err := c.Find(nil).All(&part)
	if err != nil {
		panic(err)
	}
	return part
}

func GetParticipant(eventId string, id int, session *mgo.Session) Participant {
	
	var p Participant
	defer session.Close()
	collection := session.DB(DATABASE).C(PARTICIPANTS)
	err := collection.Find(bson.M{"no": id, "event": eventId}).One(&p)
	if err != nil {
		fmt.Println(err)
	}

	return p
}

func UpdateParticipant(eventId string, id int, p Participant, session *mgo.Session) Participant {
	
	var updateParticipant Participant

	defer session.Close()
	
	collection := session.DB(DATABASE).C(PARTICIPANTS)

	err := collection.Find(bson.M{"no": id, "event": eventId}).One(&updateParticipant)
	if err != nil {
		panic(err)
	}
	updateParticipant.No = p.No
	updateParticipant.Event = eventId
	updateParticipant.Rider = p.Rider
	updateParticipant.TotalTime = p.TotalTime
	updateParticipant.TotalPoints = p.TotalPoints
	updateParticipant.MountainPoints = p.MountainPoints
	updateParticipant.SprintPoints = p.SprintPoints
	updateParticipant.U23 = p.U23
	updateParticipant.Bk = p.Bk
	updateParticipant.State = p.State	

	if err := collection.Update(bson.M{"no": id}, updateParticipant); err != nil {
		panic(err)
	}
	
	return updateParticipant
}

func DeleteParticipant(eventId string, id int, session *mgo.Session) error {
	
	defer session.Close()
	
	collection := session.DB(DATABASE).C(PARTICIPANTS)
	err := collection.Remove(bson.M{"no": id, "event": eventId})
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Object was deleted")

	return err
}