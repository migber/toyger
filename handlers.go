
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
	validToken := checkJwt(w,r)
	if (validToken){

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
 	} else {
		w.WriteHeader(http.StatusUnauthorized)		
	 }

}
// TeamsHandler return list of teams
func TeamsHandler(w http.ResponseWriter, r *http.Request) {
	validToken := checkJwt(w,r)
	w.Header().Set("Content-Type", "application/json;charset=UTF-8")	
	if (validToken){
		w.WriteHeader(http.StatusOK)
		if err := json.NewEncoder(w).Encode(Models.GetTeamsList(connection())); err != nil {
			panic(err)
		}
	} else {
		w.WriteHeader(http.StatusUnauthorized)		
	}
}

// TeamCreateHandler create handler
func TeamCreateHandler(w http.ResponseWriter, r *http.Request){
	validToken := checkJwt(w,r)	
		if (validToken){
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
	} else {
		w.WriteHeader(http.StatusUnauthorized)
	}
}

// TeamDeleteHandler team delete handler
func TeamDeleteHandler(w http.ResponseWriter, r *http.Request){
	validToken := checkJwt(w,r)
	if (validToken){
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
	} else {
		w.WriteHeader(http.StatusUnauthorized)
	}
}
// TeamUpdateHandler udpate team handler
func TeamUpdateHandler(w http.ResponseWriter, r *http.Request){
	validToken := checkJwt(w,r)
	if (validToken){
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
	} else {
		w.WriteHeader(http.StatusUnauthorized)
	}
}

// CyclistsHandler handler cyclists return cyclists
func CyclistsHandler(w http.ResponseWriter, r *http.Request) {
	validToken := checkJwt(w,r)
	if (validToken){
		vars := mux.Vars(r)
		teamID := vars["teamId"]
	
		w.Header().Set("Content-Type", "application/json;charset=UTF-8")
		w.WriteHeader(http.StatusOK)
		if err := json.NewEncoder(w).Encode(Models.GetCyclistsList(teamID, connection())); err != nil {
			panic(err)
		}
	} else {
		w.WriteHeader(http.StatusUnauthorized)
	}
}

//CyclistHandler handler cyclist return on cyclist
func CyclistHandler(w http.ResponseWriter, r *http.Request) {
	validToken := checkJwt(w,r)
	if (validToken){
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
	} else {
		w.WriteHeader(http.StatusUnauthorized)
	}
}

//CyclistCreateHandler handler create cyclist
func CyclistCreateHandler(w http.ResponseWriter, r *http.Request) {
	validToken := checkJwt(w,r)
	if (validToken){
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
		if err := r.Body.Close(); err != nil {
			panic(err)
		}
	} else {
		w.WriteHeader(http.StatusUnauthorized)
	}
}

//CyclistUpdateHandler handler update cyclist
func CyclistUpdateHandler(w http.ResponseWriter, r *http.Request) {
	validToken := checkJwt(w,r)
	if (validToken){
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
	} else {
		w.WriteHeader(http.StatusUnauthorized)
	}
}

//CyclistDeleteHandler handler delete cyclist
func CyclistDeleteHandler(w http.ResponseWriter, r *http.Request) {
	validToken := checkJwt(w,r)
	if (validToken){
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
	} else {
		w.WriteHeader(http.StatusUnauthorized)
	}
}

// CommissairesHandler handler commissaires
func CommissairesHandler(w http.ResponseWriter, r *http.Request) {
	validToken := checkJwt(w,r)
	if (validToken){
		w.Header().Set("Content-Type", "application/json;charset=UTF-8")
		w.WriteHeader(http.StatusOK)
		if err := json.NewEncoder(w).Encode(Models.GetCommissairesList(connection())); err != nil {
			panic(err)
		}
	} else {
		w.WriteHeader(http.StatusUnauthorized)
	}
}

