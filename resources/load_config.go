package resources

import (
	"github.com/spf13/viper"
)

// Config contains all configurations application
type Config struct {
	DBUserMongo               string `mapstructure:"DB_USER_MONGO"`
	DBNameDataBaseMongo       string `mapstructure:"DB_NAME_DATABASE_MONGO"`
	URLAPIVerifyUserBlocked   string `mapstructure:"URL_API_VERIFY_USER_BLOCKED"`
	VersionApp                string `mapstructure:"VERSION_APP"`
	DBHostMongo               string `mapstructure:"DB_HOST_MONGO"`
	DBPortMongo               string `mapstructure:"DB_PORT_MONGO"`
	APPPort                   string `mapstructure:"APP_PORT"`
	DBPassMongo               string `mapstructure:"DB_PASSWORD_MONGO"`
	AppEnv                    string `mapstructure:"APP_ENV"`
	JWTSecretKey              string `mapstructure:"JWT_SECRET_KEY"`
	JWTExpiredTokenTimeMinute int    `mapstructure:"JWT_EXPIRED_TOKEN_TIME_MINUTE"`
	DBTimeOutMongo            int    `mapstructure:"DB_TIME_OUT_MONGO"`
	ShowLogRequest            bool   `mapstructure:"SHOW_LOGS_REQUEST"`
	ShowMonitor               bool   `mapstructure:"SHOW_MONITOR"`
	LoggerSaveFile            bool   `mapstructure:"LOGGER_SAVE_FILE"`
}

// configuration get unique instance of Config
var configuration *Config

// ConfigurationEnv contain all configuration environment
func ConfigurationEnv() *Config {
	return configuration
}

// LoadConfig build all configuration application
func LoadConfig(pathsProperties string) {
	configuration = &Config{}
	viper.AutomaticEnv()
	viper.SetConfigFile(pathsProperties)

	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}

	if err := viper.Unmarshal(configuration); err != nil {
		panic(err)
	}
}
