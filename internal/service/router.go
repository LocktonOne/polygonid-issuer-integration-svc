package service

import (
	"github.com/go-chi/chi"
	"gitlab.com/distributed_lab/ape"
	"gitlab.com/tokene/polygonid-issuer-integration/internal/service/handlers"
	"gitlab.com/tokene/polygonid-issuer-integration/internal/service/helpers"
)

func (s *service) router() chi.Router {
	r := chi.NewRouter()

	r.Use(
		ape.RecoverMiddleware(s.log),
		ape.LoganMiddleware(s.log),
		ape.CtxMiddleware(
			helpers.CtxLog(s.log),
			helpers.CtxIssuerConfig(s.config.IssuerConfig()),
		),
	)
	r.Route("/integrations/polygonid-integration-svc", func(r chi.Router) {
		// identity = issuer did
		r.Route("/identities", func(r chi.Router) {

			r.Post("/", handlers.CreateIdentity) // create issuer (identity)

			// get identities
			r.Get("/", handlers.GetIdentities)
			r.Get("/{did}", handlers.GetIdentityDetail)

			// get identity
		})
		r.Route("/{did}", func(r chi.Router) {
			r.Route("/claims", func(r chi.Router) {
				// create credential
				r.Post("/", handlers.CreateCredential)
				// list of credentials
				r.Get("/", handlers.ListCredentials)

				r.Get("/revocation-status/{nonce}", handlers.GetRevocationStatus)
				r.Post("/revoke/{nonce}", handlers.RevokeCredential)

				r.Route("/{claim}", func(r chi.Router) {
					// get credential
					r.Get("/", handlers.GetCredential)
					r.Get("/qrcode", handlers.GetCredentialQRCode)
					// revoke credential
				})
			})
			r.Route("/state/publish", func(r chi.Router) {
				r.Post("/", handlers.PublishIdentityState)
				r.Post("/retry", handlers.RetryPublishIdentityState)
			})
		})
	})

	return r
}
