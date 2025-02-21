package db

type MongoModule struct{}

func (m *MongoModule) ProvidedMongoDB() (Database, error) {
	return NewDevDatabase(), nil
}
