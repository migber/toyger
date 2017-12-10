package models

import (
	"fmt"
	"reflect"
	"testing"

	"gopkg.in/mgo.v2"
)

func TestCreateCyclist(t *testing.T) {

	fmt.Println("++++++++++ GENERATED TEST")
	type args struct {
		teamId  string
		c       Cyclist
		session *mgo.Session
	}
	tests := []struct {
		name string
		args args
		want Cyclist
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CreateCyclist(tt.args.teamId, tt.args.c, tt.args.session); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CreateCyclist() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetCyclistsList(t *testing.T) {
	fmt.Println("GENERATED TEST GET LIST CYCLIST")

	type args struct {
		teamId  string
		session *mgo.Session
	}
	tests := []struct {
		name string
		args args
		want Cyclists
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetCyclistsList(tt.args.teamId, tt.args.session); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetCyclistsList() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetCyclist(t *testing.T) {
	fmt.Println("GENERATED TEST GET CYCLIST")

	type args struct {
		teamId  string
		id      string
		session *mgo.Session
	}
	tests := []struct {
		name string
		args args
		want Cyclist
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetCyclist(tt.args.teamId, tt.args.id, tt.args.session); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetCyclist() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetCyclistInside(t *testing.T) {
	fmt.Println("GENERATED TEST GET INSIDE CYCLIST")

	type args struct {
		teamId  string
		id      string
		session *mgo.Session
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := GetCyclistInside(tt.args.teamId, tt.args.id, tt.args.session); (err != nil) != tt.wantErr {
				t.Errorf("GetCyclistInside() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestUpdateCyclist(t *testing.T) {
	fmt.Println("GENERATED TEST UPDATE CYCLIST")

	type args struct {
		teamId  string
		uid     string
		c       Cyclist
		session *mgo.Session
	}
	tests := []struct {
		name string
		args args
		want Cyclist
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := UpdateCyclist(tt.args.teamId, tt.args.uid, tt.args.c, tt.args.session); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UpdateCyclist() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDeleteCyclist(t *testing.T) {
	
	fmt.Println("GENERATED TEST DELETE CYCLIST")
	
	type args struct {
		teamId  string
		uid     string
		session *mgo.Session
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := DeleteCyclist(tt.args.teamId, tt.args.uid, tt.args.session); (err != nil) != tt.wantErr {
				t.Errorf("DeleteCyclist() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
