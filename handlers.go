
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

func HealtchCheck(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hey I am working here %s!", html.EscapeString(r.URL.Path))
}

func MainHanlder(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
}

func teamHandler (w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	teamId := vars["teamId"]

	w.Header().Set("Content-Type", "application/json;charset=UTF-8")

	if !IsValidUUID(teamId) {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	team := Models.GetTeam(teamId)
	if (Models.Team{}) == team {
		w.WriteHeader(http.StatusNotFound )
	}
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(team); err != nil {
        http.Error(w, err.Error(), 500)
		return
	}
}

func teamsHandler(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json;charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(Models.GetTeamsList()); err != nil {
		panic(err)
	}
}

func teamCreate(w http.ResponseWriter, r *http.Request){
	
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
	
	t := Models.CreateTeam(team)
    w.Header().Set("Content-Type", "application/json; charset=UTF-8")
    w.WriteHeader(http.StatusCreated)
    if err := json.NewEncoder(w).Encode(t); err != nil {
        http.Error(w, err.Error(), 500)
		return
	}
	
	if err := r.Body.Close(); err != nil {
        panic(err)
	}
}


func cyclistsHandler(w http.ResponseWriter, r *http.Request) {
}

func cyclistHandler(w http.ResponseWriter, r *http.Request) {
}

func commissairesHandler(w http.ResponseWriter, r *http.Request) {
}

func commissaireHandler(w http.ResponseWriter, r *http.Request) {
}

func managersHandler(w http.ResponseWriter, r *http.Request) {
	
}

func managerHandler(w http.ResponseWriter, r *http.Request) {
	
}

// event dependence

func eventsHandler(w http.ResponseWriter, r *http.Request) {
}

func eventHandler(w http.ResponseWriter, r *http.Request) {
}

func participantHandler(w http.ResponseWriter, r *http.Request) {
}

func participantsHandler(w http.ResponseWriter, r *http.Request) {
}

func stagesHandler(w http.ResponseWriter, r *http.Request) {
}

func stageHandler(w http.ResponseWriter, r *http.Request) {
}

func sprintsHandler(w http.ResponseWriter, r *http.Request) {
}

func sprintHandler(w http.ResponseWriter, r *http.Request) {
}

func eventCommissairesHandler(w http.ResponseWriter, r *http.Request) {
}

func eventCommissaireHandler(w http.ResponseWriter, r *http.Request) {
}
