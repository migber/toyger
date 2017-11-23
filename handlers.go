
package main

import (
    "encoding/json"
    "fmt"
	"net/http"
	"io/ioutil"
	"html"
	Models "toyger/models"
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
func TeamHandler (w http.ResponseWriter, r *http.Request) {

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

// CyclistsHandler handler cyclists
func CyclistsHandler(w http.ResponseWriter, r *http.Request) {
}

//CyclistHandler handler cyclist
func CyclistHandler(w http.ResponseWriter, r *http.Request) {
}

//CyclistCreateHandler handler create cyclist
func CyclistCreateHandler(w http.ResponseWriter, r *http.Request) {
}

//CyclistUpdateHandler handler update cyclist
func CyclistUpdateHandler(w http.ResponseWriter, r *http.Request) {
}

//CyclistDeleteHandler handler delete cyclist
func CyclistDeleteHandler(w http.ResponseWriter, r *http.Request) {
}



// CommissairesHandler handler commissaires
func CommissairesHandler(w http.ResponseWriter, r *http.Request) {
}

// CommissaireHandler handler commissaire information
func CommissaireHandler(w http.ResponseWriter, r *http.Request) {
}

// CommissaireCreateHandler handler create commissaire
func CommissaireCreateHandler(w http.ResponseWriter, r *http.Request) {
}

// CommissaireUpdateHandler handler update commissaire information
func CommissaireUpdateHandler(w http.ResponseWriter, r *http.Request) {
}
// CommissaireDeleteHandler handler delete commissaires
func CommissaireDeleteHandler(w http.ResponseWriter, r *http.Request) {
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
}

// EventHandler handler event information
func EventHandler(w http.ResponseWriter, r *http.Request) {
}

// EventCreateHandler handler create event information
func EventCreateHandler(w http.ResponseWriter, r *http.Request) {
}

// EventDeleteHandler handler delete event information
func EventDeleteHandler(w http.ResponseWriter, r *http.Request) {
}

// EventUpdateHandler handler update event information
func EventUpdateHandler(w http.ResponseWriter, r *http.Request) {
}



// ParticipantsHandler handler participants information
func ParticipantsHandler(w http.ResponseWriter, r *http.Request) {
}
// ParticipantHandler handler participant information
func ParticipantHandler(w http.ResponseWriter, r *http.Request) {
}
// ParticipantCreateHandler handler create participant information
func ParticipantCreateHandler(w http.ResponseWriter, r *http.Request) {
}
// ParticipantDeleteHandler handler delete participant information
func ParticipantDeleteHandler(w http.ResponseWriter, r *http.Request) {
}
// ParticipantUpdateHandler handler update participant information
func ParticipantUpdateHandler(w http.ResponseWriter, r *http.Request) {
}

// StagesHandler handler stages information
func StagesHandler(w http.ResponseWriter, r *http.Request) {
}
// StageHandler handler stage information 
func StageHandler(w http.ResponseWriter, r *http.Request) {
}
// StageCreateHandler handler create stage information 
func StageCreateHandler(w http.ResponseWriter, r *http.Request) {
}
// StageDeleteHandler handler delete stage
func StageDeleteHandler(w http.ResponseWriter, r *http.Request) {
}
// StageUpdateHandler handler update stage information 
func StageUpdateHandler(w http.ResponseWriter, r *http.Request) {
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
