package geeorm

import (
	"database/sql"
	"geeorm/log"
	"geeorm/session"
)

type Engine struct {
	db *sql.DB
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
	e = &Engine{db: db}
	log.Info("connect db success")
	return e, nil
}

func (e *Engine) Close() {
	if err := e.db.Close(); err != nil {
		log.Error("Failed to close DB connection")
	}
	log.Info("close db connection success")
}

func (e *Engine) NewSession() *session.Session {
	return session.New(e.db)
}
