package v1

import "go.octolab.org/ecosystem/click/internal/domain"

// GetRequest represents `GET /api/v1/{Link.ID}` request.
type GetRequest struct {
	ID domain.ID
}

// GetResponse represents `GET /api/v1/{Link.ID}` response.
type GetResponse struct {
	Error error
	Link  domain.Link
}
