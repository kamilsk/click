package dao

import (
	"database/sql"

	"github.com/kamilsk/click/dao/postgres"
	"github.com/kamilsk/click/domain"
)

// Must returns a new instance of the Storage or panics if it cannot configure it.
func Must(configs ...Configurator) *Storage {
	instance, err := New(configs...)
	if err != nil {
		panic(err)
	}
	return instance
}

// New returns a new instance of the Storage or an error if it cannot configure it.
func New(configs ...Configurator) (*Storage, error) {
	instance := &Storage{}
	for _, configure := range configs {
		if err := configure(instance); err != nil {
			return nil, err
		}
	}
	return instance, nil
}

// Connection returns database connection Configurator.
func Connection(driver, dsn string, open, idle int) Configurator {
	return func(instance *Storage) error {
		var err error
		instance.conn, err = sql.Open(driver, dsn)
		instance.conn.SetMaxOpenConns(open)
		instance.conn.SetMaxIdleConns(idle)
		return err
	}
}

// Configurator defines a function which can use to configure the Storage.
type Configurator func(*Storage) error

// Storage is an implementation of Data Access Object.
type Storage struct {
	conn *sql.DB
}

// Connection returns current database connection.
func (l *Storage) Connection() *sql.DB {
	return l.conn
}

// Dialect returns supported database dialect.
func (l *Storage) Dialect() string {
	return postgres.Dialect()
}

// Link returns the Link with its Aliases and Targets by provided ID.
func (l *Storage) Link(id domain.UUID) (domain.Link, error) {
	return postgres.Link(l.conn, id)
}

// LinkByAlias returns the Link with its Targets and the single Alias defined by Namespace and URN.
func (l *Storage) LinkByAlias(alias domain.Alias) (domain.Link, error) {
	return postgres.LinkByAlias(l.conn, alias)
}
