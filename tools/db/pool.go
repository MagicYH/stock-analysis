package db

import (
	"database/sql"
	"fmt"

	"github.com/MagicYH/stock-analysis/tools"
)

var _connectPool map[string]*sql.DB

func init() {
	_connectPool = make(map[string]*sql.DB)
}

// InitConnection initialize the connection pools
func InitConnection(config tools.ConfigDb) {
	if len(config.Mysql) > 0 {
		for _, conf := range config.Mysql {
			initMysqlConn(conf.Name, conf.Host, conf.Port, conf.User, conf.Passwd, conf.Database)
		}
	}
	// Add other connection type here
}

// GetConnection get the *sql.DB object by connection name
func GetConnection(name string) (*sql.DB, error) {
	conn, ok := _connectPool[name]
	if !ok {
		return nil, fmt.Errorf("Connection %s not found", name)
	}
	return conn, nil
}

func initMysqlConn(name string, host string, port int, user string, passwd string, database string) {
	conn, ok := _connectPool[name]
	if ok {
		panic("Duplicate connection name: " + name)
	}

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", user, passwd, host, port, database)
	conn, err := sql.Open("mysql", dsn)
	if err != nil {
		panic(fmt.Sprintf("Open database connection %s faile", name))
	}
	_connectPool[name] = conn
}
