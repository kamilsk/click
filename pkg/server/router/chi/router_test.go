package chi_test

import (
	"testing"

	"github.com/kamilsk/click/pkg/server/router"
	. "github.com/kamilsk/click/pkg/server/router/chi"
	"github.com/stretchr/testify/assert"
)

func TestNewRouter(t *testing.T) {
	type server struct{ router.Server }
	assert.NotPanics(t, func() { _ = NewRouter(server{}) })
}
