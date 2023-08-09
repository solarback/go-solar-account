package initializers

import (
	"gopkg.in/yaml.v3"
	"log"
	"os"
)

type EnvironmentConfig struct {
	Environment string    `yaml:"environment"`
	ServerPort  int       `yaml:"serverPort"`
	Database    DbConfig  `yaml:"database"`
	JWT         JWTConfig `yaml:"jwt"`
}

type JWTConfig struct {
	JWTSecret    string `yaml:"jwt-secret"`
	JWTIssuer    string `yaml:"jwt-issuer"`
	JWTAudience  string `yaml:"jwt-audience"`
	CookieDomain string `yaml:"cookie-domain"`
}

type DbConfig struct {
	DBHost         string `yaml:"host"`
	DBUserName     string `yaml:"db-user-name"`
	DBUserPassword string `yaml:"db-user-password"`
	DBName         string `yaml:"db-name"`
	DBPort         string `yaml:"db-port"`
}

type Config struct {
	Config []EnvironmentConfig `yaml:"config"`
}

func LoadConfig(path string) (config EnvironmentConfig) {

	data, err := os.ReadFile(path)

	if err != nil {
		log.Fatalf("Error during reading file %v", err)
	}

	var configs Config
	err = yaml.Unmarshal(data, &configs)
	if err != nil {
		log.Fatalf("Error during parsing file %v", err)
	}

	return configs.Config[0]
}
