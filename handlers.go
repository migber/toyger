
package main

import (
    "encoding/json"
    "fmt"
	"net/http"
	"io/ioutil"
	"strconv"
	"html"
	Models "toyger/models"
	m "toyger/models/events"
	"github.com/gorilla/mux"
)
// HealtchCheck handler
func HealtchCheck(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hey I am working here %s!", html.EscapeString(r.URL.Path))
}
// Hanlder main handler
func Hanlder(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
}
// TeamHandler return one team handler
func TeamHandler(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	teamID := vars["teamId"]

	w.Header().Set("Content-Type", "application/json;charset=UTF-8")

	if !IsValidUUID(teamID) {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	team := Models.GetTeam(teamID, connection())
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(team); err != nil {
        http.Error(w, err.Error(), 500)
		return
	}
}
// TeamsHandler return list of teams
func TeamsHandler(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json;charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(Models.GetTeamsList(connection())); err != nil {
		panic(err)
	}
}

// TeamCreateHandler create handler
func TeamCreateHandler(w http.ResponseWriter, r *http.Request){
	
	var team Models.Team
   	body, err := ioutil.ReadAll(r.Body)
    if err != nil {
		http.Error(w, err.Error(), 500)
	}
    if err := json.Unmarshal(body, &team); err != nil {
        w.Header().Set("Content-Type", "application/json; charset=UTF-8")
        w.WriteHeader(http.StatusBadRequest) // unprocessable entity
        if err := json.NewEncoder(w).Encode(err); err != nil {
            http.Error(w, err.Error(), 500)
			return
        }
	}
	t := Models.CreateTeam(team, connection())
    w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusCreated)

	fmt.Println(r.URL.String())
	w.Header().Set("X-Frame-Options", "soooo" )
    if err := json.NewEncoder(w).Encode(t); err != nil {
        http.Error(w, err.Error(), 500)
		return
	}
	
	if err := r.Body.Close(); err != nil {
        panic(err)
	}
}

// TeamDeleteHandler team delete handler
func TeamDeleteHandler(w http.ResponseWriter, r *http.Request){

	vars := mux.Vars(r)
	teamID := vars["teamId"]

	w.Header().Set("Content-Type", "application/json;charset=UTF-8")

	if !IsValidUUID(teamID) {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err := Models.DeleteTeam(teamID, connection())
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
	}
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(Models.GetTeamsList(connection())); err != nil {
        http.Error(w, err.Error(), 500)
		return
	}
}
// TeamUpdateHandler udpate team handler
func TeamUpdateHandler(w http.ResponseWriter, r *http.Request){

	vars := mux.Vars(r)
	teamID := vars["teamId"]

	w.Header().Set("Content-Type", "application/json;charset=UTF-8")

	if !IsValidUUID(teamID) {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var team Models.Team

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), 500)
	}

	if err := r.Body.Close(); err != nil {
        panic(err)
	}

	if err := json.Unmarshal(body, &team); err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusBadRequest) // unprocessable entity
		if err := json.NewEncoder(w).Encode(err); err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
	}

	t := Models.UpdateTeam(teamID, team, connection())
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	w.Header().Set("Location", r.URL.String())
    w.WriteHeader(http.StatusOK)
    if err := json.NewEncoder(w).Encode(t); err != nil {
        http.Error(w, err.Error(), 500)
		return
	}
}

// CyclistsHandler handler cyclists return cyclists
func CyclistsHandler(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	teamID := vars["teamId"]

	w.Header().Set("Content-Type", "application/json;charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(Models.GetCyclistsList(teamID, connection())); err != nil {
		panic(err)
	}
}

