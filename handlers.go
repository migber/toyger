
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
// DATABASE name 
const DATABASE = "toyger"
//TEAMS team name in database
const TEAMS = "teams"
// EVENTS events table name in database
const EVENTS = "events"
// COMMISSAIRES table name in database
const COMMISSAIRES = "commissaires"
// STAGES table name in database 
const STAGES = "stages"
// SPRINTS table name in database
const SPRINTS = "sprints"
// CYCLISTS table name in database
const CYCLISTS = "cyclist"
// PARTICIPANTS table name in database 
const PARTICIPANTS = "participants"
// RACECOMMISSAIRE table name in database
const RACECOMMISSAIRE = "racecommissaire"
// CYCLISTSAlone
const CYCLISTSAlone = "cyclist"

// HealthCheck handler
func HealthCheck(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hey I am working here %s!", html.EscapeString(r.URL.Path))
}
// Hanlder main handler
func Hanlder(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %s", html.EscapeString(r.URL.Path))
}


// TeamHandler return one team handler
func TeamHandler(w http.ResponseWriter, r *http.Request) {
	validToken := true //checkJwt(w,r)
	if (validToken){

		vars := mux.Vars(r)
		teamID := vars["teamId"]

		w.Header().Set("Content-Type", "application/json;charset=UTF-8")

		if !IsValidUUID(teamID) {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		team, err := Models.GetTeam(teamID, connection(), DATABASE, TEAMS)
		if err !=  nil {
			w.WriteHeader(http.StatusNotFound)
		} else {
			w.WriteHeader(http.StatusOK)			
		}

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
	validToken := true //checkJwt(w,r)
	w.Header().Set("Content-Type", "application/json;charset=UTF-8")	
	if (validToken){
		
		teams, err2 := Models.GetTeamsList(connection(), DATABASE, TEAMS)
		if (err2 != nil){
			w.WriteHeader(http.StatusNotFound)
		} else {
			w.WriteHeader(http.StatusOK)
		}
		if err := json.NewEncoder(w).Encode(teams); err != nil {
			fmt.Println(err)
		}
	} else {
		w.WriteHeader(http.StatusUnauthorized)		
	}
}

// TeamCreateHandler create handler
func TeamCreateHandler(w http.ResponseWriter, r *http.Request){
	validToken := true //checkJwt(w,r)	
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
		t , errt:= Models.CreateTeam(team, connection(), DATABASE, TEAMS)
		if errt != nil {
			w.WriteHeader(http.StatusNotFound)
		} else {
			w.WriteHeader(http.StatusCreated)
		}
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	
		fmt.Println(r.URL.String())
		w.Header().Set("X-Frame-Options", "soooo" )
		if err := json.NewEncoder(w).Encode(t); err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
		
		if err := r.Body.Close(); err != nil {
			fmt.Println(err)
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
	
		err := Models.DeleteTeam(teamID, connection(), DATABASE, TEAMS)
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
		}
		teams, errt := Models.GetTeamsList(connection(), DATABASE, TEAMS)
		if errt != nil {
			w.WriteHeader(http.StatusBadRequest)
		}else {
		   w.WriteHeader(http.StatusOK)
		}
		if err := json.NewEncoder(w).Encode(teams); err != nil {
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
			fmt.Println(err)
		}
	
		if err := json.Unmarshal(body, &team); err != nil {
			w.Header().Set("Content-Type", "application/json; charset=UTF-8")
			w.WriteHeader(http.StatusBadRequest) // unprocessable entity
			if err := json.NewEncoder(w).Encode(err); err != nil {
				http.Error(w, err.Error(), 500)
				return
			}
		}
	
		t, err := Models.UpdateTeam(teamID, team, connection(), DATABASE, TEAMS)
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	
		w.Header().Set("Location", r.URL.String())
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
		} else {
			w.WriteHeader(http.StatusOK)
		}
		if err := json.NewEncoder(w).Encode(t); err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
	} else {
		w.WriteHeader(http.StatusUnauthorized)
	}
}



