package service

// Option contains rules for request processing.
type Option struct {
	// Anonymously: use zero-token instead of origin.
	Anonymously bool
	// Debug: return debug information to the client.
	Debug bool
	// NoLog: do not log link navigation.
	NoLog bool
}
