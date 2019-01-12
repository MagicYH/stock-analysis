package orm

import (
	"database/sql"
	"fmt"
	"reflect"
)

type whereStruct struct {
	Key   string
	Op    string
	Value interface{}
}

type orderStruct struct {
	Key string
	Ord string
}

type QueryBuilder interface {
	Get() []interface{}
	First() interface{}
	Insert(interface{}) bool
	Update(interface{}) bool
	Delete() bool
}

type BuilderBase struct {
	andWhere []whereStruct
	orWhere  []whereStruct
	field    []string
	group    []string
	order    []orderStruct
	offset   int
	limit    int
	model    interface{}
	fmap     map[string]string
	conn     *sql.DB
}

func NewBuilder(driver string) (QueryBuilder, error) {
	var builder QueryBuilder
	switch driver {
	case "mysql":
		builder = &MysqlBuilder{}
	default:
		return nil, fmt.Errorf("Undefine builder type: %s", driver)
	}
	return builder, nil
}

func (builder *BuilderBase) Where(key string, op string, value interface{}) *BuilderBase {
	andWhere := whereStruct{key, op, value}
	builder.andWhere = append(builder.andWhere, andWhere)
	return builder
}

func (builder *BuilderBase) OrWhere(key string, op string, value interface{}) *BuilderBase {
	orWhere := whereStruct{key, op, value}
	builder.orWhere = append(builder.orWhere, orWhere)
	return builder
}

func (builder *BuilderBase) parseModel() {
	builder.fmap = make(map[string]string, 0)
	elem := reflect.TypeOf(builder.model).Elem()
	for index := 0; index < elem.NumField(); index++ {
		field := elem.Field(index)
		builder.fmap[field.Name] = string(field.Tag)
	}
}
