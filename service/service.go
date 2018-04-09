package service

import (
	"github.com/kamilsk/click/transfer"
	"github.com/kamilsk/click/transfer/api/v1"
)

// New returns a new instance of Click! service.
func New(dao Storage) *Click {
	return &Click{dao: dao}
}

// Click is the primary application service.
type Click struct {
	dao Storage
}

// HandleGetV1 handles an input request.
func (s *Click) HandleGetV1(request v1.GetRequest) v1.GetResponse {
	var response v1.GetResponse
	response.Link, response.Error = s.dao.Link(request.ID)
	return response
}

// HandleRedirect handles an input request.
func (s *Click) HandleRedirect(request transfer.RedirectRequest) transfer.RedirectResponse {
	var response transfer.RedirectResponse
	link, err := s.dao.LinkByAlias(request.Namespace, request.URN)
	if err != nil {
		response.Error = err
		return response
	}
	response.Alias = link.Aliases.Find(request.Namespace, request.URN)
	response.Target = link.Targets.Find(response.Alias, request.Query)
	return response
}
