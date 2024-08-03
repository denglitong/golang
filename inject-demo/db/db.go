package db

type Database interface {
	Ping() (string, error)
}

type DevDatabase struct {
	URI string
}

func NewDevDatabase() Database {
	return &DevDatabase{URI: "localhost:27017"}
}

func (m *DevDatabase) Ping() (string, error) {
	return "devDb: success!", nil
}
