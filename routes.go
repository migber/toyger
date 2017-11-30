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
		CommissaireDeleteHandler,
	},
	Route{
		"CommissaireUpdate",
		"PUT",
		"/api/commissaires/{commissaireId}",
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
		"/api/events/participants",
		ParticipantsHandler,
	},
	Route{
		"Participant",
		"GET",
		"/api/events/{eventId}/participants/{participantId}",
		ParticipantHandler,
	},
	Route{
		"ParticipantCreate",
		"POST",
		"/api/events/{eventId}/participants",
		ParticipantCreateHandler,
	},
	Route{
		"ParticipantDelete",
		"DELETE",
		"/api/events/{eventId}/participants/{participantId}",
		ParticipantDeleteHandler,
	},
	Route{
		"ParticipantUpdate",
		"PUT",
		"/api/events/{eventId}/participants/{participantId}",
		ParticipantUpdateHandler,
	},

	// Stages
	Route{
		"Stages",
		"GET",
		"/api/events/{eventId}/stages",
		StagesHandler,
	},
	Route{
		"Stage",
		"GET",
		"/api/events/{eventId}/stages/{stageId}",
		StageHandler,
	},
	Route{
		"StageCreate",
		"POST",
		"/api/events/{eventId}/stages",
		StageCreateHandler,
	},
	Route{
		"StageDelete",
		"DELETE",
		"/api/events/{eventId}/stages/{stagesId}",
		StageDeleteHandler,
	},
	Route{
		"StageUpdate",
		"PUT",
		"/api/events/{eventId}/stages/{stageId}",
		StageUpdateHandler,
	},

	// Sprints 
	Route{
		"Sprints",
		"GET",
		"/api/events/{eventId}/stages/{stageId}/sprints",
		SprintsHandler,
	},
	Route{
		"Stage",
		"GET",
		"/api/events/{eventId}/stages/{stageId}/sprints/{sprintId}",
		SprintHandler,
	},
	Route{
		"SprintCreate",
		"POST",
		"/api/events/{eventId}/stages/{stageId}/sprints",
		SprintCreateHandler,
	},
	Route{
		"SprintDelete",
		"DELETE",
		"/api/events/{eventId}/stages/{stageId}/sprints/{sprintId}}",
		SprintDeleteHandler,
	},
	Route{
		"SprintUpdate",
		"PUT",
		"/api/events/{eventId}/stages/{stageId}/sprints/{sprintId}",
		SprintUpdateHandler,
	},

	// Event commissaires
	Route{
		"Event commissaires",
		"GET",
		"/api/events/{eventId}/commissaires",
		SprintsHandler,
	},
	Route{
		"Event commissaire",
		"GET",
		"/api/events/{eventId}/commissaires/{commissaireId}",
		SprintHandler,
	},
	Route{
		"Event commissaire create",
		"POST",
		"/api/events/{eventId}/commissaires",
		SprintCreateHandler,
	},
	Route{
		"Event commissaire delete",
		"DELETE",
		"/api/events/{eventId}/commissaires/{commissaireId}",
		SprintDeleteHandler,
	},
	Route{
		"Event commissaire update",
		"PUT",
		"/api/events/{eventId}/commissaires/{commissaireId}",
		SprintUpdateHandler,
	},
}
