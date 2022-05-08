package injectors

import (
	"chatroom/config"
	"chatroom/infrastructure"
	"database/sql"
)

func ProvideSqlDBConnector(config *config.GeneralConfig) (*sql.DB, error) {
	return infrastructure.ConDb(&config.PostgreSqlDB)
}
