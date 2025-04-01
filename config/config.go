package config

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"strings"

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
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")

	configDir := getConfigDir()
	viper.AddConfigPath(".")
	viper.AddConfigPath(configDir)

	viper.SetEnvPrefix("AZURE_TASK_GEN")
	viper.AutomaticEnv()

	setDefaults()

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			if err := setupInitialConfig(); err != nil {
				return fmt.Errorf("erro ao configurar configuração inicial: %w", err)
			}
		} else {
			return fmt.Errorf("erro ao ler arquivo de configuração: %w", err)
		}
	}

	if err := viper.Unmarshal(&AppConfig); err != nil {
		return fmt.Errorf("erro ao decodificar configuração: %w", err)
	}

	return nil
}

func getConfigDir() string {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "."
	}

	if runtime.GOOS == "windows" {
		return filepath.Join(homeDir, "AppData", "Local", "azure-task-gen")
	}
	return filepath.Join(homeDir, ".azure-task-gen")
}

func setupInitialConfig() error {
	fmt.Println("Configuração inicial necessária. Por favor, preencha as informações abaixo:")
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Digite sua chave da API do OpenAI: ")
	apiKey, _ := reader.ReadString('\n')
	apiKey = strings.TrimSpace(apiKey)

	fmt.Print("Digite o nome padrão do responsável: ")
	assignedTo, _ := reader.ReadString('\n')
	assignedTo = strings.TrimSpace(assignedTo)

	fmt.Print("Digite o email padrão do responsável: ")
	email, _ := reader.ReadString('\n')
	email = strings.TrimSpace(email)

	fmt.Print("Digite a atividade padrão (ex: Development): ")
	activity, _ := reader.ReadString('\n')
	activity = strings.TrimSpace(activity)

	fmt.Print("Digite o tipo de item de trabalho padrão (ex: Task): ")
	workItemType, _ := reader.ReadString('\n')
	workItemType = strings.TrimSpace(workItemType)

	fmt.Print("Digite a unidade de Story Point padrão (ex: 4): ")
	var defaultUST int
	fmt.Scanf("%d", &defaultUST)

	viper.Set("openai_api_key", apiKey)
	viper.Set("default_task.assigned_to", assignedTo)
	viper.Set("default_task.email", email)
	viper.Set("default_task.activity", activity)
	viper.Set("default_task.work_item_type", workItemType)
	viper.Set("default_task.default_ust", defaultUST)

	configDir := getConfigDir()
	if err := os.MkdirAll(configDir, 0755); err != nil {
		return fmt.Errorf("erro ao criar diretório de configuração: %w", err)
	}

	viper.SetConfigFile(filepath.Join(configDir, "config.yaml"))

	if err := viper.WriteConfig(); err != nil {
		return fmt.Errorf("erro ao salvar configuração: %w", err)
	}

	fmt.Println("Configuração inicial salva com sucesso!")
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
	configDir := getConfigDir()
	if err := os.MkdirAll(configDir, 0755); err != nil {
		return fmt.Errorf("erro ao criar diretório de configuração: %w", err)
	}

	viper.SetConfigFile(filepath.Join(configDir, "config.yaml"))

	if err := viper.WriteConfig(); err != nil {
		return fmt.Errorf("erro ao salvar configuração: %w", err)
	}

	return nil
}
