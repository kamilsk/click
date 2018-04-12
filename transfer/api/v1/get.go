package v1

import "github.com/kamilsk/click/domain"

// GetRequest represents `GET /api/v1/{Link.ID}` request.
type GetRequest struct {
	ID domain.UUID
}

// GetResponse represents `GET /api/v1/{Link.ID}` response.
type GetResponse struct {
	Link  domain.Link
	Error error
}
