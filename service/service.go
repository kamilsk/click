package service

import (
	"github.com/kamilsk/click/domain"
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
	link, err := s.dao.LinkByAlias(domain.Alias{Namespace: request.Namespace, URN: request.URN})
	if err != nil {
		response.Error = err
		return response
	}

	// simple logic now
	response.Alias = link.Aliases[0]
	if len(link.Targets) > 0 {
		response.Target = link.Targets[0]
	}

	return response
}
