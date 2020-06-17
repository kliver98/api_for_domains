package database

import "time"

type DomainDB struct {
	Name 				string 		`json:"name,omitempty"`
	SslGrade 			string 		`json:"ssl_grade,omitempty"`
	PreviousSslGrade 	string 		`json:"previous_ssl_grade,omitempty"`
	SearchedAt 			*time.Time 	`json:"-"`
}