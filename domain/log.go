package domain

// Log represents a "redirect event".
//go:generate easyjson -all
type Log struct {
	ID        uint64   `json:"id"`
	LinkID    string   `json:"link_id"`
	AliasID   uint64   `json:"alias_id"`
	TargetID  uint64   `json:"target_id"`
	URI       string   `json:"uri"`
	Code      int      `json:"code"`
	Context   Metadata `json:"context"`
	CreatedAt string   `json:"created_at"`
}

// Metadata contains context information about a "redirect event".
type Metadata struct {
	Cookie map[string]string   `json:"cookie,omitempty"`
	Header map[string][]string `json:"header,omitempty"`
}
