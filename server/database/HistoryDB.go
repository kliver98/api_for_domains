package database

import "time"

//Struct that represents a History Object from database
type HistoryDB struct {
	Domain 				string 		`json:"domain,omitempty"`
	SearchedAt 			*time.Time 	`json:"-"`
}