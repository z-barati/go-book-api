package config

import (
	"github.com/spf13/viper"
)

// Load initializes the configuration from environment variables and config files
func Load() error {
	// Set default values
	setDefaults()

	// Read from environment variables
	viper.AutomaticEnv()

	// Read from TOML config file if it exists
	viper.SetConfigName("config")
	viper.SetConfigType("toml")
	viper.AddConfigPath(".")
	viper.AddConfigPath("./config")

	// Try to read TOML config file
	if err := viper.ReadInConfig(); err != nil {
		// If TOML file doesn't exist, try .env file
		viper.SetConfigName(".env")
		viper.SetConfigType("env")
		viper.AddConfigPath(".")
		_ = viper.ReadInConfig()
	}

	return nil
}

// setDefaults sets default configuration values
func setDefaults() {
	// Server defaults
	viper.SetDefault("server.port", "8080")
	viper.SetDefault("server.host", "localhost")

	// Database defaults
	viper.SetDefault("database.host", "localhost")
	viper.SetDefault("database.port", "3306")
	viper.SetDefault("database.name", "book_db")
	viper.SetDefault("database.user", "root")
	viper.SetDefault("database.password", "password")
	viper.SetDefault("database.charset", "utf8mb4")
	viper.SetDefault("database.parse_time", true)
	viper.SetDefault("database.loc", "Local")

	// JWT defaults
	viper.SetDefault("jwt.secret", "your-secret-key")
	viper.SetDefault("jwt.expiration", 24) // hours

	// App defaults
	viper.SetDefault("app.mode", "development")
	viper.SetDefault("app.name", "Go Book API")
	viper.SetDefault("app.version", "1.0.0")
}

// GetDatabaseDSN returns the database connection string
func GetDatabaseDSN() string {
	host := viper.GetString("database.host")
	port := viper.GetString("database.port")
	name := viper.GetString("database.name")
	user := viper.GetString("database.user")
	password := viper.GetString("database.password")
	charset := viper.GetString("database.charset")
	parseTime := viper.GetBool("database.parse_time")
	loc := viper.GetString("database.loc")

	dsn := user + ":" + password + "@tcp(" + host + ":" + port + ")/" + name + "?charset=" + charset
	if parseTime {
		dsn += "&parseTime=true"
	}
	if loc != "" {
		dsn += "&loc=" + loc
	}

	return dsn
}

// GetJWTSecret returns the JWT secret key
func GetJWTSecret() string {
	return viper.GetString("jwt.secret")
}

// GetJWTExpiration returns the JWT expiration time in hours
func GetJWTExpiration() int {
	return viper.GetInt("jwt.expiration")
}
