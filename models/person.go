package models

type Person struct {
	ID   uint64 `json:"id"`
	Name string `json:"name"`
	Age  string `json:"age"`
}


type Persons []*Person
