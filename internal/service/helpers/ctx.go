package helpers

import (
	"context"
	"gitlab.com/tokene/doorman/connector"
	"gitlab.com/tokene/polygonid-issuer-integration/internal/config"
	"gitlab.com/tokene/polygonid-issuer-integration/internal/data"
	"net/http"

	"gitlab.com/distributed_lab/logan/v3"
)

type ctxKey int

const (
	logCtxKey ctxKey = iota
	issuerConfigKey
	networkConfigKey
	doormanConnectorCtxKey
	didCtxKey
)

func CtxLog(entry *logan.Entry) func(context.Context) context.Context {
	return func(ctx context.Context) context.Context {
		return context.WithValue(ctx, logCtxKey, entry)
	}
}

func Log(r *http.Request) *logan.Entry {
	return r.Context().Value(logCtxKey).(*logan.Entry)
}

func CtxIssuerConfig(entry config.IssuerConfig) func(context.Context) context.Context {
	return func(ctx context.Context) context.Context {
		return context.WithValue(ctx, issuerConfigKey, entry)
	}
}

func IssuerConfig(r *http.Request) config.IssuerConfig {
	return r.Context().Value(issuerConfigKey).(config.IssuerConfig)
}

func CtxNetworkConfig(entry config.NetworkConfig) func(context.Context) context.Context {
	return func(ctx context.Context) context.Context {
		return context.WithValue(ctx, networkConfigKey, entry)
	}
}

func NetworkConfig(r *http.Request) config.NetworkConfig {
	return r.Context().Value(networkConfigKey).(config.NetworkConfig)
}

func DoormanConnector(r *http.Request) connector.ConnectorI {
	return r.Context().Value(doormanConnectorCtxKey).(connector.ConnectorI)
}

func CtxDoormanConnector(entry connector.ConnectorI) func(context.Context) context.Context {
	return func(ctx context.Context) context.Context {
		return context.WithValue(ctx, doormanConnectorCtxKey, entry)
	}
}

func CtxDidQ(entry data.DidQ) func(context.Context) context.Context {
	return func(ctx context.Context) context.Context {
		return context.WithValue(ctx, didCtxKey, entry)
	}
}

func DidQ(r *http.Request) data.DidQ {
	return r.Context().Value(didCtxKey).(data.DidQ).New()
}

func Token(r *http.Request) string {
	return r.Context().Value("token").(string)
}

func Address(r *http.Request) string {
	return r.Context().Value("address").(string)
}
