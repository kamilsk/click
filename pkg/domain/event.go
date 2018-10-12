package domain

// RedirectEvent represents a redirect event.
//go:generate easyjson -all
type RedirectEvent struct {
	ID         uint64          `json:"id"`
	LinkID     ID              `json:"link_id"`
	AliasID    ID              `json:"alias_id"`
	TargetID   ID              `json:"target_id"`
	Identifier ID              `json:"identifier"`
	URI        string          `json:"uri"`
	Code       int             `json:"code"`
	Context    RedirectContext `json:"context"`
}

// RedirectContext contains context information about a "redirect event".
//go:generate easyjson -all
type RedirectContext struct {
	RequestID string              `json:"request_id,omitempty"`
	Cookie    map[string]string   `json:"cookie,omitempty"`
	Header    map[string][]string `json:"header,omitempty"`
	Query     map[string][]string `json:"query,omitempty"`
}
