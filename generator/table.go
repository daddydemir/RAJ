package generator

import (
	"fmt"

	"github.com/daddydemir/RAJ/models"
)

func CreateTable(model models.Object) string {

	query := "create table " + model.Name + " ("

	for _, o := range model.Fields {
		if o.Default == "" {
			query += " " + o.Name + " " + o.Type
		} else {
			query += " " + o.Name + " " + o.Type + " DEFAULT " + o.Default
		}
		if o != model.Fields[len(model.Fields)-1] {
			query += ", "
		}
	}
	query += ") ;"
	fmt.Println(query)
	return query
}
