package chi_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"go.octolab.org/ecosystem/click/internal/server/router"
	. "go.octolab.org/ecosystem/click/internal/server/router/chi"
)

func TestNewRouter(t *testing.T) {
	type server struct{ router.Server }
	assert.NotPanics(t, func() { _ = NewRouter(server{}) })
}
