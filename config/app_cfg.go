package config

type AppCfg struct {
	Environment string `mapstructure:"environment" json:"environment"`
	HttpHost    string `mapstructure:"http_host" json:"http_host`
	HttpPort    int    `mapstructure:"http_port" json:"http_port"`
}
