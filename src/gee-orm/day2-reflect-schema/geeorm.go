package geeorm

import (
	"database/sql"
	"geeorm/dialect"
	"geeorm/log"
	"geeorm/session"
)

type Engine struct {
	db      *sql.DB
	dialect dialect.Dialect
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
