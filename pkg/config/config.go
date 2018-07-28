package config

import (
	"net/url"
	"time"
)

// ApplicationConfig holds all configurations of the application.
type ApplicationConfig struct {
	Union struct {
		DBConfig         `json:"db"         xml:"db"         yaml:"db"`
		GRPCConfig       `json:"grpc"       xml:"grpc"       yaml:"grpc"`
		MigrationConfig  `json:"migration"  xml:"migration"  yaml:"migration"`
		MonitoringConfig `json:"monitoring" xml:"monitoring" yaml:"monitoring"`
		ProfilerConfig   `json:"profiler"   xml:"profiler"   yaml:"profiler"`
		ServerConfig     `json:"server"     xml:"server"     yaml:"server"`
	} `json:"config" xml:"config" yaml:"config"`
}

// DBConfig contains configuration related to database.
type DBConfig struct {
	DSN Secret `json:"dsn" xml:"dsn" yaml:"dsn"`

	dsn *url.URL
}

// DriverName returns database driver name.
func (cnf *DBConfig) DriverName() string {
	if cnf.dsn == nil {
		cnf.dsn, _ = url.Parse(string(cnf.DSN))
	}
	return cnf.dsn.Scheme
}

// GRPCConfig contains configuration related to gRPC server.
type GRPCConfig struct {
	Interface string        `json:"interface" xml:"interface" yaml:"interface"`
	Timeout   time.Duration `json:"timeout"   xml:"timeout"   yaml:"timeout"`
	Token     Secret        `json:"token"     xml:"token"     yaml:"token"`
}

// MigrationConfig contains configuration related to migrations.
type MigrationConfig struct {
	Table    string `json:"table"     xml:"table"     yaml:"table"`
	Schema   string `json:"schema"    xml:"schema"    yaml:"schema"`
	Limit    uint   `json:"limit"     xml:"limit"     yaml:"limit"`
	DryRun   bool   `json:"dry-run"   xml:"dry-run"   yaml:"dry-run"`
	WithDemo bool   `json:"with-demo" xml:"with-demo" yaml:"with-demo"`
}

// MonitoringConfig contains configuration related to monitoring.
type MonitoringConfig struct {
	Enabled   bool   `json:"enabled"   xml:"enabled"   yaml:"enabled"`
	Interface string `json:"interface" xml:"interface" yaml:"interface"`
}

// ProfilerConfig contains configuration related to profiler.
type ProfilerConfig struct {
	Enabled   bool   `json:"enabled"   xml:"enabled"   yaml:"enabled"`
	Interface string `json:"interface" xml:"interface" yaml:"interface"`
}

// ServerConfig contains configuration related to the server.
type ServerConfig struct {
	Interface         string        `json:"interface"           xml:"interface"           yaml:"interface"`
	CPUCount          uint          `json:"cpus"                xml:"cpus"                yaml:"cpus"`
	ReadTimeout       time.Duration `json:"read-timeout"        xml:"read-timeout"        yaml:"read-timeout"`
	ReadHeaderTimeout time.Duration `json:"read-header-timeout" xml:"read-header-timeout" yaml:"read-header-timeout"`
	WriteTimeout      time.Duration `json:"write-timeout"       xml:"write-timeout"       yaml:"write-timeout"`
	IdleTimeout       time.Duration `json:"idle-timeout"        xml:"idle-timeout"        yaml:"idle-timeout"`
}