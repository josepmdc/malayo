package conf

import (
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// Config struct defines the configuration for the entire application
type Config struct {
	Port           int64          `yaml:"port"`
	Domain         string         `yaml:"domain"`
	LogConfig      LoggingConfig  `yaml:"logconfig"`
	Token          string         `yaml:"token"`
	Media          Media          `yaml:"media"`
	Storage        string         `yaml:"storage"`
	PostgresConfig PostgresConfig `yaml:"db"`
}

// LoadConfig takes a command as an argument to get the command flags
// in case the user specified special settings. Then it loads the config from specified file
func LoadConfig(cmd *cobra.Command) (*Config, error) {

	if err := viper.BindPFlags(cmd.Flags()); err != nil {
		return nil, err
	}

	viper.SetEnvPrefix("MALAYO")
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.AutomaticEnv()

	if configFile, _ := cmd.Flags().GetString("config"); configFile != "" {
		viper.SetConfigFile(configFile)
	} else {
		viper.SetConfigName("config")
		viper.AddConfigPath("./")
	}

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	config := new(Config)
	if err := viper.Unmarshal(config); err != nil {
		return nil, err
	}

	return config, nil
}
