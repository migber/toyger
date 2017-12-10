package models

import (
	"fmt"
	"reflect"
	"testing"

	"gopkg.in/mgo.v2"
)

func TestCreateCommissaire(t *testing.T) {

	fmt.Println("GENERATED TEST CREATE COMMISSAIRE")
	type args struct {
		c       Commissaire
		session *mgo.Session
	}
	tests := []struct {
		name string
		args args
		want Commissaire
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CreateCommissaire(tt.args.c, tt.args.session); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CreateCommissaire() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetCommissaireInside(t *testing.T) {

	fmt.Println("GENERATED TEST GET INSIDE COMMISSAIRE")
	type args struct {
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
			if err := GetCommissaireInside(tt.args.id, tt.args.session); (err != nil) != tt.wantErr {
				t.Errorf("GetCommissaireInside() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestGetCommissairesList(t *testing.T) {
	fmt.Println("GENERATED TEST GET LIST COMMISSAIRE")

	type args struct {
		session *mgo.Session
	}
	tests := []struct {
		name string
		args args
		want Commissaires
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetCommissairesList(tt.args.session); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetCommissairesList() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetCommissaire(t *testing.T) {
	fmt.Println("GENERATED TEST GET COMMISSAIRE")

	type args struct {
		id      string
		session *mgo.Session
	}
	tests := []struct {
		name string
		args args
		want Commissaire
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetCommissaire(tt.args.id, tt.args.session); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetCommissaire() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUpdateCommissaire(t *testing.T) {
	fmt.Println("GENERATED TEST UPDATE COMMISSAIRE")
	type args struct {
		uid     string
		c       Commissaire
		session *mgo.Session
	}
	tests := []struct {
		name string
		args args
		want Commissaire
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := UpdateCommissaire(tt.args.uid, tt.args.c, tt.args.session); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UpdateCommissaire() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDeleteCommissaire(t *testing.T) {
	fmt.Println("GENERATED TEST DELETE COMMISSAIRE")	
	type args struct {
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
			if err := DeleteCommissaire(tt.args.uid, tt.args.session); (err != nil) != tt.wantErr {
				t.Errorf("DeleteCommissaire() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
