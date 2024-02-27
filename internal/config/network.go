package config

import (
	"crypto/ecdsa"
	"fmt"
	"github.com/ethereum/go-ethereum/crypto"
	"gitlab.com/distributed_lab/figure"
	"gitlab.com/distributed_lab/kit/kv"
	"gitlab.com/distributed_lab/logan/v3/errors"
	"net/url"
	"reflect"
)

type NetworkConfig struct {
	RpcUrl     *url.URL          `fig:"rpc_url,required"`
	PrivateKey *ecdsa.PrivateKey `fig:"private_key,required"`
	ChainId    int64             `fig:"chain_id,required"`
	Network    string            `fig:"network,required"`
	GasLimit   uint64            `fig:"gas_limit,required"`
}

func (c *config) NetworkConfig() NetworkConfig {
	return c.network.Do(func() interface{} {
		networkConfig := NetworkConfig{}
		err := figure.
			Out(&networkConfig).
			With(figure.BaseHooks, ecdsaHook).
			From(kv.MustGetStringMap(c.getter, "network")).
			Please()
		if err != nil {
			panic(errors.Wrap(err, "failed to figure out auth service config"))
		}
		return networkConfig
	}).(NetworkConfig)
}

var ecdsaHook = figure.Hooks{
	"*ecdsa.PrivateKey": func(value interface{}) (reflect.Value, error) {
		switch v := value.(type) {
		case string:
			privKey, err := crypto.HexToECDSA(v)
			if err != nil {
				return reflect.Value{}, errors.Wrap(err, "invalid hex private key")
			}
			return reflect.ValueOf(privKey), nil
		default:
			return reflect.Value{}, fmt.Errorf("unsupported conversion from %T", value)
		}
	},
}