// CyclistAloneHandler return one team handler
func CyclistAloneHandler(w http.ResponseWriter, r *http.Request) {
	validToken := true //checkJwt(w,r)
	if (validToken){
		vars := mux.Vars(r)
		cyclistID := vars["cyclistId"]

		w.Header().Set("Content-Type", "application/json;charset=UTF-8")

		if !IsValidUCIID(cyclistID) {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		cyclist, err := Models.GetCyclistAlone(cyclistID, connection(), DATABASE, CYCLISTS)
		if err !=  nil {
			w.WriteHeader(http.StatusNotFound)
		} else {
			w.WriteHeader(http.StatusOK)			
		}

		if err := json.NewEncoder(w).Encode(cyclist); err != nil {
			http.Error(w, err.Error(), 500)
			return
		} 
 	} else {
		w.WriteHeader(http.StatusUnauthorized)		
	 }

}
// CyclistsAloneHandler return list of cyclists
func CyclistsAloneHandler(w http.ResponseWriter, r *http.Request) {
	validToken := true //checkJwt(w,r)
	w.Header().Set("Content-Type", "application/json;charset=UTF-8")	
	if (validToken){
		
		cyclists, err2 := Models.GetCyclistsListALone(connection(), DATABASE, CYCLISTS)
		if (err2 != nil){
			w.WriteHeader(http.StatusNotFound)
		} else {
			w.WriteHeader(http.StatusOK)
		}
		if err := json.NewEncoder(w).Encode(cyclists); err != nil {
			fmt.Println(err)
		}
	} else {
		w.WriteHeader(http.StatusUnauthorized)		
	}
}

// CyclistsAloneCreateHandler create handler
func CyclistsAloneCreateHandler(w http.ResponseWriter, r *http.Request){
	validToken := true //checkJwt(w,r)	
		if (validToken){
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
		t , errt:= Models.CreateCyclistAlone(cyclist, connection(), DATABASE, CYCLISTS)
		if errt != nil {
			w.WriteHeader(http.StatusNotFound)
		} else {
			w.WriteHeader(http.StatusCreated)
		}
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	
		fmt.Println(r.URL.String())
		w.Header().Set("X-Frame-Options", "soooo" )
		if err := json.NewEncoder(w).Encode(t); err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
		
		if err := r.Body.Close(); err != nil {
			fmt.Println(err)
		}
	} else {
		w.WriteHeader(http.StatusUnauthorized)
	}
}

// CyclistsAloneDeleteHandler team delete handler
func CyclistsAloneDeleteHandler(w http.ResponseWriter, r *http.Request){
	validToken := checkJwt(w,r)
	if (validToken){
		vars := mux.Vars(r)
		cyclistID := vars["cyclistId"]
	
		w.Header().Set("Content-Type", "application/json;charset=UTF-8")
	
		if !IsValidUCIID(cyclistID) {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
	
		err := Models.DeleteCyclistAlone(cyclistID, connection(), DATABASE, CYCLISTS)
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
		}
		cyclists, errt := Models.GetCyclistsListALone(connection(), DATABASE, CYCLISTS)
		if errt != nil {
			w.WriteHeader(http.StatusBadRequest)
		}else {
		   w.WriteHeader(http.StatusOK)
		}
		if err := json.NewEncoder(w).Encode(cyclists); err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
	} else {
		w.WriteHeader(http.StatusUnauthorized)
	}
}
// CyclistAloneUpdateHandler udpate team handler
func CyclistAloneUpdateHandler(w http.ResponseWriter, r *http.Request){
	validToken := checkJwt(w,r)
	if (validToken){
		vars := mux.Vars(r)
		cyclistID := vars["cyclistId"]
	
		w.Header().Set("Content-Type", "application/json;charset=UTF-8")
	
		if !IsValidUCIID(cyclistID) {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
	
		var cyclist Models.Cyclist
	
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			http.Error(w, err.Error(), 500)
		}
	
		if err := r.Body.Close(); err != nil {
			fmt.Println(err)
		}
	
		if err := json.Unmarshal(body, &cyclist); err != nil {
			w.Header().Set("Content-Type", "application/json; charset=UTF-8")
			w.WriteHeader(http.StatusBadRequest) // unprocessable entity
			if err := json.NewEncoder(w).Encode(err); err != nil {
				http.Error(w, err.Error(), 500)
				return
			}
		}
	
		t, err := Models.UpdateCyclistAlone(cyclistID, cyclist, connection(), DATABASE, CYCLISTS)
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	
		w.Header().Set("Location", r.URL.String())
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
		} else {
			w.WriteHeader(http.StatusOK)
		}
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
	validToken := true //checkJwt(w,r)
	if (validToken){
		vars := mux.Vars(r)
		teamID := vars["teamId"]
	
		w.Header().Set("Content-Type", "application/json;charset=UTF-8")
		cyclists, err := Models.GetCyclistsList(teamID, connection(), DATABASE, CYCLISTS)
		if err != nil{
			w.WriteHeader(http.StatusNotFound)
		} else {
			w.WriteHeader(http.StatusOK)
		}
		if err := json.NewEncoder(w).Encode(cyclists); err != nil {
			fmt.Println(err)
		}
	} else {
		w.WriteHeader(http.StatusUnauthorized)
	}
}

//CyclistHandler handler cyclist return on cyclist
func CyclistHandler(w http.ResponseWriter, r *http.Request) {
	validToken := true //checkJwt(w,r)
	if (validToken){
		vars := mux.Vars(r)
		teamID := vars["teamId"]
		cyclistID := vars["cyclistId"]
	
		w.Header().Set("Content-Type", "application/json;charset=UTF-8")
	
		if !IsValidUUID(teamID) {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
	
		if !IsValidUCIID(cyclistID){
			w.WriteHeader(http.StatusBadRequest)
			return
		}
	
		cyclist, err := Models.GetCyclist(teamID, cyclistID, connection(),DATABASE, CYCLISTS)
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
		} else {
			w.WriteHeader(http.StatusOK)
		}
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
		
		c, err := Models.CreateCyclist(teamID, cyclist, connection(), DATABASE, CYCLISTS)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
		}
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		if c.UCIID == ""{
			w.WriteHeader(http.StatusConflict)
		} else {
			w.WriteHeader(http.StatusOK)	
			updateErr := Models.InsertRider(teamID, c.UCIID, connection(),
										    DATABASE, TEAMS)
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
			fmt.Println(err)
		}
	} else {
		w.WriteHeader(http.StatusUnauthorized)
	}
}

