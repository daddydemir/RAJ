package generator

import "RAJ/models"

type Generate interface {
	GenerateMySQL(model models.Object) string
	GenerateOracleDB(model models.Object) string
	GenerateSQLServer(model models.Object) string
}
