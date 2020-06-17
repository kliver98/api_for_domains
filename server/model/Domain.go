package model

type Domain struct {
	Servers 			[]Server 	`json:"servers,omitempty"`
	ServersChanged 		bool 		`json:"servers_changed,omitempty"`
	SslGrade 			string 		`json:"ssl_grade,omitempty"`
	PreviousSslGrade 	string 		`json:"previous_ssl_grade,omitempty"`
	Logo 				string 		`json:"logo,omitempty"`
	Title 				string 		`json:"title,omitempty"`
	IsDown 				bool 		`json:"is_down,omitempty"`
}