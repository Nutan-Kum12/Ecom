package config

import (
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	DBUser                 string
	DBPassword             string
	DBName                 string
	DBAddress              string
	PublicHost             string
	Port                   string
	JWTExpirationInSeconds int64
	JWTSecretKey           string
}

var Envs = initConfig()

func initConfig() Config {
	godotenv.Load()
	return Config{
		PublicHost:             getEnv("PUBLIC_HOST", "http://localhost"),
		Port:                   getEnv("PORT", "8080"),
		DBUser:                 getEnv("DB_USER", "root"),
		DBPassword:             getEnv("DB_PASSWORD", ""),
		DBName:                 getEnv("DB_NAME", ""),
		DBAddress:              fmt.Sprintf("%s:%s", getEnv("DB_HOST", "127.0.0.1"), getEnv("DB_PORT", "3306")),
		JWTExpirationInSeconds: getEnvAsInt("JWT_EXPIRATION", 3600*24*7),
		JWTSecretKey:           getEnv("JWT_SECRET", "not-secrer-secret-anymore?"),
	}
}
func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
func getEnvAsInt(name string, fallback int64) int64 {
	if value, ok := os.LookupEnv(name); ok {
		i, err := strconv.ParseInt(value, 10, 64)
		if err != nil {
			return fallback
		}
		return i
	}
	return fallback
}
