package config

import (
	"chatroom/infrastructure"
)

type GeneralConfig struct {
	App          AppCfg                            `mapstructure:"app" json:"app"`
	PostgreSqlDB infrastructure.PostgreSqlDBConfig `mapstructure:"postgresqldb" json:"postgresqldb"`
}
