package connect

import (
	"database/sql"
	"fmt"
)

/*
MysqlConnect struct
*/
type MysqlConnect struct {
	BaseConnect
	host     string
	port     int32
	user     string
	passwd   string
	database string
}

/*
Config function set the parameters that needed to connect to mysql

args[0]:
	ip addr that connect to
args[1]:
	port that connection to
args[2]:
	user use to connect
args[3]:
	passwd use to connect
args[4]:
	databse should use
*/
func (connect *MysqlConnect) Config(args ...interface{}) {
	connect.host = args[0].(string)
	connect.port = args[1].(int32)
	connect.user = args[2].(string)
	connect.passwd = args[3].(string)
	connect.database = args[4].(string)
}

// Connect to database
func (connect *MysqlConnect) Connect() error {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", connect.user, connect.passwd, connect.host, connect.port, connect.database)
	connection, err := sql.Open("mysql", dsn)

	if err != nil {
		return err
	}
	connect.conn = connection
	return nil
}

// Reconnect to mysql database
func (connect *MysqlConnect) Reconnect() error {
	connect.Disconnect()
	return connect.Connect()
}

// Disconnect to databse
func (connect *MysqlConnect) Disconnect() {
	if connect.conn != nil {
		connect.conn.Close()
		connect.conn = nil
	}
}
