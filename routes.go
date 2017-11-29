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
		"DELETE",
		"/api/teams/{teamId}",
		TeamDeleteHandler,
	},
	Route{
		"UpdateTeam",
		"PUT",
		"/api/teams/{teamId}",
		TeamUpdateHandler,
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
		"/api/commissaires/{commissaireId}",
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
		"DELETE",
		"/api/commissaires/{commissaireId}",
		CommissaireUpdateHandler,
	},
	Route{
		"CommissaireUpdate",
		"PUT",
		"/api/commissaires/{commissaireId}",
		CommissaireDeleteHandler,
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
		"Evnet Get",
		"GET",
		"/api/events/{eventId}",
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
		"DELETE",
		"/api/events/{eventId}",
		EventDeleteHandler,
	},
	Route{
		"EventUpdate",
		"PUT",
		"/api/events/{eventId}",
		EventUpdateHandler,
	},

	// Participants  
	Route{
		"Participants",
		"GET",
		"/api/event/participants",
		ParticipantsHandler,
	},
	Route{
		"Participant",
		"GET",
		"/api/event/{eventId}/participants/{participantId}",
		ParticipantHandler,
	},
	Route{
		"ParticipantCreate",
		"POST",
		"/api/event/{eventId}/participants",
		ParticipantCreateHandler,
	},
	Route{
		"ParticipantDelete",
		"DELETE",
		"/api/event/{eventId}/participants/{participantId}",
		ParticipantDeleteHandler,
	},
	Route{
		"ParticipantUpdate",
		"PUT",
		"/api/event/{eventId}/participants/{participantId}",
		ParticipantUpdateHandler,
	},

	// Stages
	Route{
		"Stages ",
		"GET",
		"/api/event/{eventId}/stages",
		StagesHandler,
	},
	Route{
		"Stage",
		"GET",
		"/api/event/{eventId}/stages/{stageId}",
		StageHandler,
	},
	Route{
		"StageCreate",
		"POST",
		"/api/event/{eventId}/stages",
		StageCreateHandler,
	},
	Route{
		"StageDelete",
		"DELETE",
		"/api/event/{eventId}/stages/{stagesId}",
		StageDeleteHandler,
	},
	Route{
		"StageUpdate",
		"PUT",
		"/api/event/{eventId}/stages/{stageId}",
		StageUpdateHandler,
	},

	// Sprints 
	Route{
		"Sprints",
		"GET",
		"/api/event/{eventId}/stages/{stageId}/sprints",
		SprintsHandler,
	},
	Route{
		"Stage",
		"GET",
		"/api/event/{eventId}/stages/{stageId}/sprints/{sprintId}",
		SprintHandler,
	},
	Route{
		"SprintCreate",
		"POST",
		"/api/event/{eventId}/stages/{stageId}/sprints",
		SprintCreateHandler,
	},
	Route{
		"SprintDelete",
		"DELETE",
		"/api/event/{eventId}/stages/{stageId}/sprints/{sprintId}}",
		SprintDeleteHandler,
	},
	Route{
		"SprintUpdate",
		"PUT",
		"/api/event/{eventId}/stages/{stageId}/sprints/{sprintId}",
		SprintUpdateHandler,
	},

	// Event commissaires
	Route{
		"Event commissaires",
		"GET",
		"/api/event/{eventId}/commissaires",
		SprintsHandler,
	},
	Route{
		"Event commissaire",
		"GET",
		"/api/event/{eventId}/commissaires/{commissaireId}",
		SprintHandler,
	},
	Route{
		"Event commissaire create",
		"POST",
		"/api/event/{eventId}/commissaires",
		SprintCreateHandler,
	},
	Route{
		"Event commissaire delete",
		"DELETE",
		"/api/event/{eventId}/commissaires/{commissaireId}",
		SprintDeleteHandler,
	},
	Route{
		"Event commissaire update",
		"PUT",
		"/api/event/{eventId}/commissaires/{commissaireId}",
		SprintUpdateHandler,
	},
}