//CyclistUpdateHandler handler update cyclist
func CyclistUpdateHandler(w http.ResponseWriter, r *http.Request) {
	validToken := true//checkJwt(w,r)
	if (validToken){
		vars := mux.Vars(r)
		teamID := vars["teamId"]
		cyclistID := vars["cyclistId"]
		w.Header().Set("Content-Type", "application/json;charset=UTF-8")
	
		if !IsValidUUID(teamID) {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		if !IsValidUCIID(cyclistID){
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		var cyclist Models.Cyclist
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			http.Error(w, err.Error(), 500)
		}
		if err := r.Body.Close(); err != nil {
			fmt.Println(err)
		}
	
		if err := json.Unmarshal(body, &cyclist); err != nil {
			w.Header().Set("Content-Type", "application/json; charset=UTF-8")
			w.WriteHeader(http.StatusBadRequest) // unprocessable entity
			if err := json.NewEncoder(w).Encode(err); err != nil {
				http.Error(w, err.Error(), 500)
				return
			}
		}
	
		c, err := Models.UpdateCyclist(teamID, cyclistID, cyclist, connection(), DATABASE, CYCLISTS)
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	
		w.Header().Set("Location", r.URL.String())
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
		} else {
			w.WriteHeader(http.StatusOK)
		}
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
		if !IsValidUCIID(cyclistID){
			w.WriteHeader(http.StatusBadRequest)
			return
		}
	
		err := Models.DeleteCyclist(teamID, cyclistID, connection(), DATABASE, CYCLISTS)
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
		}
		err = Models.DeleteRider(teamID, cyclistID, connection(),
								DATABASE, TEAMS)
		if err != nil{
			fmt.Println(fmt.Errorf("%s", err))
		}
		cyclists, err := Models.GetCyclistsList(teamID, connection(), DATABASE, CYCLISTS)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
		} else {
			w.WriteHeader(http.StatusOK)			
		}
		if err := json.NewEncoder(w).Encode(cyclists); err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
	} else {
		w.WriteHeader(http.StatusUnauthorized)
	}
}

// CommissairesHandler handler commissaires
func CommissairesHandler(w http.ResponseWriter, r *http.Request) {
	validToken := true //checkJwt(w,r)
	if (validToken){
		w.Header().Set("Content-Type", "application/json;charset=UTF-8")
		comm, err := Models.GetCommissairesList(connection())
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
		} else {
			w.WriteHeader(http.StatusOK)		
		}
		if err := json.NewEncoder(w).Encode(comm); err != nil {
			fmt.Println(err)
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
		comID := vars["commissaireID"]
	
		w.Header().Set("Content-Type", "application/json;charset=UTF-8")
	
		if !IsValidUCIID(comID) {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
	
		commissaire, err := Models.GetCommissaire(comID, connection())
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
		} else {
			w.WriteHeader(http.StatusOK)			
		}
	
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
		if !IsValidUCIID(commissaire.UCIID) {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		c, err := Models.CreateCommissaire(commissaire, connection())
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
		} else {
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
				fmt.Println(err)
			}
		}
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		
	} else {
		w.WriteHeader(http.StatusUnauthorized)
	}
}

// CommissaireUpdateHandler handler update commissaire information
func CommissaireUpdateHandler(w http.ResponseWriter, r *http.Request) {
	validToken := checkJwt(w,r)
	if (validToken){
		vars := mux.Vars(r)
		comID := vars["commissaireID"]
		
		w.Header().Set("Content-Type", "application/json;charset=UTF-8")
	
		if !IsValidUCIID(comID) {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
	
		var commissaire Models.Commissaire
	
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			http.Error(w, err.Error(), 500)
		}
	
		if err := r.Body.Close(); err != nil {
			fmt.Println(err)
		}
	
		if err := json.Unmarshal(body, &commissaire); err != nil {
			w.Header().Set("Content-Type", "application/json; charset=UTF-8")
			w.WriteHeader(http.StatusBadRequest) // unprocessable entity
			if err := json.NewEncoder(w).Encode(err); err != nil {
				http.Error(w, err.Error(), 500)
				return
			}
		}
		c, err := Models.UpdateCommissaire(comID, commissaire, connection())
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
		} else {
			w.WriteHeader(http.StatusOK)
		}
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	
		w.Header().Set("Location", r.URL.String())
		
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
		comID := vars["commissaireID"]
	
		w.Header().Set("Content-Type", "application/json;charset=UTF-8")
	
		if !IsValidUCIID(comID) {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
	
		err := Models.DeleteCommissaire(comID, connection())
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
		}
		comm, err := Models.GetCommissairesList(connection())
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
		} else {
			w.WriteHeader(http.StatusOK)			
		}
		if err := json.NewEncoder(w).Encode(comm); err != nil {
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
	validToken := true//checkJwt(w,r)
	if (validToken){
		w.Header().Set("Content-Type", "application/json;charset=UTF-8")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		events, err := m.GetEventsList(connection(), DATABASE, EVENTS)
		if err != nil {
			w.WriteHeader(http.StatusNotFound)	
		} else  {
			w.WriteHeader(http.StatusOK)		
		}
		if err := json.NewEncoder(w).Encode(events); err != nil {
			fmt.Println(err)
		}
	} else {
		w.WriteHeader(http.StatusUnauthorized)
	}
}