//CyclistHandler handler cyclist return on cyclist
func CyclistHandler(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	teamID := vars["teamId"]
	cyclistID := vars["cyclistId"]

	w.Header().Set("Content-Type", "application/json;charset=UTF-8")

	if !IsValidUUID(teamID) {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if !isValidUCIID(cyclistID){
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	cyclist := Models.GetCyclist(teamID, cyclistID, connection())
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(cyclist); err != nil {
        http.Error(w, err.Error(), 500)
		return
	}
}

//CyclistCreateHandler handler create cyclist
func CyclistCreateHandler(w http.ResponseWriter, r *http.Request) {
	
	vars := mux.Vars(r)
	teamID := vars["teamId"]
	var cyclist Models.Cyclist
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), 500)
	}

	if err := json.Unmarshal(body, &cyclist); err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusBadRequest) // unprocessable entity
		if err := json.NewEncoder(w).Encode(err); err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
	}
	
	c := Models.CreateCyclist(teamID, cyclist, connection())
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	if c.UCIID == ""{
		w.WriteHeader(http.StatusConflict)
	} else {
		w.WriteHeader(http.StatusOK)	
		updateErr := Models.InsertRider(teamID, c.UCIID, connection())
		if updateErr != nil {
			fmt.Println("Something went wrong.")
		}
		if err := json.NewEncoder(w).Encode(c); err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
	}
	
	fmt.Println(r.URL.String())
	//w.Header().Set("X-Frame-Options", "soooo" )
	if err := r.Body.Close(); err != nil {
		panic(err)
	}
}

//CyclistUpdateHandler handler update cyclist
func CyclistUpdateHandler(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	teamID := vars["teamId"]
	cyclistID := vars["cyclistId"]
	w.Header().Set("Content-Type", "application/json;charset=UTF-8")

	if !IsValidUUID(teamID) {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if !isValidUCIID(cyclistID){
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	var cyclist Models.Cyclist
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), 500)
	}
	if err := r.Body.Close(); err != nil {
        panic(err)
	}

	if err := json.Unmarshal(body, &cyclist); err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusBadRequest) // unprocessable entity
		if err := json.NewEncoder(w).Encode(err); err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
	}

	c := Models.UpdateCyclist(teamID, cyclistID, cyclist, connection())
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	w.Header().Set("Location", r.URL.String())
    w.WriteHeader(http.StatusOK)
    if err := json.NewEncoder(w).Encode(c); err != nil {
        http.Error(w, err.Error(), 500)
		return
	}
}

//CyclistDeleteHandler handler delete cyclist
func CyclistDeleteHandler(w http.ResponseWriter, r *http.Request) {
	
	vars := mux.Vars(r)
	teamID := vars["teamId"]
	cyclistID := vars["cyclistId"]
	w.Header().Set("Content-Type", "application/json;charset=UTF-8")

	if !IsValidUUID(teamID) {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if !isValidUCIID(cyclistID){
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err := Models.DeleteCyclist(teamID, cyclistID, connection())
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
	}
	err = Models.DeleteRider(teamID, cyclistID, connection())
	if err != nil{
		fmt.Println(fmt.Errorf("%s", err))
	}
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(Models.GetCyclistsList(teamID, connection())); err != nil {
        http.Error(w, err.Error(), 500)
		return
	}
}



// CommissairesHandler handler commissaires
func CommissairesHandler(w http.ResponseWriter, r *http.Request) {
	
	w.Header().Set("Content-Type", "application/json;charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(Models.GetCommissairesList(connection())); err != nil {
		panic(err)
	}
}

// CommissaireHandler handler commissaire information
func CommissaireHandler(w http.ResponseWriter, r *http.Request) {
	
	vars := mux.Vars(r)
	comId := vars["commissaireId"]
	fmt.Println(comId)

	w.Header().Set("Content-Type", "application/json;charset=UTF-8")

	if !isValidUCIID(comId) {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	commissaire := Models.GetCommissaire(comId, connection())
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(commissaire); err != nil {
        http.Error(w, err.Error(), 500)
		return
	}
}

// CommissaireCreateHandler handler create commissaire
func CommissaireCreateHandler(w http.ResponseWriter, r *http.Request) {
	
	var commissaire Models.Commissaire
	body, err := ioutil.ReadAll(r.Body)
 	if err != nil {
	 	http.Error(w, err.Error(), 500)
 	}
 	if err := json.Unmarshal(body, &commissaire); err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	 	w.WriteHeader(http.StatusBadRequest) // unprocessable entity
	 	if err := json.NewEncoder(w).Encode(err); err != nil {
			 http.Error(w, err.Error(), 500)
			 return
			}
		}
	fmt.Println(commissaire.UCIID)
	if !isValidUCIID(commissaire.UCIID) {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	c := Models.CreateCommissaire(commissaire, connection())
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	if c.UCIID == "" {
		w.WriteHeader(http.StatusConflict)
	} else {
		w.WriteHeader(http.StatusCreated)	
	}
	fmt.Println(r.URL.String())
	if err := json.NewEncoder(w).Encode(c); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	if err := r.Body.Close(); err != nil {
		panic(err)
	}
}

// CommissaireUpdateHandler handler update commissaire information
func CommissaireUpdateHandler(w http.ResponseWriter, r *http.Request) {
	
	vars := mux.Vars(r)
	comId := vars["commissaireId"]
	
	w.Header().Set("Content-Type", "application/json;charset=UTF-8")

	if !isValidUCIID(comId) {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var commissaire Models.Commissaire

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), 500)
	}

	if err := r.Body.Close(); err != nil {
        panic(err)
	}

	if err := json.Unmarshal(body, &commissaire); err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusBadRequest) // unprocessable entity
		if err := json.NewEncoder(w).Encode(err); err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
	}
	c := Models.UpdateCommissaire(comId, commissaire, connection())
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	w.Header().Set("Location", r.URL.String())
    w.WriteHeader(http.StatusOK)
    if err := json.NewEncoder(w).Encode(c); err != nil {
        http.Error(w, err.Error(), 500)
		return
	}
}
// CommissaireDeleteHandler handler delete commissaires
func CommissaireDeleteHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	comId := vars["commissaireId"]

	w.Header().Set("Content-Type", "application/json;charset=UTF-8")

	if !isValidUCIID(comId) {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err := Models.DeleteCommissaire(comId, connection())
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
	}
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(Models.GetCommissairesList(connection())); err != nil {
        http.Error(w, err.Error(), 500)
		return
	}
}


