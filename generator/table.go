package generator

import (
	"fmt"

	"github.com/daddydemir/RAJ/models"
)

func CreateTable(model models.Object) string {

	query := "create table " + model.Name + " ("

	for _, o := range model.Fields {
		query += " " + o.Name + " " + o.Type
		if o != model.Fields[len(model.Fields)-1] {
			query += ", "
		}
	}
	query += ") ;"
	fmt.Println(query)
	return query
}
