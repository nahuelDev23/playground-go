package models

import "database/sql"

type Person struct {
	ID   uint64 `json:"id"`
	Name string `json:"name"`
	Age  string `json:"age,omitempy"`
	Details  sql.NullString `json:"details,omitempy"`
}


type Persons []*Person
