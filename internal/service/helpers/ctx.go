package helpers

import (
	"context"
	"gitlab.com/tokene/polygonid-issuer-integration/internal/config"
	"net/http"

	"gitlab.com/distributed_lab/logan/v3"
)

type ctxKey int

const (
	logCtxKey ctxKey = iota
	issuerConfigKey
	networkConfigKey
	poseidonContractsConfigKey
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
