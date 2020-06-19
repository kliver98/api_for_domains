package model

//https://qualys-secure.force.com/discussions/s/article/000005828
var SSL_GRADE = map[string]int{
	"A+":0,
	"A":1,
	"B":2,
	"C":3,
	"D":4,
	"E":5,
	"F":6,
}
const MAX_SSL_GRADE = "A+"
const MIN_SSL_GRADE = "F"

type Server struct {
	Address 	string 	`json:"address,omitempty"`
	SslGrade 	string 	`json:"ssl_grade,omitempty"`
	Country 	string 	`json:"country,omitempty"`
	Owner 		string 	`json:"owner,omitempty"`
}