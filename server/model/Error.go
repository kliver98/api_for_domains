package model

type Error struct {
	Type 	string `json:"type,omitempty"`
	Message string `json:"message,omitempty"`
	Code 	string `json:"code,omitempty"`
}