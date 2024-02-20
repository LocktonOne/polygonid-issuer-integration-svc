package config

import (
	"net/url"
	"time"

	"gitlab.com/distributed_lab/figure"
	"gitlab.com/distributed_lab/kit/kv"
	"gitlab.com/distributed_lab/logan/v3/errors"
)

type IssuerConfig struct {
	Endpoint         *url.URL       `fig:"endpoint,required"`
	IssuerDID        string         `fig:"issuer_did"`
	Username         string         `fig:"username,required"`
	Password         string         `fig:"password,required"`
	CredentialSchema string         `fig:"credential_schema"`
	CredentialType   string         `fig:"credential_type"`
	Expiration       *time.Duration `fig:"expiration"`
}

func (c *config) IssuerConfig() IssuerConfig {
	return c.issuer.Do(func() interface{} {
		issuerConfig := IssuerConfig{}
		err := figure.
			Out(&issuerConfig).
			From(kv.MustGetStringMap(c.getter, "issuer")).
			Please()
		if err != nil {
			panic(errors.Wrap(err, "failed to figure out auth service config"))
		}
		return issuerConfig
	}).(IssuerConfig)
}