// ManagersHandler handler managers information
func ManagersHandler(w http.ResponseWriter, r *http.Request) {
	
}

// ManagerHandler handler manager information
func ManagerHandler(w http.ResponseWriter, r *http.Request) {
	
}
// ManagerCreateHandler handler create manager
func ManagerCreateHandler(w http.ResponseWriter, r *http.Request) {
	
}
// ManagerUpdateHandler handler update manager information
func ManagerUpdateHandler(w http.ResponseWriter, r *http.Request) {
	
}
// ManagerDeleteHandler handler delete manager 
func ManagerDeleteHandler(w http.ResponseWriter, r *http.Request) {
	
}


// Event dependence

// EventsHandler handler events information
func EventsHandler(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json;charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(m.GetEventsList(connection())); err != nil {
		panic(err)
	}
}

// EventHandler handler event information
func EventHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	eventId := vars["eventId"]

	w.Header().Set("Content-Type", "application/json;charset=UTF-8")

	if !IsValidUUID(eventId) {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	event := m.GetEvent(eventId, connection())
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(event); err != nil {
        http.Error(w, err.Error(), 500)
		return
	}
}

// EventCreateHandler handler create event information
func EventCreateHandler(w http.ResponseWriter, r *http.Request) {
	
	var event m.Event
	body, err := ioutil.ReadAll(r.Body)
 	if err != nil {
		 http.Error(w, err.Error(), 500)
 	}
 	if err := json.Unmarshal(body, &event); err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusBadRequest) // unprocessable entity
		if err := json.NewEncoder(w).Encode(err); err != nil {
			http.Error(w, err.Error(), 500)
			return
	 	}
 	}
 	e := m.CreateEvent(event, connection())
 	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
 	w.WriteHeader(http.StatusCreated)

 	fmt.Println(r.URL.String())
 	w.Header().Set("X-Frame-Options", "soooo" )
 	if err := json.NewEncoder(w).Encode(e); err != nil {
		 http.Error(w, err.Error(), 500)
		 return
		}
 
 	if err := r.Body.Close(); err != nil {
		 panic(err)
	}
}

// EventDeleteHandler handler delete event information
func EventDeleteHandler(w http.ResponseWriter, r *http.Request) {
	
	vars := mux.Vars(r)
	eventId := vars["eventId"]

	w.Header().Set("Content-Type", "application/json;charset=UTF-8")

	if !IsValidUUID(eventId) {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err := m.DeleteEvent(eventId, connection())
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
	}
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(m.GetEventsList(connection())); err != nil {
        http.Error(w, err.Error(), 500)
		return
	}
}

// EventUpdateHandler handler update event information
func EventUpdateHandler(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	eventId := vars["eventId"]

	w.Header().Set("Content-Type", "application/json;charset=UTF-8")

	if !IsValidUUID(eventId) {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var event m.Event

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), 500)
	}

	if err := r.Body.Close(); err != nil {
        panic(err)
	}

	if err := json.Unmarshal(body, &event); err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusBadRequest) // unprocessable entity
		if err := json.NewEncoder(w).Encode(err); err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
	}

	e := m.UpdateEvent(eventId, event, connection())
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	w.Header().Set("Location", r.URL.String())
    w.WriteHeader(http.StatusOK)
    if err := json.NewEncoder(w).Encode(e); err != nil {
        http.Error(w, err.Error(), 500)
		return
	}
}



