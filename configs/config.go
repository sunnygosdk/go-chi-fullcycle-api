package configs

import (
	"fmt"
	"strings"

	"github.com/go-chi/jwtauth"
	"github.com/spf13/viper"
)

type conf struct {
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

var config *conf

// init config.
// It will read the .env file and set the config.
func init() {
	viper.SetConfigFile(".env")
	viper.AddConfigPath(".")
	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("config file read error: %w", err))
	}

	required := []string{"DB_PASSWORD", "JWT_SECRET"}
	for _, key := range required {
		if !viper.IsSet(key) || viper.GetString(key) == "" {
			panic(fmt.Errorf("missing required config: %s", key))
		}
	}

	config = &conf{
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

// NewConfig returns the config.
func NewConfig() *conf {
	return config
}

// GetEnvironment returns the environment.
// The environment can be DEV or PROD.
func (c *conf) GetEnvironment() string {
	return strings.ToUpper(c.environment)
}
