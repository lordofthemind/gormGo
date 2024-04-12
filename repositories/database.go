// repositories/database.go
package repositories

// Database represents the database connection.
type Database interface {
	Create(interface{}) error
	First(interface{}, ...interface{}) error
	Find(interface{}, ...interface{}) error
	Save(interface{}) error
	Delete(interface{}, ...interface{}) error
}
