package service

import (
	"github.com/go-chi/chi"
	"gitlab.com/distributed_lab/ape"
	"gitlab.com/tokene/polygonid-issuer-integration/internal/service/handlers"
)

func (s *service) router() chi.Router {
	r := chi.NewRouter()

	r.Use(
		ape.RecoverMiddleware(s.log),
		ape.LoganMiddleware(s.log),
		ape.CtxMiddleware(
			handlers.CtxLog(s.log),
		),
	)
	r.Route("/integrations/polygonid-integration-svc", func(r chi.Router) {
		r.Route("claim/", func(r chi.Router) {
			// create credential
			// list of credentials
			r.Route("{id}", func(r chi.Router) {
				// update credential
				// get credential
			})
		})
	})

	return r
}
