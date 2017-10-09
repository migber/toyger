package main

import (
    "log"
	"net/http"
)

func main() {

	router := NewRouter()
	log.Fatal(http.ListenAndServe(":8080", router))
	// router := mux.NewRouter().StrictSlash(true)
	// router.HandleFunc("/", MainHanlder)
	// router.HandleFunc("/healthcheck", HealtchCheck)

	// // independent 
	// router.HandleFunc("/teams", teamsHandler)
	// router.HandleFunc("/teams/{teamId}", teamHandler)
	// router.HandleFunc("/cyclists", cyclistsHandler)
	// router.HandleFunc("/cyclists/{cyclistId}", cyclistHandler)
	// router.HandleFunc("/commissaires", commissairesHandler)
	// router.HandleFunc("/commissaires/{comId}", commissaireHandler)
	// router.HandleFunc("/managers", managersHandler)
	// router.HandleFunc("/managers/{managerId}", managerHandler)

	// // event dependence
	// router.HandleFunc("/events", eventsHandler)
	// router.HandleFunc("/events/{eventId}", eventHandler)
	// router.HandleFunc("/event/{eventId}/participants", participantsHandler)
	// router.HandleFunc("/event/{eventId}/participants/{participantId}", participantHandler)
	// router.HandleFunc("/events/{eventId}/stages", stagesHandler)
	// router.HandleFunc("/events/{eventId}/stages/{stageId}", stageHandler)
	// router.HandleFunc("/events/{eventId}/stages/{stageId}/sprints", sprintsHandler)
	// router.HandleFunc("/events/{eventId}/stages/{stageId}/sprints/{sprintId}", sprintHandler)
	// router.HandleFunc("/events/{eventId}/racecommissaires", eventCommissairesHandler)
	// router.HandleFunc("/events/{eventId}/racecommissaires/{commissaireId}", eventCommissaireHandler)
	
	
    // log.Fatal(http.ListenAndServe(":8080", router))
}

