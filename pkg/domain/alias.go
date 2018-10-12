package domain

// Alias represents an alias of the Link.
//go:generate easyjson -all
type Alias struct {
	ID        ID     `json:"id"`
	Namespace ID     `json:"namespace"`
	URN       string `json:"urn"`
}

// Aliases holds set of Alias of the Link and provides useful methods for the set.
type Aliases []Alias

// Find tries to find a suitable alias by provided namespace and URN
// or returns an empty Alias if nothing found.
func (set Aliases) Find(ns ID, urn string) (Alias, bool) {
	for _, alias := range set {
		if alias.Namespace == ns && alias.URN == urn {
			return alias, true
		}
	}
	return Alias{}, false
}
