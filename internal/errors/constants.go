package errors

const (
	// ClientErrorCode is a code of a client error.
	ClientErrorCode = iota
	// ResourceNotFoundCode is a code of a client error when the requested resource does not exist.
	ResourceNotFoundCode

	// ClientErrorMessage is a default message of a client error.
	ClientErrorMessage = "error"
	// LinkNotFoundMessage is a default message of the case when a Link does not exist.
	LinkNotFoundMessage = "link not found"
)

const (
	// ServerErrorCode is a code of a server error.
	ServerErrorCode = 100 + iota
	// DatabaseFailCode is a code of the server error related to database problems.
	DatabaseFailCode
	// SerializationFailCode is a code of the server error related to serialization problems.
	SerializationFailCode

	// NeutralMessage is a default message.
	NeutralMessage = "something went wrong"
	// ServerErrorMessage is a default message of a server error.
	ServerErrorMessage = "server error"
)
