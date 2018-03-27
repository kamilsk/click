package postgres

const dialect = "postgres"

// Dialect returns supported database dialect.
func Dialect() string {
	return dialect
}