// ParticipantsHandler handler participants information
func ParticipantsHandler(w http.ResponseWriter, r *http.Request) {
	
	vars := mux.Vars(r)
	eventId := vars["eventId"]

	w.Header().Set("Content-Type", "application/json;charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(m.GetParticipantsList(eventId, connection())); err != nil {
		panic(err)
	}
}
// ParticipantHandler handler participant information
func ParticipantHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	eventId := vars["eventId"]
	participantId := vars["participantId"]

	w.Header().Set("Content-Type", "application/json;charset=UTF-8")

	if !IsValidUUID(eventId) {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	participantNo, _ := strconv.Atoi(participantId)
	if !IsValidRaceNumber(participantNo){
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	participant := m.GetParticipant(eventId, participantNo, connection())
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(participant); err != nil {
        http.Error(w, err.Error(), 500)
		return
	}
}
// ParticipantCreateHandler handler create participant information
func ParticipantCreateHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	eventId := vars["eventId"]
	var participant m.Participant
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), 500)
	}

	if err := json.Unmarshal(body, &participant); err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusBadRequest) // unprocessable entity
		if err := json.NewEncoder(w).Encode(err); err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
	}
	
	p := m.CreateParticipant(eventId, participant, connection())
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	if strconv.Itoa(p.No) == ""{
		w.WriteHeader(http.StatusConflict)
	} else {
		w.WriteHeader(http.StatusOK)	
		updateErr := m.InsertEventStages(eventId, p.No, connection())
		if updateErr != nil {
			fmt.Println("Something went wrong.")
		}
		if err := json.NewEncoder(w).Encode(p); err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
	}
	
	fmt.Println(r.URL.String())
	if err := r.Body.Close(); err != nil {
		panic(err)
	}
}
// ParticipantDeleteHandler handler delete participant information
func ParticipantDeleteHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	eventId := vars["eventId"]
	participantId := vars["participantId"]
	w.Header().Set("Content-Type", "application/json;charset=UTF-8")

	if !IsValidUUID(eventId) {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	participantNo, _ := strconv.Atoi(participantId)
	if !IsValidStageId(participantNo){
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err := m.DeleteParticipant(eventId, participantNo, connection())
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
	}
	err = m.DeleteEventParticipant(eventId, participantNo, connection())
	if err != nil{
		fmt.Println(fmt.Errorf("%s", err))
	}
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(m.GetParticipantsList(eventId, connection())); err != nil {
        http.Error(w, err.Error(), 500)
		return
	}
}
// ParticipantUpdateHandler handler update participant information
func ParticipantUpdateHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	eventId := vars["eventId"]
	participantId := vars["participantId"]
	w.Header().Set("Content-Type", "application/json;charset=UTF-8")

	if !IsValidUUID(eventId) {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	participantNo, _ := strconv.Atoi(participantId)
	if !IsValidStageId(participantNo){
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var particicipant m.Participant
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), 500)
	}
	if err := r.Body.Close(); err != nil {
        panic(err)
	}

	if err := json.Unmarshal(body, &particicipant); err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusBadRequest) // unprocessable entity
		if err := json.NewEncoder(w).Encode(err); err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
	}

	p := m.UpdateParticipant(eventId, participantNo, particicipant, connection())
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	w.Header().Set("Location", r.URL.String())
    w.WriteHeader(http.StatusOK)
    if err := json.NewEncoder(w).Encode(p); err != nil {
        http.Error(w, err.Error(), 500)
		return
	}
}

