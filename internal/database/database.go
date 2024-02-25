package database

type DatabaseService interface {
	Connect() error
	Disconnect() error
}
