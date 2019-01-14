package orm

import (
	"database/sql"
	"fmt"
	"reflect"
	"regexp"
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
	Get() ([]interface{}, error)
	First() (interface{}, error)
	Insert(interface{}) bool
	Update(interface{}) bool
	Delete() bool
}

type BuilderBase struct {
	andWhere  []whereStruct
	orWhere   []whereStruct
	field     []string
	group     []string
	order     []orderStruct
	offset    int
	limit     int
	model     BaseModel
	fmap      map[string]string
	conn      *sql.DB
	queryData []interface{}
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
	elem := reflect.TypeOf(builder.model.Model).Elem()
	reg := regexp.MustCompile(`^[A-Z]`)
	for index := 0; index < elem.NumField(); index++ {
		field := elem.Field(index)
		name := field.Name
		tag := string(field.Tag)

		// 仅公共变量做记录
		if reg.MatchString(name) {
			builder.fmap[tag] = field.Name
		}
	}
}

func (builder *BuilderBase) newModeObj() interface{} {
	modelType := reflect.TypeOf(builder.model.Model)
	return reflect.New(modelType)
}
