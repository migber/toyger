package main

import (
	"gopkg.in/mgo.v2"
	//"gopkg.in/mgo.v2/bson"
	"time"
	"log"
	"github.com/koding/multiconfig"
)

// DB is db load structure
type DB struct {
	Host    string  		`required:"true"`
	Port    int
	Timeout  int            `default:"60"`
	Database string			`required:"true"`
	Username string			`required:"true"`
	Password string			`required:"true"`
}

func connection() (*mgo.Session){

	// config 
	m := multiconfig.NewWithPath("./config/config.json")
	db := new(DB)
	m.MustLoad(db)

	mongoDBDialInfo := &mgo.DialInfo{
		Addrs:    []string{db.Host},
		Timeout:  60 * time.Second,
		Database: db.Database,
		Username: db.Username,
		Password: db.Password,
	}	

	mongoSession, err := mgo.DialWithInfo(mongoDBDialInfo)
	if err != nil {
		log.Fatalf("CreateSession: %s\n", err)
	}

	return mongoSession.Clone()
}
