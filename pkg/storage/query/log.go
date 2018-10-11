package query

import "github.com/kamilsk/click/pkg/domain"

// WriteLog TODO issue#131
type WriteLog struct {
	LinkID          domain.ID
	AliasID         domain.ID
	TargetID        domain.ID
	Identifier      domain.ID
	URI             string
	Code            int
	RedirectContext domain.RedirectContext
}
