package models

import "database/sql"

func NullFrom[T any](t T) *sql.Null[T] {
	return &sql.Null[T]{
		Valid: true,
		V:     t,
	}
}
