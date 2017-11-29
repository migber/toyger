package main

import "regexp"

func IsValidUUID(uuid string) bool {
    r := regexp.MustCompile("^[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-4[a-fA-F0-9]{3}-[8|9|aA|bB][a-fA-F0-9]{3}-[a-fA-F0-9]{12}$")
    return r.MatchString(uuid)
}

func isValidUCIID(uid string) bool {
    length := len(uid)
    if (length == 11){
        return true
    } else {
        return false
    }
}
