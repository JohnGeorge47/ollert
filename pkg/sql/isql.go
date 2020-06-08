package sql

type Isql interface {
	Insert(query string, args ...interface{}) (*int64, error)
}
