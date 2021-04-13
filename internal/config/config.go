package config

type Postgres struct {
	User 		string `env:"DB_USER"`
	Pass 		string `env:"DB_PASS"`
	Host 		string `env:"DB_HOST"`
	Port 		string `env:"DB_PORT"`
	Database 	string `env:"DB_DATABASE"`
	Ssl 		string `env:"DB_SSL"`
}

func (p Postgres) GetUser() string {
	return p.User
}

func (p Postgres) GetPassword() string {
	return p.Pass
}

func (p Postgres) GetHost() string {
	return p.Host
}

func (p Postgres) GetPort() string {
	return p.Port
}

func (p Postgres) GetDatabase() string {
	return p.Database
}

func (p Postgres) GetSSL() string {
	return p.Ssl
}