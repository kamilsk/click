package domain

import "strings"

const (
	namespaceHeader = "X-Click-Namespace"
	optionsHeader   = "X-Click-Options"
	passQueryParam  = "url"
)

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

// Namespace TODO issue#131
func (context RedirectContext) Namespace() ID {
	return ID(get(context.Headers, namespaceHeader))
}

// Option TODO issue#131
func (context RedirectContext) Option() Option {
	split := func(str string) []string {
		ss := strings.Split(str, ";")
		for i, s := range ss {
			ss[i] = strings.ToLower(strings.TrimSpace(s))
		}
		return ss
	}
	is := func(where []string, what string) bool {
		for _, opt := range where {
			if opt == what {
				return true
			}
		}
		return false
	}
	options := split(get(context.Headers, optionsHeader))
	return Option{
		Anonymously: is(options, "anonym"),
		Debug:       is(options, "debug"),
		NoLog:       is(options, "nolog"),
	}
}

// Redirect TODO issue#131
func (context RedirectContext) Redirect() string {
	return get(context.Queries, passQueryParam)
}

func get(where map[string][]string, what string) string {
	if len(where[what]) > 0 {
		return where[what][0]
	}
	return ""
}
