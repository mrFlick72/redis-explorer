package connections

type Repository struct {
	Repo RepositoryOperations
}

type RepositoryOperations interface {
	GetConnectionFor(id ConnectionId) (*Connection, error)
	GetConnections() (*[]Connection, error)
	StoreConnection(connection *Connection) error
}

type ConnectionId = string
type ConnectionName = string
type HostAndPort = string
type Username = string
type Password = string

type Connection struct {
	Id          ConnectionId   `json:"ConnectionId"`
	Name        ConnectionName `json:"ConnectionName"`
	HostAndPort HostAndPort    `json:",omitempty"`
	Username    Username       `json:",omitempty"`
	Password    Password       `json:",omitempty"`
}