// EventHandler handler event information
func EventHandler(w http.ResponseWriter, r *http.Request) {
	validToken := true //checkJwt(w,r)
	if (validToken){
		vars := mux.Vars(r)
		eventID := vars["eventID"]
	
		w.Header().Set("Content-Type", "application/json;charset=UTF-8")
	
		if !IsValidUUID(eventID) {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
	
		event, err := m.GetEvent(eventID, connection(), DATABASE, EVENTS)
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
		} else {
		 w.WriteHeader(http.StatusOK)		
		}
	
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
	validToken := true //checkJwt(w,r)
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
		 e, err := m.CreateEvent(event, connection(), DATABASE, EVENTS)
		 if err != nil {
			 w.WriteHeader(http.StatusBadGateway)
		 } else {
			w.WriteHeader(http.StatusCreated)			
		 }
		 w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	
		 fmt.Println(r.URL.String())
		 w.Header().Set("X-Frame-Options", "soooo" )
		 if err := json.NewEncoder(w).Encode(e); err != nil {
			 http.Error(w, err.Error(), 500)
			 return
			}
	 
		 if err := r.Body.Close(); err != nil {
			 fmt.Println(err)
		}
	} else {
		w.WriteHeader(http.StatusUnauthorized)
	}
}

// EventDeleteHandler handler delete event information
func EventDeleteHandler(w http.ResponseWriter, r *http.Request) {
	validToken := true //checkJwt(w,r)
	if (validToken){
		vars := mux.Vars(r)
		eventID := vars["eventID"]
	
		w.Header().Set("Content-Type", "application/json;charset=UTF-8")
	
		if !IsValidUUID(eventID) {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
	
		err := m.DeleteEvent(eventID, connection(), DATABASE, EVENTS)
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
		}
		
		events, err := m.GetEventsList(connection(), DATABASE, EVENTS)
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
		} else {
			w.WriteHeader(http.StatusOK)
		}
		if err := json.NewEncoder(w).Encode(events); err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
	} else {
		w.WriteHeader(http.StatusUnauthorized)
	}
}

// EventUpdateHandler handler update event information
func EventUpdateHandler(w http.ResponseWriter, r *http.Request) {
	validToken := true //checkJwt(w,r)
	if (validToken){
		vars := mux.Vars(r)
		eventID := vars["eventID"]
	
		w.Header().Set("Content-Type", "application/json;charset=UTF-8")
	
		if !IsValidUUID(eventID) {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
	
		var event m.Event
	
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			http.Error(w, err.Error(), 500)
		}
	
		if err := r.Body.Close(); err != nil {
			fmt.Println(err)
		}
	
		if err := json.Unmarshal(body, &event); err != nil {
			w.Header().Set("Content-Type", "application/json; charset=UTF-8")
			w.WriteHeader(http.StatusBadRequest) // unprocessable entity
			if err := json.NewEncoder(w).Encode(err); err != nil {
				http.Error(w, err.Error(), 500)
				return
			}
		}
	
		e, err  := m.UpdateEvent(eventID, event, connection(), DATABASE, EVENTS)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
		} else {
			w.WriteHeader(http.StatusOK)
		}
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	
		w.Header().Set("Location", r.URL.String())
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
	validToken := true //checkJwt(w,r)
	if (validToken){
		vars := mux.Vars(r)
		eventID := vars["eventID"]
	
		w.Header().Set("Content-Type", "application/json;charset=UTF-8")
		participants , err := m.GetParticipantsList(eventID, connection(),DATABASE, PARTICIPANTS)
		if err != nil {
			w.WriteHeader(http.StatusNotFound)			
		} else {
			w.WriteHeader(http.StatusOK)		
		}
		if err := json.NewEncoder(w).Encode(participants); err != nil {
			fmt.Println(err)
		}
	} else {
		w.WriteHeader(http.StatusUnauthorized)
	}
}

