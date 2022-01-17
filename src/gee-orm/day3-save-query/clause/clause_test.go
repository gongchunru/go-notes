package clause

import (
	"reflect"
	"testing"
)

func TestSelect(t *testing.T) {
	var clause Clause
	clause.Set(LIMIT, 3)
	clause.Set(SELECT, "User", []string{"*"})
	clause.Set(WHERE, "Name = ? and Age = ?", "Tom", "20")
	clause.Set(ORDERBY, "Age ASC")
	sql, vars := clause.Build(SELECT, WHERE, ORDERBY, LIMIT)

	t.Log(sql, vars)
	if sql != "SELECT * FROM User WHERE Name = ? and Age = ? ORDER BY Age ASC LIMIT ?" {
		t.Fatal("Failed to build SQL")
	}
	if !reflect.DeepEqual(vars, []interface{}{"Tom", "20", 3}) {
		t.Fatal("Failed to build SQL Vars")
	}
}

func TestClause_Build(t *testing.T) {
	t.Run("select", func(t *testing.T) {
		TestSelect(t)
	})
}
