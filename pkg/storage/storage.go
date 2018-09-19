package storage

import (
	"database/sql"

	_ "github.com/lib/pq"

	"github.com/kamilsk/click/pkg/config"
	"github.com/kamilsk/click/pkg/domain"
	"github.com/kamilsk/click/pkg/storage/postgres"
	"github.com/pkg/errors"
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
		if err := errors.WithStack(configure(instance)); err != nil {
			return nil, err
		}
	}
	return instance, nil
}

// Connection returns database connection Configurator.
func Connection(cnf config.DBConfig) Configurator {
	return func(instance *Storage) error {
		var err error
		instance.conn, err = sql.Open(cnf.DriverName(), string(cnf.DSN))
		if err == nil {
			instance.conn.SetMaxOpenConns(cnf.MaxOpen)
			instance.conn.SetMaxIdleConns(cnf.MaxIdle)
			instance.conn.SetConnMaxLifetime(cnf.MaxLifetime)
		}
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
func (l *Storage) Link(id domain.ID) (domain.Link, error) {
	return postgres.Link(l.conn, id)
}

// LinkByAlias returns the Link with its set of Alias and set of Target defined by provided namespace and URN.
func (l *Storage) LinkByAlias(ns, urn string) (domain.Link, error) {
	return postgres.LinkByAlias(l.conn, ns, urn)
}

// Log stores a "redirect event".
func (l *Storage) Log(event domain.Log) (domain.Log, error) {
	return postgres.Log(l.conn, event)
}

// UUID returns a new generated unique identifier.
func (l *Storage) UUID() (domain.ID, error) {
	return postgres.UUID(l.conn)
}
