package conf

import (
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// Config struct defines the configuration
type Config struct {
	Port      int64         `yaml:"port"`
	LogConfig LoggingConfig `yaml:"logconfig"`
	Token     string        `yaml:"token"`
	Movies    Media         `yaml:"movies"`
	Music     Media         `yaml:"music"`
	Tv        Media         `yaml:"tv"`
}

// TODO: Move somewhere else
type Media struct {
	Path string `yaml:"path"`
	API  string `yaml:"api"`
	Key  string `yaml:"key"`
	JSON string `yaml:"json"`
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
