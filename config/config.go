package config

import (
	"fmt"
	"os"
	"strings"

	"github.com/joho/godotenv"
)

type DBConfig struct {
	DBType string
	DBPort string
	DBHost string
	DBUser string
	DBPass string
	DBFile string
}

type APIConfig struct {
	ServerPort  string
	ServerHost  string
	CORSOrigins []string
	CORSMethods []string
}

type AuthConfig struct {
	JWTSecretKey string
}

type EmailConfig struct {
	MailHost     string
	MailPort     string
	MailUsername string
	MailPassword string
}

type DebugConfig struct {
	DebugMode bool
	LogOutputPath string
}

// Add more configs as needed
type Config struct {
	DBConfig    DBConfig
	APIConfig   APIConfig
	AuthConfig  AuthConfig
	EmailConfig EmailConfig
	DebugConfig DebugConfig
}

func LoadConfig() (*Config, error) {
	if os.Getenv("ENV") == "DEV" {
		err := godotenv.Load(".env.development")
		if err != nil {
			return nil, fmt.Errorf("error loading .env.development file")
		}
 	} else {
		err := godotenv.Load(".env.production")
		if err != nil {
			return nil, fmt.Errorf("error loading .env.production file")
		}
	}

	parseCSV := func(envVar string) []string {
        return strings.Split(os.Getenv(envVar), ",")
    }

	config := &Config{
        DBConfig: DBConfig{
            DBType: os.Getenv("DB_TYPE"),
            DBPort: os.Getenv("DB_PORT"),
            DBHost: os.Getenv("DB_HOST"),
            DBUser: os.Getenv("DB_USER"),
			DBFile: os.Getenv("DB_File"),
            DBPass: os.Getenv("DB_PASSWORD"),
        },
        APIConfig: APIConfig{
            ServerPort:  os.Getenv("SERVER_PORT"),
            ServerHost:  os.Getenv("SERVER_HOST"),
			CORSOrigins: parseCSV("CORS_ORIGINS"),  
            CORSMethods: parseCSV("CORS_METHODS"), 
        },
        AuthConfig: AuthConfig{
            JWTSecretKey: os.Getenv("JWT_SECRET_KEY"),
        },
        EmailConfig: EmailConfig{
            MailHost:     os.Getenv("MAIL_HOST"),
            MailPort:     os.Getenv("MAIL_PORT"),
            MailUsername: os.Getenv("MAIL_USERNAME"),
            MailPassword: os.Getenv("MAIL_PASSWORD"),
        },
        DebugConfig: DebugConfig{
            DebugMode: os.Getenv("DEBUG_MODE") == "true",
			LogOutputPath: os.Getenv("LOG_OUTPUT_PATH"),
        },
    }

    return config, nil
}

func GetEnv(key string, fallback ...string) (string, error) {
    value, exists := os.LookupEnv(key)
    if !exists {
        if len(fallback) > 0 {
            return fallback[0], nil
        }
        return "", fmt.Errorf("environment variable %s is not set", key)
    }
    return value, nil
}