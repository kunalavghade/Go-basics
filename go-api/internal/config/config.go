package config

import (
	"os"
	"strconv"
	"time"

	"github.com/joho/godotenv"
)

type Config struct {
	Server   ServerConfig
	Database DatabaseConfig
	JWT      JWTConfig
	AWS      AWSConfig
	Upload   UploadConfig
	Logger   LoggerConfig
}

type ServerConfig struct {
	Port    string
	GinMode string
}

type DatabaseConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	Name     string
	SSLMode  string
}

type JWTConfig struct {
	Secret                string
	AccessTokenExpiresIn  time.Duration
	RefreshTokenExpiresIn time.Duration
}

type LoggerConfig struct {
	Level string
}

type AWSConfig struct {
	Region          string
	AccessKeyID     string
	SecretAccessKey string
	S3Bucket        string
	S3Endpoint      string
}

type UploadConfig struct {
	Path        string
	MaxFileSize int64 // bytes
}

func Load() (*Config, error) {
	_ = godotenv.Load()
	jwtEx, _ := time.ParseDuration(getEnv("JWT_TOKEN_EXPIRES_IN", "1h"))
	jwtRe, _ := time.ParseDuration(getEnv("JWT_REFRESH_TOKEN_EXPIRES_IN", "168h"))
	uploadPath := getEnv("UPLOAD_PATH", "./uploads")
	maxSize, _ := strconv.ParseInt(getEnv("UPLOAD_MAX_FILE_SIZE", "5242880"), 10, 64)
	return &Config{
		Server: ServerConfig{
			Port:    getEnv("SERVER_PORT", "8080"),
			GinMode: getEnv("GIN_MODE", "release"),
		},
		Database: DatabaseConfig{
			Host:     getEnv("DB_HOST", "localhost"),
			Port:     getEnv("DB_PORT", "5432"),
			User:     getEnv("DB_USER", "postgres"),
			Password: getEnv("DB_PASSWORD", "postgres"),
			Name:     getEnv("DB_NAME", "postgres"),
			SSLMode:  getEnv("DB_SSL_MODE", "disable"),
		},
		JWT: JWTConfig{
			Secret:                getEnv("JWT_SECRET", ""),
			AccessTokenExpiresIn:  jwtEx,
			RefreshTokenExpiresIn: jwtRe,
		},
		AWS: AWSConfig{
			Region:          getEnv("AWS_REGION", ""),
			AccessKeyID:     getEnv("AWS_ACCESS_KEY_ID", ""),
			SecretAccessKey: getEnv("AWS_SECRET_ACCESS_KEY", ""),
			S3Bucket:        getEnv("AWS_S3_BUCKET", ""),
			S3Endpoint:      getEnv("AWS_S3_ENDPOINT", ""),
		},
		Upload: UploadConfig{
			Path:        uploadPath,
			MaxFileSize: maxSize,
		},
		Logger: LoggerConfig{
			Level: getEnv("LOG_LEVEL", "info"),
		},
	}, nil
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
