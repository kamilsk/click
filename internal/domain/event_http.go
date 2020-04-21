package domain

import (
	"net/http"
	"net/url"
	"strings"
)

const (
	cookieHeader     = "Cookie"
	identifierHeader = "X-Passport-ID"
	namespaceHeader  = "X-Click-Namespace"
	optionsHeader    = "X-Click-Options"
	refererHeader    = "Referer"
	requestHeader    = "X-Request-ID"
	userAgentHeader  = "User-Agent"
	passQueryParam   = "url"
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

// Header TODO issue#131
func (context RedirectContext) Header(key string) string {
	return http.Header(context.Headers).Get(key)
}

// Identifier TODO issue#131
func (context RedirectContext) Identifier() *ID {
	header := context.Header(identifierHeader)
	if header != "" {
		id := ID(header)
		return &id
	}
	return nil
}

// Namespace TODO issue#131
func (context RedirectContext) Namespace() ID {
	return ID(context.Header(namespaceHeader))
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
	options := split(context.Header(optionsHeader))
	return Option{
		Anonymously: is(options, "anonym"),
		Debug:       is(options, "debug"),
		NoLog:       is(options, "nolog"),
	}
}

// Redirect TODO issue#131
func (context RedirectContext) Redirect() string {
	return url.Values(context.Queries).Get(passQueryParam)
}

// Referer TODO issue#131
func (context RedirectContext) Referer() string {
	return context.Header(refererHeader)
}

// Request TODO issue#131
func (context RedirectContext) Request() *ID {
	header := context.Header(requestHeader)
	if header != "" {
		id := ID(header)
		return &id
	}
	return nil
}

// UserAgent TODO issue#131
func (context RedirectContext) UserAgent() string {
	return context.Header(userAgentHeader)
}
