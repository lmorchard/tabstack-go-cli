package cmd

import (
	"fmt"
	"os"

	"github.com/lmorchard/tabstack-go-cli/internal/config"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	cfgFile string
	log     = logrus.New()
	cfg     *config.Config
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "tabstack",
	Short: "A brief description of your application",
	Long: `A longer description of what your application does and how it works.

This can be multiple lines and should provide helpful context about the
purpose and usage of your CLI tool.`,
	// API errors and other runtime failures shouldn't drag the full Cobra
	// usage block along with them — Cobra still prints the error itself.
	SilenceUsage: true,
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		initConfig()
		setupLogging()
	},
}

// Execute adds all child commands to the root command and sets appropriate flags.
func Execute() {
	// Cobra prints "Error: ..." to stderr itself when SilenceErrors is false
	// (the default). All we do here is propagate the non-zero exit.
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

func init() {
	// Configuration file flag
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is ./tabstack.yaml)")

	// Logging flags
	rootCmd.PersistentFlags().BoolP("verbose", "v", false, "verbose output")
	rootCmd.PersistentFlags().Bool("debug", false, "debug output")
	rootCmd.PersistentFlags().Bool("log-json", false, "output logs in JSON format")

	// Tabstack API flags
	rootCmd.PersistentFlags().String("api-key", "", "Tabstack API key (env: TABSTACK_API_KEY)")
	rootCmd.PersistentFlags().String("base-url", "", "Tabstack API base URL (env: TABSTACK_BASE_URL)")

	// Bind flags to viper
	_ = viper.BindPFlag("verbose", rootCmd.PersistentFlags().Lookup("verbose"))
	_ = viper.BindPFlag("debug", rootCmd.PersistentFlags().Lookup("debug"))
	_ = viper.BindPFlag("log_json", rootCmd.PersistentFlags().Lookup("log-json"))
	_ = viper.BindPFlag("api_key", rootCmd.PersistentFlags().Lookup("api-key"))
	_ = viper.BindPFlag("base_url", rootCmd.PersistentFlags().Lookup("base-url"))
	_ = viper.BindEnv("api_key", "TABSTACK_API_KEY")
	_ = viper.BindEnv("base_url", "TABSTACK_BASE_URL")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag
		viper.SetConfigFile(cfgFile)
	} else {
		// Search for config in current directory
		viper.AddConfigPath(".")
		viper.SetConfigType("yaml")
		viper.SetConfigName("tabstack")
	}

	// Set defaults
	viper.SetDefault("verbose", false)
	viper.SetDefault("debug", false)
	viper.SetDefault("log_json", false)

	// Read in environment variables that match
	viper.AutomaticEnv()

	// If a config file is found, read it in
	if err := viper.ReadInConfig(); err != nil {
		if cfgFile != "" {
			// Only error if config was explicitly specified
			fmt.Fprintf(os.Stderr, "Error reading config file: %v\n", err)
			os.Exit(1)
		}
	}
}

// setupLogging configures the logger based on configuration
func setupLogging() {
	if viper.GetBool("log_json") {
		log.SetFormatter(&logrus.JSONFormatter{})
	} else {
		log.SetFormatter(&logrus.TextFormatter{
			FullTimestamp: true,
		})
	}

	if viper.GetBool("debug") {
		log.SetLevel(logrus.DebugLevel)
	} else if viper.GetBool("verbose") {
		log.SetLevel(logrus.InfoLevel)
	} else {
		log.SetLevel(logrus.WarnLevel)
	}
}

// GetConfig returns the application configuration, loading it if necessary
func GetConfig() *config.Config {
	if cfg == nil {
		cfg = &config.Config{
			Verbose: viper.GetBool("verbose"),
			Debug:   viper.GetBool("debug"),
			LogJSON: viper.GetBool("log_json"),
			APIKey:  viper.GetString("api_key"),
			BaseURL: viper.GetString("base_url"),
		}
	}
	return cfg
}

// GetLogger returns the configured logger
func GetLogger() *logrus.Logger {
	return log
}
