package database

type Database struct {
}

func New() (*Database, error) {
    return &Database{}, nil
}