// StagesHandler handler stages information
func StagesHandler(w http.ResponseWriter, r *http.Request) {
	
	vars := mux.Vars(r)
	eventId := vars["eventId"]

	w.Header().Set("Content-Type", "application/json;charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(m.GetStageList(eventId, connection())); err != nil {
		panic(err)
	}
}
// StageHandler handler stage information 
func StageHandler(w http.ResponseWriter, r *http.Request) {
	
	vars := mux.Vars(r)
	eventId := vars["eventId"]
	stageId := vars["stageId"]

	w.Header().Set("Content-Type", "application/json;charset=UTF-8")

	if !IsValidUUID(eventId) {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	intStageId, _ := strconv.Atoi(stageId)
	if !IsValidStageId(intStageId){
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	stage := m.GetStage(eventId, intStageId, connection())
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(stage); err != nil {
        http.Error(w, err.Error(), 500)
		return
	}
}
// StageCreateHandler handler create stage information 
func StageCreateHandler(w http.ResponseWriter, r *http.Request) {
	
	vars := mux.Vars(r)
	eventId := vars["eventId"]
	var stage m.Stage
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), 500)
	}

	if err := json.Unmarshal(body, &stage); err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusBadRequest) // unprocessable entity
		if err := json.NewEncoder(w).Encode(err); err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
	}
	
	s := m.CreateStage(eventId, stage, connection())
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	if strconv.Itoa(s.ID) == ""{
		w.WriteHeader(http.StatusConflict)
	} else {
		w.WriteHeader(http.StatusOK)	
		updateErr := m.InsertEventStages(eventId, s.ID, connection())
		if updateErr != nil {
			fmt.Println("Something went wrong.")
		}
		if err := json.NewEncoder(w).Encode(s); err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
	}
	
	fmt.Println(r.URL.String())
	if err := r.Body.Close(); err != nil {
		panic(err)
	}
}
// StageDeleteHandler handler delete stage
func StageDeleteHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	eventId := vars["eventId"]
	stageId := vars["stageId"]
	w.Header().Set("Content-Type", "application/json;charset=UTF-8")

	if !IsValidUUID(eventId) {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	intStageId, _ := strconv.Atoi(stageId)
	if !IsValidStageId(intStageId){
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err := m.DeleteStage(eventId, intStageId, connection())
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
	}
	err = m.DeleteEventStage(eventId, intStageId, connection())
	if err != nil{
		fmt.Println(fmt.Errorf("%s", err))
	}
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(m.GetStageList(eventId, connection())); err != nil {
        http.Error(w, err.Error(), 500)
		return
	}
}
// StageUpdateHandler handler update stage information 
func StageUpdateHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	eventId := vars["eventId"]
	stageId := vars["stageId"]
	w.Header().Set("Content-Type", "application/json;charset=UTF-8")

	if !IsValidUUID(eventId) {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	intStageId, _ := strconv.Atoi(stageId)
	if !IsValidStageId(intStageId){
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var stage m.Stage
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), 500)
	}
	if err := r.Body.Close(); err != nil {
        panic(err)
	}

	if err := json.Unmarshal(body, &stage); err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusBadRequest) // unprocessable entity
		if err := json.NewEncoder(w).Encode(err); err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
	}

	s := m.UpdateStage(eventId, intStageId, stage, connection())
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	w.Header().Set("Location", r.URL.String())
    w.WriteHeader(http.StatusOK)
    if err := json.NewEncoder(w).Encode(s); err != nil {
        http.Error(w, err.Error(), 500)
		return
	}
}

// SprintsHandler handler sprints information
func SprintsHandler(w http.ResponseWriter, r *http.Request) {
}
// SprintHandler handler sprint information
func SprintHandler(w http.ResponseWriter, r *http.Request) {
}
// SprintCreateHandler handler create sprint
func SprintCreateHandler(w http.ResponseWriter, r *http.Request) {
}
// SprintDeleteHandler handler delete sprint 
func SprintDeleteHandler(w http.ResponseWriter, r *http.Request) {
}
// SprintUpdateHandler handler update sprint information
func SprintUpdateHandler(w http.ResponseWriter, r *http.Request) {
}


// EventCommissairesHandler handler event commissaire information
func EventCommissairesHandler(w http.ResponseWriter, r *http.Request) {
}
// EventCommissaireHandler handler event commissaire information
func EventCommissaireHandler(w http.ResponseWriter, r *http.Request) {
}
// EventCommissaireCreateHandler handler create event commissaire information
func EventCommissaireCreateHandler(w http.ResponseWriter, r *http.Request) {
}
// EventCommissaireDeleteHandler handler delete event commissaire 
func EventCommissaireDeleteHandler(w http.ResponseWriter, r *http.Request) {
}
// EventCommissaireUpdateHandler handler update event commissaire information
func EventCommissaireUpdateHandler(w http.ResponseWriter, r *http.Request) {
}
