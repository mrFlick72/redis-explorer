package connections

type Repository interface {
	GetConnectionFor(name ConnectionName) (*Connection, error)
	GetConnections() (*[]Connection, error)
	StoreConnection(connection *Connection) error
}

type ConnectionName string
type Host string
type Port int16
type Username string
type Password string

type Connection struct {
	Name     ConnectionName `json:"ConnectionName"`
	Host     Host           `json:",omitempty"`
	Port     Port
	Username Username `json:",omitempty"`
	Password Password `json:",omitempty"`
}
