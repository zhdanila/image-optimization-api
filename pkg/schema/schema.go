package schema

import (
	"fmt"
	"go.uber.org/zap"
	"reflect"
	"strings"
)

type DataHolder interface {
	Data() []any
}

type Columns interface {
	Columns(except ...string) []string
	SelectColumns() string
}

type Schema interface {
	Columns
	Data(DataHolder) ([]any, error)
	PkColumn1() string
	TableName() string
	ValidateColumns(data map[string]any) error
}

type tableSchema struct {
	entity        any
	escapeColumns bool

	pkColumns    []string
	columns      []string
	validColumns map[string]struct{}
	tableName    string
}

func NewSchema(entity any) Schema {
	res := &tableSchema{
		entity: entity,
	}

	res.columns, res.tableName, res.pkColumns = mustBuild(entity)
	res.validColumns = make(map[string]struct{})
	for _, col := range res.columns {
		res.validColumns[col] = struct{}{}
	}

	return res
}

func (s *tableSchema) Columns(except ...string) []string {
	if except == nil {
		return s.columns
	}

	fs := make([]string, 0, len(s.columns)-len(except))
Fields:
	for i := range s.columns {
		for j := range except {
			if strings.Trim(s.columns[i], `"'`) == strings.Trim(except[j], `"'`) {
				continue Fields
			}
		}
		fs = append(fs, s.columns[i])
	}

	return fs
}

func (s *tableSchema) ColumnsWithComma(except ...string) string {
	columns := s.Columns(except...)

	return strings.Join(columns, ", ")
}

func (s *tableSchema) PkColumn1() string {
	return s.pkColumns[0]
}

func (s *tableSchema) SelectColumns() string {
	return s.ColumnsWithComma()
}

func (s *tableSchema) TableName() string {
	return s.tableName
}

func (s *tableSchema) Data(ent DataHolder) ([]any, error) {
	data := ent.Data()
	if len(s.Columns()) != len(data) {
		return nil, fmt.Errorf("mismatch data and columns length")
	}

	return data, nil
}

// ValidateData validates the keys of the given map against the schema columns.
func (s *tableSchema) ValidateColumns(data map[string]any) error {
	for key := range data {
		if _, exists := s.validColumns[key]; !exists {
			return fmt.Errorf("invalid key: %s", key)
		}
	}

	return nil
}

func mustBuild(entity any) (columns []string, tableName string, pkColumns []string) {
	v := reflect.TypeOf(entity)
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}

	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i)

		if dbTag, ok := field.Tag.Lookup("db"); ok {
			columns = append(columns, dbTag)
		}

		if i == 0 {
			if tableTag, ok := field.Tag.Lookup("table_name"); ok {
				tableName = tableTag
			}
		}

		if _, ok := field.Tag.Lookup("pk_number"); ok {
			pkColumns = append(pkColumns, field.Tag.Get("db"))
		}
	}

	if tableName != `` && pkColumns == nil {
		zap.L().Error(fmt.Sprintf("pk_number tag not found for table %s", tableName))
	}

	return columns, tableName, pkColumns
}
