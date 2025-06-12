package models

import "database/sql"

func NullFrom[T comparable](t T) *sql.Null[T] {
	return &sql.Null[T]{
		Valid: t != *new(T),
		V:     t,
	}
}
