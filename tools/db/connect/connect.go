package connect

import (
	"database/sql"
	"errors"
)

/*
ConnectInterface interface of database connect
*/
type ConnectInterface interface {
	GetName() string
	Config(args ...interface{})
	Connect() error
	Reconnect() error
	Disconnect()
}

/*
BaseConnect struct
*/
type BaseConnect struct {
	name string
	conn *sql.DB
}

// GetName get the name of connect
func (connect *BaseConnect) GetName() string {
	return connect.name
}

// GetConn get the conn
func (connect *BaseConnect) GetConn() *sql.DB {
	return connect.conn
}

/*
NewConnect return connect with specified driver
*/
func NewConnect(name string, driver string, args ...interface{}) (ConnectInterface, error) {
	var err error
	var connect ConnectInterface
	err = nil
	switch driver {
	case "mysql":
		connect := &MysqlConnect{BaseConnect: BaseConnect{name: name}}
		connect.Config(args)
	default:
		connect = nil
		err = errors.New("Can't find driver")
	}
	return connect, err
}
