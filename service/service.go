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

// HandlePass handles an input request.
func (s *Click) HandlePass(request transfer.PassRequest) transfer.PassResponse {
	var response transfer.PassResponse

	{ // TODO encrypt/decrypt marker
		marker := domain.UUID(request.EncryptedMarker)
		if !marker.IsValid() {
			marker, response.Error = s.dao.UUID()
			if response.Error != nil {
				return response
			}
		}
		response.EncryptedMarker = string(marker)
	}

	return response
}

// HandleRedirect handles an input request.
func (s *Click) HandleRedirect(request transfer.RedirectRequest) transfer.RedirectResponse {
	var response transfer.RedirectResponse

	{ // TODO encrypt/decrypt marker
		marker := domain.UUID(request.EncryptedMarker)
		if !marker.IsValid() {
			marker, response.Error = s.dao.UUID()
			if response.Error != nil {
				return response
			}
		}
		response.EncryptedMarker = string(marker)
	}

	link, err := s.dao.LinkByAlias(request.Namespace, request.URN)
	if err != nil {
		response.Error = err
		return response
	}
	response.Alias = link.Aliases.Find(request.Namespace, request.URN)
	response.Target = link.Targets.Find(response.Alias, request.Query)
	response.Alias.LinkID, response.Target.LinkID = link.ID, link.ID
	return response
}

// LogRedirectEvent stores a "redirect event".
func (s *Click) LogRedirectEvent(event domain.Log) {
	_, _ = s.dao.Log(event)
}
