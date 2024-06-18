package config

import (
	"gitlab.com/distributed_lab/kit/comfig"
	"gitlab.com/distributed_lab/kit/copus"
	"gitlab.com/distributed_lab/kit/copus/types"
	"gitlab.com/distributed_lab/kit/kv"
	"gitlab.com/distributed_lab/kit/pgdb"
	doormanCfg "gitlab.com/tokene/doorman/connector/config"
)

type Config interface {
	comfig.Logger
	pgdb.Databaser
	types.Copuser
	comfig.Listenerer
	doormanCfg.DoormanConfiger

	IssuerConfig() IssuerConfig
	NetworkConfig() NetworkConfig
}

type config struct {
	comfig.Logger
	pgdb.Databaser
	types.Copuser
	comfig.Listenerer
	doormanCfg.DoormanConfiger

	getter kv.Getter

	issuer            comfig.Once
	network           comfig.Once
	poseidonContracts comfig.Once
}

func New(getter kv.Getter) Config {
	return &config{
		getter:          getter,
		Databaser:       pgdb.NewDatabaser(getter),
		DoormanConfiger: doormanCfg.NewDoormanConfiger(getter),
		Copuser:         copus.NewCopuser(getter),
		Listenerer:      comfig.NewListenerer(getter),
		Logger:          comfig.NewLogger(getter, comfig.LoggerOpts{}),
	}
}
