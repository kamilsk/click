package transfer

// PassRequest represents `GET /pass?url={URI}` request.
type PassRequest struct {
	EncryptedMarker string
}

// PassResponse represents `GET /pass?url={URI}` response.
type PassResponse struct {
	EncryptedMarker string
	Error           error
}
