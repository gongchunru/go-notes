package schema

import (
	"geeorm/dialect"
	"go/ast"
	"reflect"
)

type Field struct {
	Name string
	Type string
	// 约束条件 Tag
	Tag string
}

type Schema struct {
	Model  interface{}
	Name   string
	Fields []*Field
	// FieldNames 包含所有的字段名(列名)
	FieldNames []string
	fieldMap   map[string]*Field
}

func (schema *Schema) GetField(name string) *Field {
	return schema.fieldMap[name]
}

//Parse 函数，将任意的对象解析为 Schema 实例。
func Parse(dest interface{}, d dialect.Dialect) *Schema {
	// reflect.Indirect() 获取指针指向的实例。
	modelType := reflect.Indirect(reflect.ValueOf(dest)).Type()

	schema := &Schema{
		Model: dest,
		// 表名
		Name:     modelType.Name(),
		fieldMap: make(map[string]*Field),
	}

	for i := 0; i < modelType.NumField(); i++ {
		p := modelType.Field(i)
		if !p.Anonymous && ast.IsExported(p.Name) {
			field := &Field{
				Name: p.Name,
				Type: d.DataTypeOf(reflect.Indirect(reflect.New(p.Type))),
			}
			if v, ok := p.Tag.Lookup("geeorm"); ok {
				field.Tag = v
			}
			schema.Fields = append(schema.Fields, field)
			schema.FieldNames = append(schema.FieldNames, p.Name)
			schema.fieldMap[p.Name] = field
		}
	}
	return schema
}

// RecordValues  从对象中找到对应的值，按顺序平铺为 sql 的插入语句
// 比如: ("Tom", 18), ("Same", 25)
//func (schema *Schema) RecordValues(dest interface{}) []interface{} {
//	destValue := reflect.Indirect(reflect.ValueOf(dest))
//	var fieldValues []interface{}
//	for _, field := range schema.Fields {
//		fieldValues = append(fieldValues, destValue.FieldByName(field.Name).Interface())
//	}
//	return fieldValues
//}
// Values return the values of dest's member variables
func (schema *Schema) RecordValues(dest interface{}) []interface{} {
	destValue := reflect.Indirect(reflect.ValueOf(dest))
	var fieldValues []interface{}
	for _, field := range schema.Fields {
		fieldValues = append(fieldValues, destValue.FieldByName(field.Name).Interface())
	}
	return fieldValues
}
