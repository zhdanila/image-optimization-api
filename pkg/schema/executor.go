package schema

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/jmoiron/sqlx"
)

type SqSQLer interface {
	ToSql() (string, []interface{}, error)
}

type Entity interface {
	ToData() []any
}

func NewExecutor[E any](db *sqlx.DB, se E) Executor[E] {
	return Executor[E]{
		db:     db,
		schema: NewSchema(se),
		defaultIterFn: func(iter *sqlx.Rows) ([]E, error) {
			res := make([]E, 0)
			for iter.Next() {
				it := new(E)
				if err := iter.StructScan(it); err != nil {
					return nil, err
				}
				res = append(res, *it)
			}
			if err := iter.Err(); err != nil {
				return nil, err
			}
			return res, nil
		},
	}
}

type Executor[E any] struct {
	db            *sqlx.DB
	schema        Schema
	defaultIterFn func(*sqlx.Rows) ([]E, error)
}

func (r *Executor[E]) Iterate(ctx context.Context, query string, args ...interface{}) ([]E, error) {
	rows, err := r.db.QueryxContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return r.defaultIterFn(rows)
}

func (r *Executor[E]) IterateAfterSq(ctx context.Context, sb SqSQLer, addArgs ...any) ([]E, error) {
	sql, args, err := sb.ToSql()
	if err != nil {
		return nil, err
	}

	params := make(map[string]any)
	for i, arg := range args {
		i++
		params[fmt.Sprintf(`p%d`, i)] = arg
	}
	if addArgs != nil {
		if len(addArgs)%2 != 0 {
			return nil, fmt.Errorf(`additional arguments must be even`)
		}
		for i := 0; i < len(addArgs); i += 2 {
			key, ok := addArgs[i].(string)
			if !ok {
				return nil, fmt.Errorf(`additional argument key must be a string`)
			}
			params[key] = addArgs[i+1]
		}
	}

	stmt, args, err := sqlx.Named(sql, params)
	if err != nil {
		return nil, err
	}

	return r.Iterate(ctx, stmt, args...)
}

func (r *Executor[E]) RowAfterSq(ctx context.Context, sb SqSQLer) (*E, error) {
	sqler, args, err := sb.ToSql()
	if err != nil {
		return nil, err
	}

	rows, err := r.db.QueryxContext(ctx, sqler, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	if rows.Next() {
		var entity E
		if err := rows.StructScan(&entity); err != nil {
			return nil, err
		}
		return &entity, nil
	}

	return nil, sql.ErrNoRows
}

func (r *Executor[E]) Row(ctx context.Context, stmt string, args ...interface{}) (*E, error) {
	ents, err := r.Iterate(ctx, stmt, args...)
	if err != nil {
		return nil, err
	}

	if len(ents) == 0 {
		return nil, sql.ErrNoRows
	}

	return &ents[0], nil
}

func IterateSpecial[T Entity](ctx context.Context, db *sqlx.DB, query string, args ...interface{}) ([]T, error) {
	rows, err := db.QueryxContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var res []T
	for rows.Next() {
		var inst T
		if err := rows.StructScan(&inst); err != nil {
			return nil, err
		}
		res = append(res, inst)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return res, nil
}
