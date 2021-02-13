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
	ConnectTo(connectionId connections.ConnectionId) error
	DisconnectFrom(connectionId connections.ConnectionId) error

	GetDatabases(connectionId connections.ConnectionId) (*[]Database, error)
	FlushAllDatabases(connectionId connections.ConnectionId) error
	FlushDatabaseFor(connectionId connections.ConnectionId, id DatabaseId) error

	Save(connectionId connections.ConnectionId, object Object) error
	GetObjectsFor(connectionId connections.ConnectionId, id DatabaseId, page int, pageSize int) (*[]Object, error)
	DeleteObjectFor(connectionId connections.ConnectionId, id ObjetsId) (*Object, error)
}
