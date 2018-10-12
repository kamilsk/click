package transfer

import (
	"github.com/kamilsk/click/pkg/domain"
	"github.com/kamilsk/click/pkg/service"
)

// RedirectRequest represents `GET /{Alias.URN}` request.
type RedirectRequest struct {
	Namespace domain.ID
	URN       string
	Query     map[string][]string
	Option    service.Option
	Redirect  domain.Redirect
}

// RedirectResponse represents `GET /{Alias.URN}` response.
type RedirectResponse struct {
	Target domain.Target
	Error  error
}