// ParticipantHandler handler participant information
func ParticipantHandler(w http.ResponseWriter, r *http.Request) {
	validToken := true //checkJwt(w,r)
	if(validToken){
		vars := mux.Vars(r)
		eventID := vars["eventID"]
		participantID := vars["participantId"]
	
		w.Header().Set("Content-Type", "application/json;charset=UTF-8")
	
		if !IsValidUUID(eventID) {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
	
		participantNo, _ := strconv.Atoi(participantID)
		if !IsValidRaceNumber(participantNo){
			w.WriteHeader(http.StatusBadRequest)
			return
		}
	
		participant, err := m.GetParticipant(eventID, participantNo, connection(),
										DATABASE, PARTICIPANTS)
		if err != nil {
			w.WriteHeader(http.StatusNotFound)					
		} else {
		 	w.WriteHeader(http.StatusOK)		
		}
	
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
	validToken := true //checkJwt(w,r)
	if (validToken) {
		vars := mux.Vars(r)
		eventID := vars["eventID"]
		
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
			}
		}
		
		p, err := m.CreateParticipant(eventID, participant, connection(),
								DATABASE, PARTICIPANTS)
	    if err != nil {
			w.WriteHeader(http.StatusBadRequest)
		} else {
			if strconv.Itoa(p.No) == ""{
				w.WriteHeader(http.StatusConflict)
			} else {
				w.WriteHeader(http.StatusOK)	
				updateErr := m.InsertEventParticipants(eventID, p.No, connection(), DATABASE, EVENTS)
				if updateErr != nil {
					fmt.Println("Something went wrong.")
				}
				if err := json.NewEncoder(w).Encode(p); err != nil {
					http.Error(w, err.Error(), 500)
				}
			}
		}
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		fmt.Println(r.URL.String())
		if err := r.Body.Close(); err != nil {
			fmt.Println(err)
		}
	} else {
		w.WriteHeader(http.StatusUnauthorized)
	}
}

// ParticipantDeleteHandler handler delete participant information
func ParticipantDeleteHandler(w http.ResponseWriter, r *http.Request) {
	validToken := true //checkJwt(w,r)
	if (validToken){
		vars := mux.Vars(r)
		eventID := vars["eventID"]
		participantID := vars["participantId"]
		w.Header().Set("Content-Type", "application/json;charset=UTF-8")
	
		if !IsValidUUID(eventID) {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		participantNo, _ := strconv.Atoi(participantID)
		if !IsValidStageID(participantNo){
			w.WriteHeader(http.StatusBadRequest)
			return
		}
	
		err := m.DeleteParticipant(eventID, participantNo, connection(),
								  DATABASE, PARTICIPANTS)
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
		}
		err = m.DeleteEventParticipant(eventID, participantNo, connection(),
									   DATABASE, EVENTS)
		if err != nil{
			fmt.Println(fmt.Errorf("%s", err))
		}
		participants, err := m.GetParticipantsList(eventID, connection(),
												   DATABASE, PARTICIPANTS)
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
		} else {
			w.WriteHeader(http.StatusOK)
		}
		if err := json.NewEncoder(w).Encode(participants); err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
	} else {
		w.WriteHeader(http.StatusUnauthorized)
	}
}

// ParticipantUpdateHandler handler update participant information
func ParticipantUpdateHandler(w http.ResponseWriter, r *http.Request) {
	validToken := true //checkJwt(w,r)
	if (validToken){
		vars := mux.Vars(r)
		eventID := vars["eventID"]
		participantID := vars["participantId"]
		w.Header().Set("Content-Type", "application/json;charset=UTF-8")
	
		if !IsValidUUID(eventID) {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		participantNo, _ := strconv.Atoi(participantID)
		if !IsValidStageID(participantNo){
			w.WriteHeader(http.StatusBadRequest)
			return
		}
	
		var particicipant m.Participant
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			http.Error(w, err.Error(), 500)
		}
		if err := r.Body.Close(); err != nil {
			fmt.Println(err)
		}
	
		if err := json.Unmarshal(body, &particicipant); err != nil {
			w.Header().Set("Content-Type", "application/json; charset=UTF-8")
			w.WriteHeader(http.StatusBadRequest) // unprocessable entity
			if err := json.NewEncoder(w).Encode(err); err != nil {
				http.Error(w, err.Error(), 500)
				return
			}
		}
	
		p, err := m.UpdateParticipant(eventID, participantNo, particicipant, connection(),
								DATABASE, PARTICIPANTS)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
		} else {
			w.WriteHeader(http.StatusOK)		
		}
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	
		w.Header().Set("Location", r.URL.String())
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
	validToken := true //checkJwt(w,r)
	if (validToken) {
		vars := mux.Vars(r)
		eventID := vars["eventID"]
	
		w.Header().Set("Content-Type", "application/json;charset=UTF-8")
		stages, err := m.GetStageList(eventID, connection(),
									  DATABASE, STAGES)
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
		} else {
			w.WriteHeader(http.StatusOK)		
		}
		if err := json.NewEncoder(w).Encode(stages); err != nil {
			fmt.Println(err)
		}
	} else {
		w.WriteHeader(http.StatusUnauthorized)
	}
}

