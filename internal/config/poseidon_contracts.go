package config

import (
	"gitlab.com/distributed_lab/figure"
	"gitlab.com/distributed_lab/kit/kv"
	"gitlab.com/distributed_lab/logan/v3/errors"
)

type PoseidonContractsConfig struct {
	El1 *string `fig:"el1"`
	El2 *string `fig:"el2"`
	El3 *string `fig:"el3"`
}

func (c *config) PoseidonContractsConfig() PoseidonContractsConfig {
	return c.poseidonContracts.Do(func() interface{} {
		poseidonContractsConfig := PoseidonContractsConfig{}
		err := figure.
			Out(&poseidonContractsConfig).
			With(figure.BaseHooks, ecdsaHook).
			From(kv.MustGetStringMap(c.getter, "poseidon_contracts")).
			Please()
		if err != nil {
			panic(errors.Wrap(err, "failed to figure out auth service config"))
		}
		return poseidonContractsConfig
	}).(PoseidonContractsConfig)
}
