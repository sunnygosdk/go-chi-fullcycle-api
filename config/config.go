package config

import (
	"fmt"
	"strings"

	"github.com/go-chi/jwtauth"
	"github.com/spf13/viper"
)

type Config struct {
	dbDriver      string
	dbUser        string
	dbPassword    string
	dbHost        string
	dbPort        string
	dbName        string
	webServerPort string
	jwtSecret     string
	jwtExpiresIn  int
	environment   string
	tokenAuth     *jwtauth.JWTAuth
}

func LoadConfig() *Config {
	viper.SetConfigFile(".env")
	viper.AddConfigPath(".")
	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("config error: %w", err))
	}

	required := []string{"DB_PASSWORD", "JWT_SECRET"}
	for _, key := range required {
		if !viper.IsSet(key) || viper.GetString(key) == "" {
			panic(fmt.Errorf("missing required config: %s", key))
		}
	}

	return &Config{
		dbDriver:      viper.GetString("DB_DRIVER"),
		dbUser:        viper.GetString("DB_USER"),
		dbPassword:    viper.GetString("DB_PASSWORD"),
		dbHost:        viper.GetString("DB_HOST"),
		dbPort:        viper.GetString("DB_PORT"),
		dbName:        viper.GetString("DB_NAME"),
		webServerPort: viper.GetString("WEB_SERVER_PORT"),
		jwtSecret:     viper.GetString("JWT_SECRET"),
		jwtExpiresIn:  viper.GetInt("JWT_EXPIRES_IN"),
		environment:   viper.GetString("ENVIRONMENT"),
		tokenAuth:     jwtauth.New("HS256", []byte(viper.GetString("JWT_SECRET")), nil),
	}
}

// GetConnectionInfo returns the database connection info.
// The first value is a boolean that indicates if the connection is for test.
// The second value is the connection string.
func (c *Config) GetConnectionInfo() (bool, string) {
	if c.GetEnvironment() == "TEST" {
		return true, "file::memory:"
	}
	return false, fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true&multiStatements=true", c.dbUser, c.dbPassword, c.dbHost, c.dbPort, c.dbName)
}

// GetWebServerPort returns the web server port.
func (c *Config) GetWebServerPort() string {
	return c.webServerPort
}

// GetEnvironment returns the environment.
// The environment can be DEV, TEST or PROD.
func (c *Config) GetEnvironment() string {
	return strings.ToUpper(c.environment)
}
