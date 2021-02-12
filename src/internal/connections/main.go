package connections

type Repository struct {
	Repo2 RepositoryOperations
	Repo  RepositoryOperations
}

type RepositoryOperations interface {
	GetConnectionFor(name ConnectionName) (*Connection, error)
	GetConnections() (*[]Connection, error)
	StoreConnection(connection *Connection) error
}

type ConnectionName = string
type HostAndPort = string
type Username = string
type Password = string

type Connection struct {
	Name        ConnectionName `json:"ConnectionName"`
	HostAndPort HostAndPort    `json:",omitempty"`
	Username    Username       `json:",omitempty"`
	Password    Password       `json:",omitempty"`
}