// StageHandler handler stage information 
func StageHandler(w http.ResponseWriter, r *http.Request) {
	validToken := true //checkJwt(w,r)
	if (validToken) {
		vars := mux.Vars(r)
		eventID := vars["eventID"]
		stageID := vars["stageID"]
	
		w.Header().Set("Content-Type", "application/json;charset=UTF-8")
		if !IsValidUUID(eventID) {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
	
		intstageID, _ := strconv.Atoi(stageID)
		if !IsValidStageID(intstageID){
			w.WriteHeader(http.StatusBadRequest)
			return
		}
	
		stage, err := m.GetStage(eventID, intstageID, connection(),
							DATABASE, STAGES)
		if err != nil {
			w.WriteHeader(http.StatusNotFound)			
		} else {
			w.WriteHeader(http.StatusOK)
		}
	
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
	validToken := true //checkJwt(w,r)
	if (validToken){
		vars := mux.Vars(r)
		eventID := vars["eventID"]
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
		
		s, err := m.CreateStage(eventID, stage, connection(), DATABASE, STAGES)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
		} else {
			if strconv.Itoa(s.ID) == ""{
				w.WriteHeader(http.StatusConflict)
			} else {
				w.WriteHeader(http.StatusOK)	
				updateErr := m.InsertEventStages(eventID, s.ID, connection(), DATABASE, EVENTS)
				if updateErr != nil {
					fmt.Println("Something went wrong.")
				}
				if err := json.NewEncoder(w).Encode(s); err != nil {
					http.Error(w, err.Error(), 500)
					return
				}
			}
		}
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		fmt.Println(r.URL.String())
		if err := r.Body.Close(); err != nil {
			fmt.Println(err)
		}
	} else {
		w.WriteHeader(http.StatusUnauthorized)
	}
}

// StageDeleteHandler handler delete stage
func StageDeleteHandler(w http.ResponseWriter, r *http.Request) {
	validToken := true //checkJwt(w,r)
	if (validToken){
		vars := mux.Vars(r)
		eventID := vars["eventID"]
		stageID := vars["stageID"]
		w.Header().Set("Content-Type", "application/json;charset=UTF-8")
	
		if !IsValidUUID(eventID) {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		intstageID, _ := strconv.Atoi(stageID)
		if !IsValidStageID(intstageID){
			w.WriteHeader(http.StatusBadRequest)
			return
		}
	
		err := m.DeleteStage(eventID, intstageID, connection(), DATABASE, STAGES)
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
		}
		err = m.DeleteEventStage(eventID, intstageID, connection(), DATABASE, EVENTS)
		if err != nil{
			fmt.Println(fmt.Errorf("%s", err))
		}
		stages, err := m.GetStageList(eventID, connection(),
									  DATABASE, STAGES)
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
		} else {
			w.WriteHeader(http.StatusOK)		
		}
		if err := json.NewEncoder(w).Encode(stages); err != nil {
			http.Error(w, err.Error(), 500)
			return
		}	
	} else {
		w.WriteHeader(http.StatusUnauthorized)
	}
}

// StageUpdateHandler handler update stage information 
func StageUpdateHandler(w http.ResponseWriter, r *http.Request) {
	validToken := true //checkJwt(w,r)
	if (validToken){
		vars := mux.Vars(r)
		eventID := vars["eventID"]
		stageID := vars["stageID"]
		w.Header().Set("Content-Type", "application/json;charset=UTF-8")
	
		if !IsValidUUID(eventID) {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		intstageID, _ := strconv.Atoi(stageID)
		if !IsValidStageID(intstageID){
			w.WriteHeader(http.StatusBadRequest)
			return
		}
	
		var stage m.Stage
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			http.Error(w, err.Error(), 500)
		}
		if err := r.Body.Close(); err != nil {
			fmt.Println(err)
		}
	
		if err := json.Unmarshal(body, &stage); err != nil {
			w.Header().Set("Content-Type", "application/json; charset=UTF-8")
			w.WriteHeader(http.StatusBadRequest) // unprocessable entity
			if err := json.NewEncoder(w).Encode(err); err != nil {
				http.Error(w, err.Error(), 500)
				return
			}
		}
	
		s, err := m.UpdateStage(eventID, intstageID, stage, connection(),
						  DATABASE, STAGES)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
		} else {
			w.WriteHeader(http.StatusOK)		
		}
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	
		w.Header().Set("Location", r.URL.String())
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
	validToken := true //checkJwt(w,r)
	if (validToken) {
		vars := mux.Vars(r)
		eventID := vars["eventID"]
		stageID := vars["stageID"]
	
		intstageID, _ := strconv.Atoi(stageID)
	
		w.Header().Set("Content-Type", "application/json;charset=UTF-8")
		sprints, err := m.GetSprintList(eventID, intstageID, connection(),
										DATABASE, SPRINTS)
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
		} else {
			w.WriteHeader(http.StatusOK)			
		}
		if err := json.NewEncoder(w).Encode(sprints); err != nil {
			fmt.Println(err)
		}
	} else {
		w.WriteHeader(http.StatusUnauthorized)
	}
}

