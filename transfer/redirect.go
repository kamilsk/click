package transfer

import "github.com/kamilsk/click/domain"

// RedirectRequest represents `GET /{Alias}` request.
type RedirectRequest struct {
	Namespace string
	URN       string
}

// RedirectResponse represents `GET /{Alias}` response.
type RedirectResponse struct {
	Alias  domain.Alias
	Target domain.Target
	Error  error
}
