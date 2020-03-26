package database

import (
	"context"
	"fmt"
	"github.com/gocraft/dbr/v2"
)

type Database struct {
	Conn *dbr.Connection
}

func NewDataBase(opt *DatabaseOptions, stopCh <-chan struct{}) (*Database, error) {
	url := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", opt.Username, opt.Password, opt.Host, opt.Port, opt.Database)
	conn, err := dbr.Open("mysql", url+"?parseTime=1&multiStatements=1&charset=utf8mb4&collation=utf8mb4_unicode_ci", nil)
	if err != nil {
		return nil, err
	}
	conn.SetMaxIdleConns(opt.MaxIdleConnections)
	conn.SetMaxOpenConns(opt.MaxOpenConnections)
	conn.SetConnMaxLifetime(opt.MaxConnectionLifeTime)

	err = conn.Ping()
	if err != nil {
		return nil, err
	}
	db := &Database{
		Conn: conn,
	}

	go func() {
		<-stopCh
		if db.Conn != nil {
			db.Conn.Close()
		}
	}()

	return db, nil
}

func (db *Database) NewConn(ctx context.Context) *Conn {
	var conn *dbr.Connection
	if db == nil {
		return nil
	} else {
		conn = db.Conn
	}
	return &Conn{
		Session: conn.NewSession(&EventReceiver{ctx}),
		ctx:     ctx,
	}

}
