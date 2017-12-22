package main

import (
	"net/http"
)
// Route struct of describing route
type Route struct{
	Name 		string
	Method 		string
	Pattern		string
	HandlerFunc	http.HandlerFunc
}

// Routes have all Rautes
type Routes []Route

var routes = Routes {
	Route{
		"Main",
		"GET",
		"/",
		Hanlder,
	},
	Route{
		"Main",
		"GET",
		"/health",
		HealthCheck,
	},

	// Teams 
	Route{
		"Teams",
		"GET",
		"/api/teams",
		TeamsHandler,
	},
	Route{
		"Team",
		"GET",
		"/api/teams/{teamId}",
		TeamHandler,
	},
	Route{
		"TeamCreate",
		"POST",
		"/api/teams",
		TeamCreateHandler,
	},
	Route{
		"DeleteTeam",
		"GET",
		"/api/teams/{teamId}/delete",
		TeamDeleteHandler,
	},
	Route{
		"DeleteTeam",
		"DELETE",
		"/api/teams/{teamId}",
		TeamDeleteHandler,
	},
	Route{
		"UpdateTeam",
		"POST",
		"/api/teams/{teamId}/update",
		TeamUpdateHandler,
	},
	Route{
		"UpdateTeam",
		"PUT",
		"/api/teams/{teamId}",
		TeamUpdateHandler,
	},

	// Cyclist alone
	Route{
		"Cyclistss",
		"GET",
		"/api/cyclists",
		CyclistsAloneHandler,
	},
	Route{
		"CyclistAlone",
		"GET",
		"/api/cyclists/{cyclistId}",
		CyclistAloneHandler,
	},
	Route{
		"CyclistAloneCreate",
		"POST",
		"/api/cyclists",
		CyclistsAloneCreateHandler,
	},
	Route{
		"CyclistAloneDelete",
		"GET",
		"/api/cyclists/{cyclistId}/delete",
		CyclistsAloneDeleteHandler,
	},
	Route{
		"CyclistAloneUpdate",
		"POST",
		"/api/cyclists/{cyclistId}/update",
		CyclistAloneUpdateHandler,
	},

	// Cyclists 
	Route{
		"Cyclists",
		"GET",
		"/api/teams/{teamId}/cyclists",
		CyclistsHandler,
	},
	Route{
		"Cyclist",
		"GET",
		"/api/teams/{teamId}/cyclists/{cyclistId}",
		CyclistHandler,
	},
	Route{
		"CyclistCreate",
		"POST",
		"/api/teams/{teamId}/cyclists",
		CyclistCreateHandler,
	},
	Route{
		"CyclistDelete",
		"DELETE",
		"/api/teams/{teamId}/cyclists/{cyclistId}",
		CyclistDeleteHandler,
	},
	Route{
		"CyclistUpdate",
		"PUT",
		"/api/teams/{teamId}/cyclists/{cyclistId}",
		CyclistUpdateHandler,
	},

	// Commissaires 
	Route{
		"Commissaires",
		"GET",
		"/api/commissaires",
		CommissairesHandler,
	},
	Route{
		"CommissaireGet",
		"GET",
		"/api/commissaires/{commissaireID}",
		CommissaireHandler,
	},
	Route{
		"CommissaireCreate",
		"POST",
		"/api/commissaires",
		CommissaireCreateHandler,
	},
	Route{
		"CommissaireDelete",
		"GET",
		"/api/commissaires/{commissaireID}/delete",
		CommissaireDeleteHandler,
	},
	Route{
		"CommissaireUpdate",
		"POST",
		"/api/commissaires/{commissaireID}/update",
		CommissaireUpdateHandler,
	},

	// Manager 
	Route{
		"Managers",
		"GET",
		"/api/managers",
		ManagersHandler,
	},
	Route{
		"Managers Get",
		"GET",
		"/api/managers/{managerId}",
		ManagerHandler,
	},
	Route{
		"ManagerCreate",
		"POST",
		"/api/managers",
		ManagerCreateHandler,
	},
	Route{
		"ManagerDelete",
		"DELETE",
		"/api/managers/{managerId}",
		ManagerDeleteHandler,
	},
	Route{
		"ManagerUpdate",
		"PUT",
		"/api/managers/{managerId}",
		ManagerUpdateHandler,
	},

	// Event 
	Route{
		"Events",
		"GET",
		"/api/events",
		EventsHandler,
	},
	Route{
		"Event Get",
		"GET",
		"/api/events/{eventID}",
		EventHandler,
	},
	Route{
		"EventCreate",
		"POST",
		"/api/events",
		EventCreateHandler,
	},
	Route{
		"EventDelete",
		"GET",
		"/api/events/{eventID}/delete",
		EventDeleteHandler,
	},
	Route{
		"EventUpdate",
		"POST",
		"/api/events/{eventID}/update",
		EventUpdateHandler,
	},

	// Participants  
	Route{
		"Participants",
		"GET",
		"/api/events/{eventID}/participants",
		ParticipantsHandler,
	},
	Route{
		"Participant",
		"GET",
		"/api/events/{eventID}/participants/{participantId}",
		ParticipantHandler,
	},
	Route{
		"ParticipantCreate",
		"POST",
		"/api/events/{eventID}/participants",
		ParticipantCreateHandler,
	},
	Route{
		"ParticipantDelete",
		"DELETE",
		"/api/events/{eventID}/participants/{participantId}",
		ParticipantDeleteHandler,
	},
	Route{
		"ParticipantUpdate",
		"PUT",
		"/api/events/{eventID}/participants/{participantId}",
		ParticipantUpdateHandler,
	},

	// Stages
	Route{
		"Stages",
		"GET",
		"/api/events/{eventID}/stages",
		StagesHandler,
	},
	Route{
		"Stage",
		"GET",
		"/api/events/{eventID}/stages/{stageID}",
		StageHandler,
	},
	Route{
		"StageCreate",
		"POST",
		"/api/events/{eventID}/stages",
		StageCreateHandler,
	},
	Route{
		"StageDelete",
		"DELETE",
		"/api/events/{eventID}/stages/{stagesId}",
		StageDeleteHandler,
	},
	Route{
		"StageUpdate",
		"PUT",
		"/api/events/{eventID}/stages/{stageID}",
		StageUpdateHandler,
	},

	// Sprints 
	Route{
		"Sprints",
		"GET",
		"/api/events/{eventID}/stages/{stageID}/sprints",
		SprintsHandler,
	},
	Route{
		"Stage",
		"GET",
		"/api/events/{eventID}/stages/{stageID}/sprints/{sprintID}",
		SprintHandler,
	},
	Route{
		"SprintCreate",
		"POST",
		"/api/events/{eventID}/stages/{stageID}/sprints",
		SprintCreateHandler,
	},
	Route{
		"SprintDelete",
		"DELETE",
		"/api/events/{eventID}/stages/{stageID}/sprints/{sprintID}}",
		SprintDeleteHandler,
	},
	Route{
		"SprintUpdate",
		"PUT",
		"/api/events/{eventID}/stages/{stageID}/sprints/{sprintID}",
		SprintUpdateHandler,
	},

	// Event commissaires
	Route{
		"Event commissaires",
		"GET",
		"/api/events/{eventID}/commissaires",
		EventCommissairesHandler,
	},
	Route{
		"Event commissaire",
		"GET",
		"/api/events/{eventID}/commissaires/{commissaireID}",
		EventCommissaireHandler,
	},
	Route{
		"Event commissaire create",
		"POST",
		"/api/events/{eventID}/commissaires",
		EventCommissaireCreateHandler,
	},
	Route{
		"Event commissaire delete",
		"DELETE",
		"/api/events/{eventID}/commissaires/{commissaireID}",
		EventCommissaireDeleteHandler,
	},
	Route{
		"Event commissaire update",
		"PUT",
		"/api/events/{eventID}/commissaires/{commissaireID}",
		EventCommissaireUpdateHandler,
	},
}
