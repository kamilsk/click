package domain

// Link represents a "redirect entity".
//go:generate easyjson -all
type Link struct {
	ID      ID      `json:"id"`
	Name    string  `json:"name"`
	Aliases Aliases `json:"aliases"`
	Targets Targets `json:"targets"`
}
