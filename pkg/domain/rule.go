package domain

const (
	// AND used by rules as `&&`.
	AND byte = iota
	// OR used by rules as `||`.
	OR
)

const (
	tagKey = "tag"
)

// Rule represents a rule of the Target.
//go:generate easyjson -all
type Rule struct {
	Description string            `json:"description,omitempty"`
	AliasID     uint64            `json:"alias,omitempty"`
	Tags        []string          `json:"tags,omitempty"`
	Conditions  map[string]string `json:"conditions,omitempty"`
	Match       byte              `json:"match,omitempty"`
}

// Calculate calculates weight of Rule's Target.
func (v Rule) Calculate(alias Alias, query map[string][]string) int {
	// default rule
	if v.AliasID == 0 && len(v.Tags) == 0 && len(v.Conditions) == 0 {
		return 1
	}
	// other rules will be ordered by weight and ID
	toCheck := make([]func(Rule, Alias, map[string][]string) bool, 0, 3)
	if v.AliasID != 0 {
		toCheck = append(toCheck, checkAlias)
	}
	if len(v.Tags) > 0 {
		toCheck = append(toCheck, checkTags)
	}
	if len(v.Conditions) > 0 {
		toCheck = append(toCheck, checkConditions)
	}
	weight := 0
	switch v.Match {
	case OR:
		for _, check := range toCheck {
			if check(v, alias, query) {
				weight += 2
			}
		}
	case AND:
		for _, check := range toCheck {
			if !check(v, alias, query) {
				return 0
			}
			weight += 4
		}
	}
	return weight
}

func checkAlias(rule Rule, alias Alias, _ map[string][]string) bool {
	return rule.AliasID == alias.ID
}

func checkTags(rule Rule, _ Alias, query map[string][]string) bool {
	if tags, ok := query[tagKey]; ok && len(tags) > 0 {
		tag := tags[0]
		for i := range rule.Tags {
			if tag == rule.Tags[i] {
				return true
			}
		}
	}
	return false
}

func checkConditions(rule Rule, _ Alias, query map[string][]string) bool {
	for key, value := range rule.Conditions {
		if values, ok := query[key]; ok {
			if len(values) > 0 && values[0] == value {
				return true
			}
		}
	}
	return false
}
