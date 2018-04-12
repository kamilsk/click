package transfer

import "github.com/kamilsk/click/domain"

// RedirectRequest represents `GET /{Alias.URN}` request.
type RedirectRequest struct {
	EncryptedMarker string
	Namespace       string
	URN             string
	Query           map[string][]string
}

// RedirectResponse represents `GET /{Alias.URN}` response.
type RedirectResponse struct {
	EncryptedMarker string
	Alias           domain.Alias
	Target          domain.Target
	Error           error
}
