package schema

import (
	"github.com/aws/aws-sdk-go/service/s3"
)

func NewRepository[E any](db *s3.S3, se E) Repository[E] {
	return Repository[E]{
		db:     db,
		schema: NewSchema(se),
	}
}

type Repository[E any] struct {
	db     *s3.S3
	schema Schema
}

func (r *Repository[E]) Schema() Schema {
	return r.schema
}

func (r *Repository[E]) Db() *s3.S3 {
	return r.db
}
