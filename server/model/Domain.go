package model

//Struct that represents a Domain Object for json response of API
type Domain struct {
	Servers 			[]Server 	`json:"servers,omitempty"`
	ServersChanged 		string 		`json:"servers_changed,omitempty"`
	SslGrade 			string 		`json:"ssl_grade,omitempty"`
	PreviousSslGrade 	string 		`json:"previous_ssl_grade,omitempty"`
	Logo 				string 		`json:"logo,omitempty"`
	Title 				string 		`json:"title,omitempty"`
	IsDown 				string 		`json:"is_down,omitempty"`
}