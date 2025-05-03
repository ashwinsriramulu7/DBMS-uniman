package models

type Student struct {
	ID             int
	Name           string
	MobileNumber   string
	Email          string
	ProgramEnrolled string
	Type           string // "UG", "PG", "Phd"
}

