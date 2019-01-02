package orm

import "fmt"

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
	BuildSql() (string, []interface{})
}

type BuilderBase struct {
	andWhere []whereStruct
	orWhere  []whereStruct
	field    []string
	group    []string
	order    []orderStruct
	offset   int
	limit    int
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

func (builder *BuilderBase) Get() {

}

func (builder *BuilderBase) Insert() {

}

func (builder *BuilderBase) Update() {

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
