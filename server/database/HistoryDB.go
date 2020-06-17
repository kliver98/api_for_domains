package database

import "time"

type HistoryDB struct {
	Domain 				string 		`json:"domain,omitempty"`
	SearchedAt 			*time.Time 	`json:"-"`
}