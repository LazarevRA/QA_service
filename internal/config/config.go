package config

import "os"

type Config struct {
	DBHost     string
	DBPort     string
	DBUser     string
	DBPassword string
	DBName     string
	ServerPort string
}

// Создает и возвращает объект Config с подтянутыми туда переменныеми окружения.
// Если переменные окружения не были найдены по ключам - использует стандартные.
func Load() *Config {
	return &Config{
		DBHost:     getEnv("DB_HOST", "localhost"),
		DBPort:     getEnv("DB_PORT", "5432"),
		DBUser:     getEnv("DB_USER", "postgres"),
		DBPassword: getEnv("DB_PASSWORD", "password"),
		DBName:     getEnv("DB_NAME", "qa_db"),
		ServerPort: getEnv("SERVER_PORT", "8080"),
	}
}

// Формирует строку DSN из данных конфигурации
func (c *Config) GetDSN() string {
	return "host=" + c.DBHost +
		" user=" + c.DBUser +
		" password=" + c.DBPassword +
		" dbname=" + c.DBName +
		" port+" + c.DBPort +
		" sslmode=disable TimeZone=UTC"
}

// Проверяет есть ли нужная переменная окружения по ключу,
// если такой переменной нет, возвращает default
func getEnv(key, defaultVal string) string {

	val := os.Getenv(key)

	if val != "" {
		return val
	}
	return defaultVal
}
