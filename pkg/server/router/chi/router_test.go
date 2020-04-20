package chi_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/kamilsk/click/pkg/server/router"
	. "github.com/kamilsk/click/pkg/server/router/chi"
)

func TestNewRouter(t *testing.T) {
	type server struct{ router.Server }
	assert.NotPanics(t, func() { _ = NewRouter(server{}) })
}
