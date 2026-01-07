package model

import "time"

type Student struct {
	ID			int			`json:"id"`
	Name		string		`json:"name"`
	Age			int			`json:"age"`
	DateOfBirth	time.Time	`json:"date_of_birth"`
}

type StudentRequest struct {
	Name		string	`json:"name"`
	Age			int		`json:"age"`
	DateOfBirth	string	`json:"date_of_birth"`  // yy-mm-dd
}

