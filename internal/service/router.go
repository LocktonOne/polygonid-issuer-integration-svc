package service

import (
	"github.com/go-chi/chi"
	"gitlab.com/distributed_lab/ape"
	"gitlab.com/tokene/polygonid-issuer-integration/internal/data/postgres"
	"gitlab.com/tokene/polygonid-issuer-integration/internal/service/handlers"
	"gitlab.com/tokene/polygonid-issuer-integration/internal/service/helpers"
	"gitlab.com/tokene/polygonid-issuer-integration/internal/service/middlewares"
)

func (s *service) router() chi.Router {
	r := chi.NewRouter()

	r.Use(
		ape.RecoverMiddleware(s.log),
		ape.LoganMiddleware(s.log),
		ape.CtxMiddleware(
			helpers.CtxLog(s.log),
			helpers.CtxIssuerConfig(s.config.IssuerConfig()),
			helpers.CtxNetworkConfig(s.config.NetworkConfig()),
			helpers.CtxDoormanConnector(s.config.DoormanConnector()),
			helpers.CtxDidQ(postgres.NewDidQ(s.config.DB())),
		),

		middlewares.Login(),
	)
	r.Route("/integrations/polygonid-issuer-integration", func(r chi.Router) {
		// identity = issuer did
		r.Route("/identities", func(r chi.Router) {
			r.Post("/", handlers.CreateIdentity) // create issuer (identity)
			r.Get("/", handlers.GetIdentities)
			r.Get("/address/{address}", handlers.GetIdentityByAddress)
			r.Get("/{did}", handlers.GetIdentityDetail)
		})

		r.Route("/{did}", func(r chi.Router) {
			r.Route("/claims", func(r chi.Router) {
				r.Post("/", handlers.CreateCredential)
				r.Get("/", handlers.ListCredentials)

				r.Get("/revocation-status/{nonce}", handlers.GetRevocationStatus)
				r.Post("/revoke/{nonce}", handlers.RevokeCredential)

				r.Route("/{claim}", func(r chi.Router) {
					r.Get("/", handlers.GetCredential)
					r.Get("/qrcode", handlers.GetCredentialQRCode)
				})
			})
			r.Route("/state/publish", func(r chi.Router) {
				r.Post("/", handlers.PublishIdentityState)
				r.Post("/retry", handlers.RetryPublishIdentityState)
			})
		})
		r.Route("/verifier", func(r chi.Router) {
			r.Post("/", handlers.DeployOnChainVerifier)
			r.Post("/request/set", handlers.SetZKPRequest)
		})
		// The library linking is not working
		//r.Post("/state", handlers.DeployState)
		r.Post("/validator", handlers.DeployValidator)

	})

	return r
}
