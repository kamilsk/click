package domain

const (
	AND byte = iota
	OR
	XOR
)

// Rule represents a rule of the Target.
//go:generate easyjson -all
type Rule struct {
	Description string            `json:"description"`
	AliasID     uint64            `json:"alias"`
	Tags        []string          `json:"tag"`
	Conditions  map[string]string `json:"conditions"`
	Match       byte              `json:"match"`
}
