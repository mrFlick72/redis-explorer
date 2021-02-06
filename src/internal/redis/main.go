package redis

type DatabaseId int

type Database struct {
	Addr     string
	Username string
	Password string
	Db       int
}

type Page struct {
	Object   *[]Object
	Page     int
	PageSize int
}

type Object struct {
	content map[string]string
	Id      string
}

func (object *Object) ValueFor(key string) (string, error) {
	return object.content[key], nil
}

type ObjetsId string

type Repository interface {
	GetDatabases() (*[]Database, error)
	FlushAllDatabases() error
	FlushDatabaseFor(id DatabaseId) error

	Save(object Object) error
	GetObjectsFor(id DatabaseId, page int, pageSize int) (*[]Object, error)
	DeleteObjectFor(id ObjetsId) (*Object, error)
}
