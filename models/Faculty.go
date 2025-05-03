package models

type Faculty struct {
	ID           int
	Name         string
	MobileNumber string
	Email        string
	Address      string
	Type         string // "ACADEMIC" or "NON ACADEMIC"
	Title        string
}

