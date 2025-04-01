package config

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/viper"
)

type Config struct {
	OpenAIAPIKey string `mapstructure:"openai_api_key"`
	DefaultTask  struct {
		AssignedTo   string `mapstructure:"assigned_to"`
		Email        string `mapstructure:"email"`
		Activity     string `mapstructure:"activity"`
		WorkItemType string `mapstructure:"work_item_type"`
		DefaultUST   int    `mapstructure:"default_ust"`
	} `mapstructure:"default_task"`
}

var AppConfig Config

func Init() error {
	// Set up Viper
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")

	// Add config paths
	viper.AddConfigPath(".")
	viper.AddConfigPath("$HOME/.azure-task-gen")

	// Set environment variables
	viper.SetEnvPrefix("AZURE_TASK_GEN")
	viper.AutomaticEnv()

	// Set defaults
	setDefaults()

	// Read config file
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			return fmt.Errorf("erro ao ler arquivo de configuração: %w", err)
		}
	}

	// Unmarshal config
	if err := viper.Unmarshal(&AppConfig); err != nil {
		return fmt.Errorf("erro ao decodificar configuração: %w", err)
	}

	return nil
}

func setDefaults() {
	viper.SetDefault("openai_api_key", os.Getenv("OPENAI_API_KEY"))
	viper.SetDefault("default_task.assigned_to", "Ygor Azambuja")
	viper.SetDefault("default_task.activity", "Development")
	viper.SetDefault("default_task.work_item_type", "Task")
	viper.SetDefault("default_task.default_ust", 4)
}

func SaveConfig() error {
	// Create config directory if it doesn't exist
	configDir := filepath.Join(os.Getenv("HOME"), ".azure-task-gen")
	if err := os.MkdirAll(configDir, 0755); err != nil {
		return fmt.Errorf("erro ao criar diretório de configuração: %w", err)
	}

	// Set config file path
	viper.SetConfigFile(filepath.Join(configDir, "config.yaml"))

	// Write config
	if err := viper.WriteConfig(); err != nil {
		return fmt.Errorf("erro ao salvar configuração: %w", err)
	}

	return nil
}
