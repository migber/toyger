package main

import "regexp"

// IsValidUUID  check if UUID is valid
func IsValidUUID(uuid string) bool {
    r := regexp.MustCompile("^[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-4[a-fA-F0-9]{3}-[8|9|aA|bB][a-fA-F0-9]{3}-[a-fA-F0-9]{12}$")
    return r.MatchString(uuid)
}

//IsValidUCIID checks if UCIID is valid
func IsValidUCIID(uid string) bool {
    length := len(uid)
    if (length == 11){
        return true
    } 
    return false
}

// IsValidStageID checks if stage id is valid
func IsValidStageID(id int) bool {
    if (id > 0 && id < 30){
        return true
    } 
    return false
}

// IsValidRaceNumber check if race number is valid
func IsValidRaceNumber(id int) bool {
    if (id > 0 && id < 201){
        return true
    }
    return false
}