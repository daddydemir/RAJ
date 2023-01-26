package models

type Object struct {
	Name   string
	Fields []Field
}

type Field struct {
	Name string
	Type string
	/*
		attributes is not required for this time
		e.g.
		PK
		data length
	*/
}
