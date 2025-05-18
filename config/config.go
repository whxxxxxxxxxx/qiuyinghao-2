package config

import (
	"fmt"
	"os"
)

// Config 应用配置
type Config struct {
	DBHost     string
	DBPort     string
	DBUser     string
	DBPassword string
	DBName     string
	ServerPort string
}

// LoadConfig 加载配置
func LoadConfig() *Config {
	return &Config{
		DBHost:     getEnv("DB_HOST", "db4free.net"),
		DBPort:     getEnv("DB_PORT", "3306"),
		DBUser:     getEnv("DB_USER", "qiuyinghaouser"),
		DBPassword: getEnv("DB_PASSWORD", "qiuyinghaopassword"),
		DBName:     getEnv("DB_NAME", "qiuyinghao"),
		ServerPort: getEnv("SERVER_PORT", "8080"),
	}
}

// GetDSN 获取数据库连接字符串
func (c *Config) GetDSN() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		c.DBUser, c.DBPassword, c.DBHost, c.DBPort, c.DBName)
}

// 获取环境变量，如果不存在则返回默认值
func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}