// CommissaireHandler handler commissaire information
func CommissaireHandler(w http.ResponseWriter, r *http.Request) {
	validToken := checkJwt(w,r)
	if (validToken){
		vars := mux.Vars(r)
		comId := vars["commissaireId"]
	
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
	} else {
		w.WriteHeader(http.StatusUnauthorized)
	}
}

// CommissaireCreateHandler handler create commissaire
func CommissaireCreateHandler(w http.ResponseWriter, r *http.Request) {
	validToken := checkJwt(w,r)
	if (validToken){
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
	} else {
		w.WriteHeader(http.StatusUnauthorized)
	}
}

// CommissaireUpdateHandler handler update commissaire information
func CommissaireUpdateHandler(w http.ResponseWriter, r *http.Request) {
	validToken := checkJwt(w,r)
	if (validToken){
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
	} else {
		w.WriteHeader(http.StatusUnauthorized)
	}
}

// CommissaireDeleteHandler handler delete commissaires
func CommissaireDeleteHandler(w http.ResponseWriter, r *http.Request) {
	validToken := checkJwt(w,r)
	if (validToken){
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
	} else {
		w.WriteHeader(http.StatusUnauthorized)
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
	validToken := checkJwt(w,r)
	if (validToken){
		w.Header().Set("Content-Type", "application/json;charset=UTF-8")
		w.WriteHeader(http.StatusOK)
		if err := json.NewEncoder(w).Encode(m.GetEventsList(connection())); err != nil {
			panic(err)
		}
	} else {
		w.WriteHeader(http.StatusUnauthorized)
	}
}

// EventHandler handler event information
func EventHandler(w http.ResponseWriter, r *http.Request) {
	validToken := checkJwt(w,r)
	if (validToken){
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
	} else {
		w.WriteHeader(http.StatusUnauthorized)
	}
}

// EventCreateHandler handler create event information
func EventCreateHandler(w http.ResponseWriter, r *http.Request) {
	validToken := checkJwt(w,r)
	if (validToken){
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
	} else {
		w.WriteHeader(http.StatusUnauthorized)
	}
}

// EventDeleteHandler handler delete event information
func EventDeleteHandler(w http.ResponseWriter, r *http.Request) {
	validToken := checkJwt(w,r)
	if (validToken){
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
	} else {
		w.WriteHeader(http.StatusUnauthorized)
	}
}

// EventUpdateHandler handler update event information
func EventUpdateHandler(w http.ResponseWriter, r *http.Request) {
	validToken := checkJwt(w,r)
	if (validToken){
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
	} else {
		w.WriteHeader(http.StatusUnauthorized)
	}
}

// ParticipantsHandler handler participants information
func ParticipantsHandler(w http.ResponseWriter, r *http.Request) {
	validToken := checkJwt(w,r)
	if (validToken){
		vars := mux.Vars(r)
		eventId := vars["eventId"]
	
		w.Header().Set("Content-Type", "application/json;charset=UTF-8")
		w.WriteHeader(http.StatusOK)
		if err := json.NewEncoder(w).Encode(m.GetParticipantsList(eventId, connection())); err != nil {
			panic(err)
		}
	} else {
		w.WriteHeader(http.StatusUnauthorized)
	}
}

// ParticipantHandler handler participant information
func ParticipantHandler(w http.ResponseWriter, r *http.Request) {
	validToken := checkJwt(w,r)
	if(validToken){
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
	} else {
		w.WriteHeader(http.StatusUnauthorized)
	}
}

// ParticipantCreateHandler handler create participant information
func ParticipantCreateHandler(w http.ResponseWriter, r *http.Request) {
	validToken := checkJwt(w,r)
	if (validToken) {
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
	} else {
		w.WriteHeader(http.StatusUnauthorized)
	}
}

// ParticipantDeleteHandler handler delete participant information
func ParticipantDeleteHandler(w http.ResponseWriter, r *http.Request) {
	validToken := checkJwt(w,r)
	if (validToken){
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
	} else {
		w.WriteHeader(http.StatusUnauthorized)
	}
}

// ParticipantUpdateHandler handler update participant information
func ParticipantUpdateHandler(w http.ResponseWriter, r *http.Request) {
	validToken := checkJwt(w,r)
	if (validToken){
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
	} else {
		w.WriteHeader(http.StatusUnauthorized)
	}
}

// StagesHandler handler stages information
func StagesHandler(w http.ResponseWriter, r *http.Request) {
	validToken := checkJwt(w,r)
	if (validToken) {
		vars := mux.Vars(r)
		eventId := vars["eventId"]
	
		w.Header().Set("Content-Type", "application/json;charset=UTF-8")
		w.WriteHeader(http.StatusOK)
		if err := json.NewEncoder(w).Encode(m.GetStageList(eventId, connection())); err != nil {
			panic(err)
		}
	} else {
		w.WriteHeader(http.StatusUnauthorized)
	}
}

// StageHandler handler stage information 
func StageHandler(w http.ResponseWriter, r *http.Request) {
	validToken := checkJwt(w,r)
	if (validToken) {
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
	} else {
		w.WriteHeader(http.StatusUnauthorized)
	}
}

// StageCreateHandler handler create stage information 
func StageCreateHandler(w http.ResponseWriter, r *http.Request) {
	validToken := checkJwt(w,r)
	if (validToken){
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
	} else {
		w.WriteHeader(http.StatusUnauthorized)
	}
}

// StageDeleteHandler handler delete stage
func StageDeleteHandler(w http.ResponseWriter, r *http.Request) {
	validToken := checkJwt(w,r)
	if (validToken){
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
	} else {
		w.WriteHeader(http.StatusUnauthorized)
	}
}

// StageUpdateHandler handler update stage information 
func StageUpdateHandler(w http.ResponseWriter, r *http.Request) {
	validToken := checkJwt(w,r)
	if (validToken){
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
	} else {
		w.WriteHeader(http.StatusUnauthorized)
	}
}

// SprintsHandler handler sprints information
func SprintsHandler(w http.ResponseWriter, r *http.Request) {
	validToken := checkJwt(w,r)
	if (validToken) {
		vars := mux.Vars(r)
		eventId := vars["eventId"]
		stageId := vars["stageId"]
	
		intStageId, _ := strconv.Atoi(stageId)
	
		w.Header().Set("Content-Type", "application/json;charset=UTF-8")
		w.WriteHeader(http.StatusOK)
		if err := json.NewEncoder(w).Encode(m.GetSprintList(eventId, intStageId, connection())); err != nil {
			panic(err)
		}
	} else {
		w.WriteHeader(http.StatusUnauthorized)
	}
}

// SprintHandler handler sprint information
func SprintHandler(w http.ResponseWriter, r *http.Request) {
	validToken := checkJwt(w,r)
	if (validToken) {
		vars := mux.Vars(r)
		eventId := vars["eventId"]
		stageId := vars["stageId"]
		sprintId := vars["sprintId"]
	
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
		if !IsValidUUID(sprintId) {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
	
		sprint := m.GetSprint(eventId, intStageId, sprintId, connection())
		w.WriteHeader(http.StatusOK)
	
		if err := json.NewEncoder(w).Encode(sprint); err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
	} else {
		w.WriteHeader(http.StatusUnauthorized)
	}
}

// SprintCreateHandler handler create sprint
func SprintCreateHandler(w http.ResponseWriter, r *http.Request) {
	validToken := checkJwt(w,r)
	if (validToken){
		vars := mux.Vars(r)
		eventId := vars["eventId"]
		stageId := vars["stageId"]
		intStageId, _ := strconv.Atoi(stageId)
	
		var sprint m.Sprint
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			http.Error(w, err.Error(), 500)
		}
	
		if err := json.Unmarshal(body, &sprint); err != nil {
			w.Header().Set("Content-Type", "application/json; charset=UTF-8")
			w.WriteHeader(http.StatusBadRequest) // unprocessable entity
			if err := json.NewEncoder(w).Encode(err); err != nil {
				http.Error(w, err.Error(), 500)
				return
			}
		}
		
		s := m.CreateSprint(eventId, intStageId, sprint, connection())
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		if s.Id == ""{
			w.WriteHeader(http.StatusConflict)
		} else {
			w.WriteHeader(http.StatusOK)	
			updateErr := m.AddSprints(eventId, intStageId, sprint, connection())
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
	} else {
		w.WriteHeader(http.StatusUnauthorized)
	}
}

// SprintDeleteHandler handler delete sprint 
func SprintDeleteHandler(w http.ResponseWriter, r *http.Request) {
	validToken := checkJwt(w,r)
	if (validToken){
		vars := mux.Vars(r)
		eventId := vars["eventId"]
		stageId := vars["stageId"]
		sprintId := vars["sprintId"]
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
	
		err := m.DeleteSprint(eventId, intStageId, sprintId, connection())
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
		}
		err = m.DeleteStageSprint(eventId, intStageId, sprintId, connection())
		if err != nil{
			fmt.Println(fmt.Errorf("%s", err))
		}
		w.WriteHeader(http.StatusOK)
	
		if err := json.NewEncoder(w).Encode(m.GetSprintList(eventId, intStageId, connection())); err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
	} else {
		w.WriteHeader(http.StatusUnauthorized)
	}
}

// SprintUpdateHandler handler update sprint information
func SprintUpdateHandler(w http.ResponseWriter, r *http.Request) {
	validToken := checkJwt(w,r)
	if (validToken){
		vars := mux.Vars(r)
		eventId := vars["eventId"]
		stageId := vars["stageId"]
		sprintId := vars["sprintId"]
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
		if !IsValidUUID(eventId) {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
	
		var sprint m.Sprint
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			http.Error(w, err.Error(), 500)
		}
		if err := r.Body.Close(); err != nil {
			panic(err)
		}
	
		if err := json.Unmarshal(body, &sprint); err != nil {
			w.Header().Set("Content-Type", "application/json; charset=UTF-8")
			w.WriteHeader(http.StatusBadRequest) // unprocessable entity
			if err := json.NewEncoder(w).Encode(err); err != nil {
				http.Error(w, err.Error(), 500)
				return
			}
		}
	
		s := m.UpdateSprint(eventId, intStageId, sprintId, sprint, connection())
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		err = m.UpdateStageSprint(intStageId, sprintId, sprint, connection())
		if err != nil{
			fmt.Println(fmt.Errorf("%s", err))
		}
		w.Header().Set("Location", r.URL.String())
		w.WriteHeader(http.StatusOK)
		if err := json.NewEncoder(w).Encode(s); err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
	} else {
		w.WriteHeader(http.StatusUnauthorized)
	}
}

// EventCommissairesHandler handler event commissaire information
func EventCommissairesHandler(w http.ResponseWriter, r *http.Request) {
	validToken := checkJwt(w,r)
	if (validToken){
		vars := mux.Vars(r)
		eventId := vars["eventId"]
	
		w.Header().Set("Content-Type", "application/json;charset=UTF-8")
		w.WriteHeader(http.StatusOK)
		if err := json.NewEncoder(w).Encode(m.GetRaceCommissairesList(eventId, connection())); err != nil {
			panic(err)
		}	
	} else {
		w.WriteHeader(http.StatusUnauthorized)
	}
}

// EventCommissaireHandler handler event commissaire information
func EventCommissaireHandler(w http.ResponseWriter, r *http.Request) {
	validToken := checkJwt(w,r)
	if (validToken){
		vars := mux.Vars(r)
		eventId := vars["eventId"]
		raceCommId := vars["commissaireId"]
	
		w.Header().Set("Content-Type", "application/json;charset=UTF-8")
	
		if !IsValidUUID(eventId) {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
	
		if !isValidUCIID(raceCommId) {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
	
		raceCommissaire := m.GetRaceCommissaire(eventId, raceCommId, connection())
		w.WriteHeader(http.StatusOK)
	
		if err := json.NewEncoder(w).Encode(raceCommissaire); err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
	} else {
		w.WriteHeader(http.StatusUnauthorized)
	}
}

// EventCommissaireCreateHandler handler create event commissaire information
func EventCommissaireCreateHandler(w http.ResponseWriter, r *http.Request) {
	validToken := checkJwt(w,r)
	if (validToken){
		vars := mux.Vars(r)
		eventId := vars["eventId"]
		var raceCommissaire m.RaceCommissaire
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			http.Error(w, err.Error(), 500)
		}
	
		if err := json.Unmarshal(body, &raceCommissaire); err != nil {
			w.Header().Set("Content-Type", "application/json; charset=UTF-8")
			w.WriteHeader(http.StatusBadRequest) // unprocessable entity
			if err := json.NewEncoder(w).Encode(err); err != nil {
				http.Error(w, err.Error(), 500)
				return
			}
		}
		
		s := m.CreateRaceCommissaire(eventId, raceCommissaire, connection())
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		if s.Commissaire.UCIID == ""{
			w.WriteHeader(http.StatusConflict)
		} else {
			w.WriteHeader(http.StatusOK)	
			updateErr := m.InsertEventCommissaire(eventId, s.Commissaire.UCIID, connection())
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
	} else {
		w.WriteHeader(http.StatusUnauthorized)
	}
}

// EventCommissaireDeleteHandler handler delete event commissaire 
func EventCommissaireDeleteHandler(w http.ResponseWriter, r *http.Request) {
	validToken := checkJwt(w,r)
	if (validToken) {
		vars := mux.Vars(r)
		eventId := vars["eventId"]
		rcId := vars["commissaireId"]
		w.Header().Set("Content-Type", "application/json;charset=UTF-8")
	
		if !IsValidUUID(eventId) {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		if !isValidUCIID(rcId){
			w.WriteHeader(http.StatusBadRequest)
			return
		}
	
		err := m.DeleteRaceCommissaire(eventId, rcId, connection())
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
		}
		err = m.DeleteEventCommissaire(eventId, rcId, connection())
		if err != nil{
			fmt.Println(fmt.Errorf("%s", err))
		}
		w.WriteHeader(http.StatusOK)
	
		if err := json.NewEncoder(w).Encode(m.GetRaceCommissairesList(eventId, connection())); err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
	} else {
		w.WriteHeader(http.StatusUnauthorized)
	}
}

// EventCommissaireUpdateHandler handler update event commissaire information
func EventCommissaireUpdateHandler(w http.ResponseWriter, r *http.Request) {
	validToken := checkJwt(w,r)
	if (validToken) {
		vars := mux.Vars(r)
		eventId := vars["eventId"]
		commissaireId := vars["commissaireId"]
		w.Header().Set("Content-Type", "application/json;charset=UTF-8")
	
		if !IsValidUUID(eventId) {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		if !isValidUCIID(commissaireId){
			w.WriteHeader(http.StatusBadRequest)
			return
		}
	
		var raceCommiss m.RaceCommissaire
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			http.Error(w, err.Error(), 500)
		}
		if err := r.Body.Close(); err != nil {
			panic(err)
		}
	
		if err := json.Unmarshal(body, &raceCommiss); err != nil {
			w.Header().Set("Content-Type", "application/json; charset=UTF-8")
			w.WriteHeader(http.StatusBadRequest) // unprocessable entity
			if err := json.NewEncoder(w).Encode(err); err != nil {
				http.Error(w, err.Error(), 500)
				return
			}
		}
	
		rc := m.UpdateRaceCommissaire(eventId, commissaireId, raceCommiss, connection())
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	
		w.Header().Set("Location", r.URL.String())
		w.WriteHeader(http.StatusOK)
		if err := json.NewEncoder(w).Encode(rc); err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
	} else {
		w.WriteHeader(http.StatusUnauthorized)
	}
}
