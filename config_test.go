package main

import (
	"fmt"
	"testing"
	"github.com/koding/multiconfig"
)

func TestConfig(t *testing.T) {
	
	fmt.Println("Test config")
	m := multiconfig.NewWithPath("./config/config.test.json")
	db := new(DB)
	m.MustLoad(db)

	assertEqual(t, db.Database, "toygertesting", "")
	assertEqual(t, db.Host, "ds129386.mlab.com:29386", "")
	assertEqual(t, db.Password, "test", "")
	assertEqual(t, db.Username, "test", "")
}

func assertEqual(t *testing.T, a interface{}, b interface{}, message string) {
	if a == b {
	  return
	  }
	  if len(message) == 0 {
		  message = fmt.Sprintf("%v != %v", a, b)
	  }
	  t.Fatal(message)
  }
 
func TestConnection(t *testing.T){
	fmt.Println("Check Connection return open session connection")
	session:= connection()
	sessionClone := session.Clone()
	sessionClone.Close()
	if (session == sessionClone){
		t.Error("Session is closed")
	}
}