// SprintHandler handler sprint information
func SprintHandler(w http.ResponseWriter, r *http.Request) {
	validToken := true //checkJwt(w,r)
	if (validToken) {
		vars := mux.Vars(r)
		eventID := vars["eventID"]
		stageID := vars["stageID"]
		sprintID := vars["sprintID"]
	
		w.Header().Set("Content-Type", "application/json;charset=UTF-8")
	
		if !IsValidUUID(eventID) {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		intstageID, _ := strconv.Atoi(stageID)
		if !IsValidStageID(intstageID){
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		if !IsValidUUID(sprintID) {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
	
		sprint, err := m.GetSprint(eventID, intstageID, sprintID, connection(),
							 DATABASE, SPRINTS)
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
		} else {
			w.WriteHeader(http.StatusOK)
		}
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
	validToken := true //checkJwt(w,r)
	if (validToken){
		vars := mux.Vars(r)
		eventID := vars["eventID"]
		stageID := vars["stageID"]
		intstageID, _ := strconv.Atoi(stageID)
	
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
		
		s, err := m.CreateSprint(eventID, intstageID, sprint, connection(),
							DATABASE, SPRINTS)
		if err != nil { 
			w.WriteHeader(http.StatusBadRequest)
		} else {
			if s.Id == ""{
				w.WriteHeader(http.StatusConflict)
			} else {
				w.WriteHeader(http.StatusOK)	
				updateErr := m.AddSprints(eventID, intstageID, sprint, connection(),
										 DATABASE, SPRINTS)
				if updateErr != nil {
					fmt.Println("Something went wrong.")
				}
				if err := json.NewEncoder(w).Encode(s); err != nil {
					http.Error(w, err.Error(), 500)
					return
				}
			}
		}
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		
		fmt.Println(r.URL.String())
		if err := r.Body.Close(); err != nil {
			fmt.Println(err)
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
		eventID := vars["eventID"]
		stageID := vars["stageID"]
		sprintID := vars["sprintID"]
		w.Header().Set("Content-Type", "application/json;charset=UTF-8")
	
		if !IsValidUUID(eventID) {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		intstageID, _ := strconv.Atoi(stageID)
		if !IsValidStageID(intstageID){
			w.WriteHeader(http.StatusBadRequest)
			return
		}
	
		err := m.DeleteSprint(eventID, intstageID, sprintID, connection(),
							 DATABASE, SPRINTS)
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
		}
		err = m.DeleteStageSprint(eventID, intstageID, sprintID, connection(),
								 DATABASE, SPRINTS)
		if err != nil{
			fmt.Println(fmt.Errorf("%s", err))
		}
		
		sprints, err := m.GetSprintList(eventID, intstageID, connection(),
										DATABASE, SPRINTS)
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
		} else {
			w.WriteHeader(http.StatusOK)			
		}
		if err := json.NewEncoder(w).Encode(sprints); err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
	} else {
		w.WriteHeader(http.StatusUnauthorized)
	}
}

// SprintUpdateHandler handler update sprint information
func SprintUpdateHandler(w http.ResponseWriter, r *http.Request) {
	validToken := true// checkJwt(w,r)
	if (validToken){
		vars := mux.Vars(r)
		eventID := vars["eventID"]
		stageID := vars["stageID"]
		sprintID := vars["sprintID"]
		w.Header().Set("Content-Type", "application/json;charset=UTF-8")
	
		if !IsValidUUID(eventID) {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		intstageID, _ := strconv.Atoi(stageID)
		if !IsValidStageID(intstageID){
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		if !IsValidUUID(eventID) {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
	
		var sprint m.Sprint
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			http.Error(w, err.Error(), 500)
		}
		if err := r.Body.Close(); err != nil {
			fmt.Println(err)
		}
	
		if err := json.Unmarshal(body, &sprint); err != nil {
			w.Header().Set("Content-Type", "application/json; charset=UTF-8")
			w.WriteHeader(http.StatusBadRequest) // unprocessable entity
			if err := json.NewEncoder(w).Encode(err); err != nil {
				http.Error(w, err.Error(), 500)
				return
			}
		}
	
		s, err := m.UpdateSprint(eventID, intstageID, sprintID, sprint, connection(),
						   DATABASE, SPRINTS)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
		} else {
			w.WriteHeader(http.StatusOK)		
		}
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		err = m.UpdateStageSprint(intstageID, sprintID, sprint, connection(),
								 DATABASE, SPRINTS)
		if err != nil{
			fmt.Println(fmt.Errorf("%s", err))
		}
		w.Header().Set("Location", r.URL.String())
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
	validToken := true //checkJwt(w,r)
	if (validToken){
		vars := mux.Vars(r)
		eventID := vars["eventID"]
	
		w.Header().Set("Content-Type", "application/json;charset=UTF-8")
		eventComm, err := m.GetRaceCommissairesList(eventID, connection(), 
													DATABASE, RACECOMMISSAIRE)
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
		} else {
			w.WriteHeader(http.StatusOK)		
		}
		if err := json.NewEncoder(w).Encode(eventComm); err != nil {
			fmt.Println(err)
		}	
	} else {
		w.WriteHeader(http.StatusUnauthorized)
	}
}

// EventCommissaireHandler handler event commissaire information
func EventCommissaireHandler(w http.ResponseWriter, r *http.Request) {
	validToken := true // checkJwt(w,r)
	if (validToken){
		vars := mux.Vars(r)
		eventID := vars["eventID"]
		raceCommID := vars["commissaireID"]
	
		w.Header().Set("Content-Type", "application/json;charset=UTF-8")
	
		if !IsValidUUID(eventID) {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
	
		if !IsValidUCIID(raceCommID) {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
	
		raceCommissaire, err := m.GetRaceCommissaire(eventID, raceCommID, connection(),
												DATABASE, RACECOMMISSAIRE)
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
		} else {
			w.WriteHeader(http.StatusOK)
		}
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
	validToken := true //checkJwt(w,r)
	if (validToken){
		vars := mux.Vars(r)
		eventID := vars["eventID"]
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
		
		s, err:= m.CreateRaceCommissaire(eventID, raceCommissaire, connection(), DATABASE, RACECOMMISSAIRE)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
		} else {
			if s.Commissaire.UCIID == ""{
				w.WriteHeader(http.StatusConflict)
			} else {
				w.WriteHeader(http.StatusOK)	
				updateErr := m.InsertEventCommissaire(eventID, s.Commissaire.UCIID, connection(), DATABASE, EVENTS)
				if updateErr != nil {
					fmt.Println("Something went wrong.")
				}
				if err := json.NewEncoder(w).Encode(s); err != nil {
					http.Error(w, err.Error(), 500)
					return
				}
			}
		}
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		
		fmt.Println(r.URL.String())
		if err := r.Body.Close(); err != nil {
			fmt.Println(err)
		}
	} else {
		w.WriteHeader(http.StatusUnauthorized)
	}
}

// EventCommissaireDeleteHandler handler delete event commissaire 
func EventCommissaireDeleteHandler(w http.ResponseWriter, r *http.Request) {
	validToken := true //checkJwt(w,r)
	if (validToken) {
		vars := mux.Vars(r)
		eventID := vars["eventID"]
		rcID := vars["commissaireID"]
		w.Header().Set("Content-Type", "application/json;charset=UTF-8")
	
		if !IsValidUUID(eventID) {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		if !IsValidUCIID(rcID){
			w.WriteHeader(http.StatusBadRequest)
			return
		}
	
		err := m.DeleteRaceCommissaire(eventID, rcID, connection(), DATABASE, RACECOMMISSAIRE)
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
		}
		err = m.DeleteEventCommissaire(eventID, rcID, connection(), DATABASE, EVENTS)
		if err != nil{
			fmt.Println(fmt.Errorf("%s", err))
		}
		comms, err := m.GetRaceCommissairesList(eventID, connection(), DATABASE, RACECOMMISSAIRE)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
		} else {
			w.WriteHeader(http.StatusOK)		
		}
		if err := json.NewEncoder(w).Encode(comms); err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
	} else {
		w.WriteHeader(http.StatusUnauthorized)
	}
}

// EventCommissaireUpdateHandler handler update event commissaire information
func EventCommissaireUpdateHandler(w http.ResponseWriter, r *http.Request) {
	validToken := true //checkJwt(w,r)
	if (validToken) {
		vars := mux.Vars(r)
		eventID := vars["eventID"]
		commissaireID := vars["commissaireID"]
		w.Header().Set("Content-Type", "application/json;charset=UTF-8")
	
		if !IsValidUUID(eventID) {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		if !IsValidUCIID(commissaireID){
			w.WriteHeader(http.StatusBadRequest)
			return
		}
	
		var raceCommiss m.RaceCommissaire
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			http.Error(w, err.Error(), 500)
		}
		if err := r.Body.Close(); err != nil {
			fmt.Println(err)
		}
	
		if err := json.Unmarshal(body, &raceCommiss); err != nil {
			w.Header().Set("Content-Type", "application/json; charset=UTF-8")
			w.WriteHeader(http.StatusBadRequest) // unprocessable entity
			if err := json.NewEncoder(w).Encode(err); err != nil {
				http.Error(w, err.Error(), 500)
				return
			}
		}
	
		rc, err := m.UpdateRaceCommissaire(eventID, commissaireID, raceCommiss, connection(), DATABASE, RACECOMMISSAIRE)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
		} else {
			w.WriteHeader(http.StatusOK)		
		}
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	
		w.Header().Set("Location", r.URL.String())
		if err := json.NewEncoder(w).Encode(rc); err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
	} else {
		w.WriteHeader(http.StatusUnauthorized)
	}
}
