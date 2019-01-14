package orm

type BaseModel struct {
	Model interface{}
	Conn  string
	Table string
}
