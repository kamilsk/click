package transfer

import "github.com/kamilsk/click/pkg/domain"

// RedirectRequest represents `GET /{Alias.URN}` request.
type RedirectRequest struct {
	Namespace string
	URN       string
	Query     map[string][]string
}

// RedirectResponse represents `GET /{Alias.URN}` response.
type RedirectResponse struct {
	Alias  domain.Alias
	Target domain.Target
	Error  error
}
