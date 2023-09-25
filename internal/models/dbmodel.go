package models

type Object struct {
	Name   string  // struct model
	Fields []Field // table Model
}

type Field struct {
	Name      string
	Type      string
	Encrypted bool
	Show      bool
	Default   string
	/*
		attributes is not required for this time
		e.g.
		PK
		data length
	*/
}

const (
	NUMBER  string = "number"
	STRING  string = "varchar2"
	DATE    string = "date"
	BOOLEAN string = "boolean"
)
