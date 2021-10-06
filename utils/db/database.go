package database

// Database interface of database.
type Database interface {
	Connect() error
	Disconnect() error
	GetConnection() interface{}
}
