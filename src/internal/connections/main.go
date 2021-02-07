package connections

type Repository interface {
	GetConnectionFor(name ConnectionName) (*Connection, error)
	GetConnections() (*[]Connection, error)
	StoreConnection(connection *Connection) error
}

type ConnectionName string
type Username string
type Password string

type Connection struct {
	Name     ConnectionName
	Username Username
	Password Password
}
