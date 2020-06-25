package model

//Struct that represents a History Object for json response of API
type History struct {
	Items []string 	`json:"items,omitempty"`
	Error Error 	`json:"error,omitempty"`
}
