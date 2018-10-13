package domain

const passQueryParam = "url"

// Option contains rules for request processing.
type Option struct {
	// Anonymously: use zero-identifier instead of origin.
	Anonymously bool
	// Debug: return debug information to the client.
	Debug bool
	// NoLog: do not log link navigation.
	NoLog bool
}

// RedirectEvent represents a redirect event.
//go:generate easyjson -all
type RedirectEvent struct {
	NamespaceID ID              `json:"namespace_id"`
	LinkID      *ID             `json:"link_id,omitempty"`
	AliasID     *ID             `json:"alias_id,omitempty"`
	TargetID    *ID             `json:"target_id,omitempty"`
	Identifier  *ID             `json:"identifier,omitempty"`
	Context     RedirectContext `json:"context"`
	Code        int             `json:"code"`
	URL         string          `json:"url"`
}

// Redirect TODO issue#131
func (event RedirectEvent) Redirect() string {
	if event.URL != "" {
		return event.URL
	}
	return event.Context.Redirect()
}

// RedirectContext contains context information about a "redirect event".
//go:generate easyjson -all
type RedirectContext struct {
	Cookies map[string]string   `json:"cookies,omitempty"`
	Headers map[string][]string `json:"headers,omitempty"`
	Queries map[string][]string `json:"queries,omitempty"`
}

// Redirect TODO issue#131
func (context RedirectContext) Redirect() string {
	if len(context.Queries[passQueryParam]) > 0 {
		return context.Queries[passQueryParam][0]
	}
	return ""
}
