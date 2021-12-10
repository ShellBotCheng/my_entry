package mysql

import (
	"database/sql"
	"github.com/pkg/errors"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

var (
	Db *sql.DB
)

// Config mysql config.
type Config struct {
	DSN             string // write data source name.
	MaxOpenConn     int    // open pool
	MaxIdleConn     int    // idle pool
	ConnMaxLifeTime int
}

// Init 初始化数据库
func Init(cfg *Config) *sql.DB {
	Db = NewMySQL(cfg)
	return Db
}

// GetDB 返回默认的数据库
func GetDB() *sql.DB {
	return Db
}

// NewMySQL new db and retry connection when has error.
func NewMySQL(c *Config) (db *sql.DB) {

	db, err := connect(c, c.DSN)
	if err != nil {
		panic(err)
	}
	return
}

func connect(c *Config, dataSourceName string) (*sql.DB, error) {
	d, err := sql.Open("mysql", dataSourceName)
	if err != nil {
		err = errors.WithStack(err)
		return nil, err
	}
	d.SetMaxOpenConns(c.MaxOpenConn)
	d.SetMaxIdleConns(c.MaxIdleConn)
	d.SetConnMaxLifetime(time.Duration(c.ConnMaxLifeTime))
	return d, nil
}
