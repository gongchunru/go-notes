package dialect

import (
	"reflect"
)

type sqlite3 struct{}

func (s sqlite3) DataTypeOf(typ reflect.Value) string {
	//TODO implement me
	panic("implement me")
}

func (s sqlite3) TableExistSQL(tableName string) (string, []interface{}) {
	//TODO implement me
	panic("implement me")
}

var _ Dialect = (*sqlite3)(nil)

func init() {
	RegisterDialect("sqlite3", &sqlite3{})
}
