package schema

import (
	"github.com/jmoiron/sqlx"
)

func NewRepository[E any](db *sqlx.DB, se E) Repository[E] {
	return Repository[E]{
		db:     db,
		schema: NewSchema(se),
	}
}

type Repository[E any] struct {
	db     *sqlx.DB
	schema Schema
}

func (r *Repository[E]) Schema() Schema {
	return r.schema
}

func (r *Repository[E]) Db() *sqlx.DB {
	return r.db
}
