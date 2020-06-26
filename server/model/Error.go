package model

//Struct that represents an Error Object for json response of API
type Error struct {
	Type 	string `json:"type,omitempty"`
	Message string `json:"message,omitempty"`
	Code 	string `json:"code,omitempty"`
}