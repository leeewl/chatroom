package infrastructure

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	_ "github.com/lib/pq"
)

var db *sql.DB

type DBDriver string

type PostgreSqlDBConfig struct {
	Driver          DBDriver      `mapstructure:"driver" json:"driver"`
	DBName          string        `mapstructure:"db_name" json:"db_name"`
	DBHost          string        `mapstructure:"db_host" json:"db_host"`
	DBPort          int           `mapstructure:"db_port" json:"db_port"`
	DialTimeOut     time.Duration `mapstructure:"dial_timeout" json:"dial_timeout"` // second
	MaxIdleConns    int           `mapstructure:"max_idle_conns" json:"max_idle_conns"`
	MaxOpenConns    int           `mapstructure:"max_open_conns" json:"max_open_conns"`
	ConnMaxLifetime time.Duration `mapstructure:"conn_max_life_time" json:"conn_max_life_time"` // second
	Username        string        `mapstructure:"username" json:"username"`
	Password        string        `mapstructure:"password" json:"password"`
	SSLMode         string        `mapstructure:"ssl_mode" json:"ssl_mode"`
	URI             string        `mapstructure:"uri" json:"uri"`
}

func ConDb(config *PostgreSqlDBConfig) (db *sql.DB, err error) {
	connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		config.DBHost, config.DBPort, config.Username, config.Password, config.DBName)
	db, err = sql.Open(string(config.Driver), connStr)
	if err != nil {
		return
	}
	ctx := context.Background()
	err = db.PingContext(ctx)
	if err != nil {
		return
	}
	return
}
