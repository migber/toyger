package main

import (
    "fmt"
    "html"
    "log"
	"net/http"
	"encoding/json"
	"models"
	
	"github.com/gorilla/mux"
	"github.com/satori/go.uuid"
)

func main() {

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", MainHanlder)
	router.HandleFunc("/healthcheck", HealtchCheck)

	// independent 
	router.HandleFunc("/teams", teamsHandler)
	router.HandleFunc("/teams/{teamId}", teamHandler)
	router.HandleFunc("/cyclists", cyclistsHandler)
	router.HandleFunc("/cyclists/{cyclistId}", cyclistHandler)
	router.HandleFunc("/commissaires", commissairesHandler)
	router.HandleFunc("/commissaires/{comId}", commissaireHandler)
	router.HandleFunc("/managers", managersHandler)
	router.HandleFunc("/managers/{managerId}", managerHandler)

	// event dependence
	router.HandleFunc("/events", eventsHandler)
	router.HandleFunc("/events/{eventId}", eventHandler)
	router.HandleFunc("/event/{eventId}/participants", participantsHandler)
	router.HandleFunc("/event/{eventId}/participants/{participantId}", participantHandler)
	router.HandleFunc("/events/{eventId}/stages", stagesHandler)
	router.HandleFunc("/events/{eventId}/stages/{stageId}", stageHandler)
	router.HandleFunc("/events/{eventId}/stages/{stageId}/sprints", sprintsHandler)
	router.HandleFunc("/events/{eventId}/stages/{stageId}/sprints/{sprintId}", sprintHandler)
	router.HandleFunc("/events/{eventId}/racecommissaires", eventCommissairesHandler)
	router.HandleFunc("/events/{eventId}/racecommissaires/{commissaireId}", eventCommissaireHandler)
	
	
    log.Fatal(http.ListenAndServe(":8080", router))
}

func HealtchCheck(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hey I am working here %s!", html.EscapeString(r.URL.Path))
}

func MainHanlder(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
}

func teamHandler (w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	teamId := vars["teamId"]
	fmt.Fprintln(w, "Team show:", teamId)
}

func teamsHandler(w http.ResponseWriter, r *http.Request) {

	teams := Teams{
		Team {Id: uuid.NewV4(),
			  Name: "Team 1",
			  Manager {
				  Id: uuid.NewV4(),
				  Name:	"Manager1",
				  Surname:	"ManagerSurname",
				  Birthdate:	"2017-05-06"
				  Phone:		"+37063015590"
				  Nationality:	"Lithuanian"
			  }},
		Team {Id: uuid.NewV4(),
		Name: "Team 2",
		Manager {
			Id: uuid.NewV4(),
			Name:	"Manager2",
			Surname:	"ManagerSurname2",
			Birthdate:	"2017-05-08"
			Phone:		"+37063015590"
			Nationality:	"Lithuanian"
		}},
	}
	
	json.NewEncoder(w).Encode(teams)
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

