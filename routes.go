package main

import (
    "net/http"
)

type Route struct{
	Name 		string
	Method 		string
	Pattern		string
	HandlerFunc	http.HandlerFunc
}

type Routes []Route

var routes = Routes {
	Route{
		"Main",
		"GET",
		"/",
		MainHanlder,
	},
	Route{
		"Teams",
		"GET",
		"/api/teams",
		teamsHandler,
	},
	Route{
		"Team",
		"GET",
		"/api/teams/{teamId}",
		teamHandler,
	},
	Route{
		"TeamCreate",
		"POST",
		"/api/teams",
		teamCreateHandler,
	},
	Route{
		"DeleteTeam",
		"DELETE",
		"/api/teams/{teamId}",
		teamDeleteHandler,
	},
	Route{
		"UpdateTeam",
		"PUT",
		"/api/teams/{teamId}",
		teamUpdateHandler,
	},
}