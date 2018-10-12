package domain

import "sort"

// Target represents a target of the Link.
//go:generate easyjson -all
type Target struct {
	ID   ID     `json:"id"`
	URI  string `json:"uri"`
	Rule Rule   `json:"rule"`
}

// Targets holds set of Target of the Link and provides useful methods for the set.
type Targets []Target

// Find tries to find a suitable target of the Link for the specific context
// or returns an empty Target if nothing found.
func (set Targets) Find(alias Alias, query map[string][]string) Target {
	var (
		result     Target
		index, max = -1, 0
	)
	sort.Sort(sort.Reverse(targetsByID(set)))
	for i, target := range set {
		if weight := target.Rule.Calculate(alias, query); weight >= max {
			index, max = i, weight
		}
	}
	if index > -1 {
		result = set[index]
	}
	return result
}

type targetsByID Targets

func (set targetsByID) Len() int { return len(set) }

func (set targetsByID) Less(i, j int) bool { return set[i].ID < set[j].ID }

func (set targetsByID) Swap(i, j int) { set[i], set[j] = set[j], set[i] }
