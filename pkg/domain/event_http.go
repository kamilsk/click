package domain

import (
	"net/http"
	"strings"
)

const (
	cookieHeader    = "Cookie"
	namespaceHeader = "X-Click-Namespace"
	optionsHeader   = "X-Click-Options"
)

// FromCookies returns converted value from request's cookies.
func FromCookies(cookies []*http.Cookie) map[string]string {
	converted := make(map[string]string)
	for _, cookie := range cookies {
		if cookie.HttpOnly && cookie.Secure {
			converted[cookie.Name] = cookie.Value
		}
	}
	return converted
}

// FromHeaders returns converted value from request's headers.
func FromHeaders(headers http.Header) map[string][]string {
	converted := make(map[string][]string)
	for key, values := range headers {
		if !strings.EqualFold(cookieHeader, key) {
			converted[key] = values
		}
	}
	return converted
}

// FromRequest returns converted value from a request.
func FromRequest(req *http.Request) map[string][]string {
	return req.URL.Query()
}

// Namespace TODO issue#131
func (context RedirectContext) Namespace() ID {
	return ID(http.Header(context.Headers).Get(namespaceHeader))
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
	options := split(http.Header(context.Headers).Get(optionsHeader))
	return Option{
		Anonymously: is(options, "anonym"),
		Debug:       is(options, "debug"),
		NoLog:       is(options, "nolog"),
	}
}
