package config

type DatabaseConfig interface {
	Host() string
	Port() string
	Database() string
	User() string
	Password() string
}

type databaseConfig struct{}

func (d databaseConfig) Host() string {
	return env.dbHost
}

func (d databaseConfig) Port() string {
	return env.dbPort
}

func (d databaseConfig) Database() string {
	return env.dbDatabase
}

func (d databaseConfig) User() string {
	return env.dbUser
}

func (d databaseConfig) Password() string {
	return env.dbPassword
}

func NewDatabaseConfig() DatabaseConfig {
	return databaseConfig{}
}
