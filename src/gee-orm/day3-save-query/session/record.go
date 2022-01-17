package session

import (
	"geeorm/clause"
	"reflect"
)

// 记录增删改查相关的代码
//
//func (s *Session) Insert(values ...interface{}) (int64, error) {
//	recordValues := make([]interface{}, 0)
//	for _, value := range values {
//		table := s.Model(value).RefTable()
//		s.clause.Set(clause.INSERT, table.Name, table.FieldNames)
//		recordValues = append(recordValues, table.RecordValues(value))
//	}
//	s.clause.Set(clause.VALUES, recordValues...)
//	sql, vars := s.clause.Build(clause.INSERT, clause.VALUES)
//	result, err := s.Raw(sql, vars...).Exec()
//	if err != nil {
//		return 0, err
//	}
//	return result.RowsAffected()
//}

// Insert one or more records in database
func (s *Session) Insert(values ...interface{}) (int64, error) {
	recordValues := make([]interface{}, 0)
	for _, value := range values {
		table := s.Model(value).RefTable()
		s.clause.Set(clause.INSERT, table.Name, table.FieldNames)
		recordValues = append(recordValues, table.RecordValues(value))
	}

	s.clause.Set(clause.VALUES, recordValues...)
	sql, vars := s.clause.Build(clause.INSERT, clause.VALUES)
	result, err := s.Raw(sql, vars...).Exec()
	if err != nil {
		return 0, err
	}

	return result.RowsAffected()
}

func (s *Session) Find(values interface{}) error {
	destSlice := reflect.Indirect(reflect.ValueOf(values))
	// 获取切片的单个元素的类型 destType
	destType := destSlice.Type().Elem()
	table := s.Model(reflect.New(destType).Elem().Interface()).RefTable()

	s.clause.Set(clause.SELECT, table.Name, table.FieldNames)

	sql, vars := s.clause.Build(clause.SELECT, clause.WHERE, clause.ORDERBY, clause.LIMIT)
	rows, err := s.Raw(sql, vars...).QueryRows()
	if err != nil {
		return err
	}

	for rows.Next() {
		// 利用反射创建 destType 的实例 dest
		dest := reflect.New(destType).Elem()
		var values []interface{}
		for _, name := range table.FieldNames {
			// 将 dest 的所有字段平铺开，构造切片 values。
			values = append(values, dest.FieldByName(name).Addr().Interface())
		}
		if err := rows.Scan(values...); err != nil {
			return err
		}
		destSlice.Set(reflect.Append(destSlice, dest))
	}
	return rows.Close()
}
