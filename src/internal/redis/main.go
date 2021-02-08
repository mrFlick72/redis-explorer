package redis

import "github.com/mrflick72/redis-explorer/src/internal/connections"

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
	ConnectTo(connectionName *connections.ConnectionName) error
	DisconnectFrom(connectionName *connections.ConnectionName) error

	GetDatabases(connectionName *connections.ConnectionName) (*[]Database, error)
	FlushAllDatabases(connectionName *connections.ConnectionName) error
	FlushDatabaseFor(connectionName *connections.ConnectionName, id DatabaseId) error

	Save(connectionName *connections.ConnectionName, object Object) error
	GetObjectsFor(connectionName *connections.ConnectionName, id DatabaseId, page int, pageSize int) (*[]Object, error)
	DeleteObjectFor(connectionName *connections.ConnectionName, id ObjetsId) (*Object, error)
}
