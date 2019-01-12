package orm

import "database/sql"

type MysqlBuilder struct {
	BuilderBase
}

func NewMysqlBuilder(conn *sql.DB, model interface{}) *MysqlBuilder {
	return &MysqlBuilder{BuilderBase: BuilderBase{conn: conn, model: model}}
}

func (builder *MysqlBuilder) Get() []interface{} {
	sql := ""
	data := make([]interface{}, 0)

}
