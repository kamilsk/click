package domain

// Log represents a "redirect event".
//go:generate easyjson -all
type Log struct {
	ID        uint64       `json:"id"`
	LinkID    string       `json:"link_id"`
	AliasID   uint64       `json:"alias_id"`
	TargetID  uint64       `json:"target_id"`
	URI       string       `json:"uri"`
	Code      int8         `json:"code"`
	Context   EventContext `json:"context"`
	CreatedAt string       `json:"created_at"`
}

// EventContext contains context information about "redirect event".
type EventContext struct {
	Cookie map[string]string   `json:"cookie,omitempty"`
	Header map[string][]string `json:"header,omitempty"`
}
