package transfer

// PassRequest represents `GET /pass?url={URI}` request.
type PassRequest struct{}

// PassResponse represents `GET /pass?url={URI}` response.
type PassResponse struct {
	Error error
}
