package cmd

import (
	"fmt"
	"log"
	"malayo/api"
	"malayo/conf"
	"malayo/indexing"
	"net/http"
	"time"

	"github.com/spf13/cobra"
)

var index bool

// RootCommand is the main command for starting the server. It reads the command flags port and config
func RootCommand() *cobra.Command {
	rootCmd := cobra.Command{
		Use: "Start the server",
		Run: run,
	}

	rootCmd.Flags().IntP("port", "p", 0, "the port to use")
	rootCmd.PersistentFlags().StringP("config", "c", "config.yaml", "config file")
	rootCmd.PersistentFlags().StringP("mediaPath", "m", "$HOME/Videos", "media directory")
	rootCmd.Flags().BoolVarP(&index, "index", "i", false, "indexes your media library")

	return &rootCmd
}

func run(cmd *cobra.Command, args []string) {
	config, err := conf.LoadConfig(cmd)
	if err != nil {
		log.Fatal("Failed to load config: " + err.Error())
	}

	logger, err := conf.ConfigureLogging(&config.LogConfig)
	if err != nil {
		log.Fatal("Failed to configure logging: " + err.Error())
	}

	logger.Infof("Starting with config: %+v", config)

	if index == true {
		indexing.IndexMediaLibrary(config.MediaPath)
	}

	startServer(config)
}

func startServer(config *conf.Config) {

	handler := api.NewRouter(config)

	srv := &http.Server{
		Addr:         fmt.Sprintf("%s:%d", "localhost", config.Port),
		Handler:      handler,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  120 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}
