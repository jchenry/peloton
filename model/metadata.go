package model

type MetadataMappings struct {
	Instructors []Instructor `json:"instructors"`
	ClassTypes  []ClassType  `json:"class_types"`
}

type ClassType struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}
