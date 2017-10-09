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
		"/teams",
		teamsHandler,
	},
	Route{
		"Team",
		"GET",
		"/teams/{teamId}",
		teamHandler,
	},
	Route{
		"TeamCreate",
		"POST",
		"/teams",
		teamCreate,
	},
}