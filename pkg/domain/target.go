package domain

import "sort"

// Target represents a target of the Link.
//go:generate easyjson -all
type Target struct {
	ID   ID     `json:"id"`
	Rule Rule   `json:"rule"`
	URL  string `json:"url"`

	weight int
}

// Targets holds set of Target of the Link and provides useful methods for the set.
type Targets []Target

// Find tries to find a suitable target of the Link for the specific context
// or returns an empty Target if nothing found.
func (set Targets) Find(alias Alias, query map[string][]string) (Target, bool) {
	var index, max = -1, 0
	for i, target := range set {
		if target.weight = target.Rule.Calculate(alias, query); target.weight >= max {
			index, max = i, target.weight
		}
	}
	if index > -1 {
		return set[index], true
	}
	return Target{}, false
}

type targetsByID Targets

func (set targetsByID) Len() int { return len(set) }

func (set targetsByID) Less(i, j int) bool { return set[i].ID < set[j].ID }

func (set targetsByID) Swap(i, j int) { set[i], set[j] = set[j], set[i] }
