//go:generate echo $PWD/$GOPACKAGE/$GOFILE
//go:generate mockgen -package dao_test -destination $PWD/dao/mock_db_test.go database/sql/driver Conn,Driver,Stmt,Rows
package dao_test

import (
	"database/sql"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/kamilsk/click/dao"
	"github.com/kamilsk/click/domain"
	"github.com/stretchr/testify/assert"
)

const (
	DSN  = "stub://localhost"
	UUID = domain.UUID("41ca5e09-3ce2-4094-b108-3ecc257c6fa4")
)

func TestMust_WithInvalidConfiguration(t *testing.T) {
	var configs = []dao.Configurator{dao.Connection("", "", 0, 0)}
	assert.Panics(t, func() { dao.Must(configs...) })
}

func TestMust_WithoutConfiguration(t *testing.T) {
	var configs []dao.Configurator
	assert.NotPanics(t, func() { dao.Must(configs...) })
}

func TestStorage(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	var (
		drv = NewMockDriver(ctrl)
	)

	var configs = []dao.Configurator{dao.Connection(t.Name(), DSN, 1, 1)}
	sql.Register(t.Name(), drv)
	service, err := dao.New(configs...)
	assert.NoError(t, err)

	assert.NotNil(t, service.Connection())
	assert.Equal(t, "postgres", service.Dialect())
}
