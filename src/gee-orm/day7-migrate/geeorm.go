package geeorm

import (
	"database/sql"
	"fmt"
	"geeorm/dialect"
	"geeorm/log"
	"geeorm/session"
	"strings"
)

type Engine struct {
	db      *sql.DB
	dialect dialect.Dialect
}

type TxFunc func(*session.Session) (interface{}, error)

func (engine *Engine) Transaction(f TxFunc) (result interface{}, err error) {
	s := engine.NewSession()
	if err := s.Begin(); err != nil {
		return nil, err
	}
	defer func() {
		if p := recover(); p != nil {
			_ = s.RollBack()
			panic(p)
		} else if err != nil {
			_ = s.RollBack()
		} else {
			err = s.Commit()
		}
	}()
	return f(s)
}

func NewEngine(driver, source string) (e *Engine, err error) {
	db, err := sql.Open(driver, source)
	if err != nil {
		log.Error(err)
		return
	}
	if err = db.Ping(); err != nil {
		log.Error(err)
		return
	}
	// make sure the specific dialect exists
	dial, ok := dialect.GetDialect(driver)
	if !ok {
		log.Errorf("dialect %s Not Found", dial)
	}
	e = &Engine{db: db, dialect: dial}
	log.Info("connect db success")
	return
}

func (e *Engine) Close() {
	if err := e.db.Close(); err != nil {
		log.Error("Failed to close DB connection")
	}
	log.Info("close db connection success")
}

func (e *Engine) NewSession() *session.Session {
	return session.New(e.db, e.dialect)
}

//   输出 [a - b]的内容
func difference(a []string, b []string) (diff []string) {
	mapB := make(map[string]bool)
	for _, v := range b {
		mapB[v] = true
	}

	for _, v := range a {
		if _, ok := mapB[v]; !ok {
			diff = append(diff, v)
		}
	}
	return
}

func (engine *Engine) Migrate(value interface{}) error {
	_, err := engine.Transaction(func(s *session.Session) (result interface{}, err error) {
		if !s.Model(value).HasTable() {
			log.Infof("table %s doesn't exist", s.RefTable().Name)
			return nil, s.CreateTable()
		}

		table := s.RefTable()
		// 旧表
		rows, _ := s.Raw(fmt.Sprintf("SELECT * FROM %s LIMIT 1", table.Name)).QueryRows()
		columns, _ := rows.Columns()
		// 新表 - 旧表 新增的字段
		addCols := difference(table.FieldNames, columns)
		delCols := difference(columns, table.FieldNames)

		log.Infof("added cols %v,deleted cols %v", addCols, delCols)

		for _, col := range addCols {
			f := table.GetField(col)
			sqlStr := fmt.Sprintf("ALTER TABLE %s ADD COLUMN %s %s", table.Name, f.Name, f.Type)
			if _, err = s.Raw(sqlStr).Exec(); err != nil {
				log.Error(err)
				return
			}
		}
		if len(delCols) == 0 {
			return
		}

		tmp := "tmp_" + table.Name
		fieldStr := strings.Join(table.FieldNames, ",")
		s.Raw(fmt.Sprintf("CREATE TABLE %s AS SELECT %s from %s;", tmp, fieldStr, table.Name))
		s.Raw(fmt.Sprintf("DROP TABLE %s;", table.Name))
		s.Raw(fmt.Sprintf("ALTER TABLE %s RENAME TO %s", tmp, table.Name))
		_, err = s.Exec()
		return
	})
	return err
}
