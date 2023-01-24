package config

type MySQL struct {
	Address  string `env:"MYSQL_ADDRESS" envDefault:"localhost:3306"`
	Username string `env:"MYSQL_USERNAME" envDefault:"root"`
	Password string `env:"MYSQL_PASSWORD" envDefault:"root"`
	Database string `env:"MYSQL_DATABASE" envDefault:"boilerplate"`
}

type I18n struct {
	Fallback string `env:"I18N_FALLBACK_LANGUAGE" envDefault:"en"`
	Dir      string `env:"I18N_DIR" envDefault:"./src/locales"`
	Locales  string `env:"I18N_LOCALES" envDefault:"en,tr"`
}

type Server struct {
	Host string `env:"SERVER_HOST" envDefault:"localhost"`
	Port int    `env:"SERVER_PORT" envDefault:"3000"`
}

type Cors struct {
	AllowedOrigins   []string `env:"CORS_ALLOWED_ORIGINS" envSeparator:"," envDefault:"*"`
	AllowedMethods   []string `env:"CORS_ALLOWED_METHODS" envSeparator:"," envDefault:"GET,POST,PUT,DELETE,OPTIONS"`
	AllowedHeaders   []string `env:"CORS_ALLOWED_HEADERS" envSeparator:"," envDefault:"*"`
	AllowCredentials bool     `env:"CORS_ALLOW_CREDENTIALS" envDefault:"true"`
}

type App struct {
	Protocol string `env:"PROTOCOL" envDefault:"http"`
	MySQL    MySQL
	Server   Server
	Cors     Cors
	I18n     I18n
}
