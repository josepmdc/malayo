package cmd

import (
	"fmt"
	"log"
	"malayo/api"
	"malayo/conf"
	postgres "malayo/database"
	"malayo/indexing"
	"malayo/json"
	"malayo/services"
	"net/http"
	"time"

	"github.com/spf13/cobra"
)

var index bool

// RootCommand is the main command for starting the server. It reads the command flags port and config
func RootCommand() *cobra.Command {
	rootCmd := cobra.Command{
		Use: "[options]",
		Run: run,
	}

	rootCmd.Flags().IntP("port", "p", 0, "--port [number] | -p [number]")
	rootCmd.PersistentFlags().StringP("config", "c", "config.yaml", "--config [path] | -c [path]")
	rootCmd.PersistentFlags().StringP("mediaPath", "m", "$HOME/Videos", "--mediaPath [path] | -m [path]")
	rootCmd.Flags().BoolVarP(&index, "index", "i", false, "--index=[true|false] | -i=[true|false]")

	return &rootCmd
}

func run(cmd *cobra.Command, _ []string) {
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
		err := indexing.IndexMediaLibrary(&config.Media)
		if err != nil {
			logger.Errorf("Unable to index library. Error: \n%s", err.Error())
		}
	}

	mediaService := createMediaService(config)

	startServer(config, mediaService)
}

func createMediaService(config *conf.Config) services.MediaService {
	mediaService := services.NewMediaService()
	switch config.Storage {
	case "json":
		mediaService.MovieRepository = json.NewMovieRepo(config.Media.Movies)
		mediaService.TvRepository = json.NewTvRepo(config.Media.Tv)
	case "postgres":
		// TODO Implement Postgres Storage
		mediaService.MovieRepository = postgres.NewMovieRepo(config.Media.Movies)
		//mediaService.TvRepository = postgres.NewTvRepo(config.Media.Tv)
	}
	return mediaService
}

func startServer(config *conf.Config, service services.MediaService) {

	handler := api.NewRouter(config, service)

	srv := &http.Server{
		Addr:         fmt.Sprintf("%s:%d", "localhost", config.Port),
		Handler:      handler,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  120 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}
