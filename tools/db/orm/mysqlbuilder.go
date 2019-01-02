package orm

type MysqlBuilder struct {
	BuilderBase
}

func (builder *MysqlBuilder) BuildSql() (string, []interface{}) {
	sql := ""
	data := make([]interface{}, 0)
	return sql, data
}
