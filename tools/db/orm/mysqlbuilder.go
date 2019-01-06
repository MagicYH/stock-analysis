package orm

import "database/sql"

type MysqlBuilder struct {
	BuilderBase
}

func NewMysqlBuilder(conn *sql.DB, mode interface{}) *MysqlBuilder {
	return &MysqlBuilder{BuilderBase: BuilderBase{conn: conn, mode: mode}}
}

func (builder *MysqlBuilder) Get() []interface{} {
	sql := ""
	data := make([]interface{}, 0)

}
