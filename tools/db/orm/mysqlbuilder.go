package orm

import (
	"database/sql"
	"fmt"
	"reflect"
	"strings"
)

type MysqlBuilder struct {
	BuilderBase
}

func NewMysqlBuilder(conn *sql.DB, model BaseModel) *MysqlBuilder {
	return &MysqlBuilder{BuilderBase: BuilderBase{conn: conn, model: model, queryData: make([]interface{}, 0)}}
}

func (builder *MysqlBuilder) Get() ([]interface{}, error) {
	strField := builder.getStrField()
	strWhere := builder.getStrWhere()
	sql := fmt.Sprintf("SELECT %s FROM %s WHERE %s", strField, builder.model.Table, strWhere)

	modelType := reflect.TypeOf(builder.model.Model)
	rows, err := builder.conn.Query(sql, builder.queryData...)
	defer rows.Close()
	if err != nil {
		return make([]interface{}, 0), err
	}

	fields, err := rows.Columns()
	if err != nil {
		return make([]interface{}, 0), err
	}

	buffer := make([]interface{}, len(fields))
	output := make([]interface{}, 0)
	for rows.Next() {
		obj := reflect.New(modelType)
		for i := 0; i < len(fields); i++ {
			fieldName := builder.fmap[fields[i]]
			objField := obj.Elem().FieldByName(fieldName)
			buffer[i] = objField.Pointer()
		}
		rows.Scan(buffer...)
		output = append(output, obj)
	}
	return output, nil
}

func (builder *MysqlBuilder) First() (interface{}, error) {
	return nil, nil
}

func (builder *MysqlBuilder) Insert(interface{}) bool {
	return true
}

func (builder *MysqlBuilder) Update(interface{}) bool {
	return true
}

func (builder *MysqlBuilder) Delete() bool {
	return true
}

func (builder *MysqlBuilder) getStrField() string {
	newFields := make([]string, len(builder.field))
	for i := 0; i < len(builder.field); i++ {
		newFields[i] = fmt.Sprintf("`%s`", builder.field[i])
	}
	return strings.Join(newFields, ",")
}

func (builder *MysqlBuilder) getStrWhere() string {
	andStrWhere := builder.getStrAndWhere()
	orStrWhere := builder.getStrOrWhere()
	var strWhere string
	if andStrWhere != "" && orStrWhere != "" {
		strWhere = fmt.Sprintf("%s OR %s", andStrWhere, orStrWhere)
	} else if andStrWhere != "" {
		strWhere = andStrWhere
	} else {
		strWhere = orStrWhere
	}
	return strWhere
}

func (builder *MysqlBuilder) getStrAndWhere() string {
	strWhere := ""
	whereArray := make([]string, len(builder.andWhere))
	if len(whereArray) == 0 {
		return strWhere
	}
	for i := 0; i < len(builder.andWhere); i++ {
		oneWhere := builder.andWhere[i]
		switch oneWhere.Op {

		default:
			whereArray[i] = fmt.Sprintf("`%s` %s ?", oneWhere.Key, oneWhere.Op)
			builder.queryData = append(builder.queryData, oneWhere.Value)
		}
	}
	strWhere = strings.Join(whereArray, " AND ")
	return strWhere
}

func (builder *MysqlBuilder) getStrOrWhere() string {
	strWhere := ""
	whereArray := make([]string, len(builder.orWhere))
	if len(whereArray) == 0 {
		return strWhere
	}
	for i := 0; i < len(builder.orWhere); i++ {
		oneWhere := builder.andWhere[i]
		switch oneWhere.Op {

		default:
			whereArray[i] = fmt.Sprintf("`%s` %s '%s'", oneWhere.Key, oneWhere.Op, oneWhere.Value)
		}
	}
	strWhere = strings.Join(whereArray, " OR ")
	return strWhere
}
