package models

import "database/sql"

type Person struct {
	ID       uint64         `json:"id"`
	Name     string         `json:"name"`
	Age      string         `json:"age"`
	Password string         `json:"password"`
	Details  sql.NullString `json:"details"`
}

type Persons []*Person
