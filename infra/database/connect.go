package database

import (
	"fmt"

	"github.com/takahiroaoki/batch-hub/config"
)

func DataSourceName(config config.DatabaseConfig) string {
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4", config.User(), config.Password(), config.Host(), config.Port(), config.Database())
}

func DatabaseURL(config config.DatabaseConfig) string {
	return fmt.Sprintf("%s://%s", "mysql", DataSourceName(config))
}
