package middlewares

import (
	"context"
	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/ape/problems"
	"gitlab.com/tokene/polygonid-issuer-integration/internal/service/helpers"
	"net/http"
)

func Login() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			address, token, err := helpers.ValidateJwt(r)
			if err != nil {
				helpers.Log(r).WithError(err).Error("failed to validate token")
				ape.Render(w, problems.BadRequest(err))
				return
			}
			//doorman := helpers.DoormanConnector(r)
			//if err = doorman.CheckPermissionID("CREATE", "*", token); err != nil {
			//	helpers.Log(r).WithError(err).Error("haven't permission for this operation")
			//	ape.Render(w, problems.BadRequest(err))
			//	return
			//}
			ctx := context.WithValue(r.Context(), "token", token)
			ctx = context.WithValue(ctx, "address", address)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}
