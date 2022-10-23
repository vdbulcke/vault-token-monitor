package cmd

import (
	"errors"
	"log"
	"os"

	"github.com/spf13/cobra"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"

	"github.com/vdbulcke/vault-token-monitor/internal"
	"github.com/vdbulcke/vault-token-monitor/internal/server"
)

// args var
var configFilename string

func init() {
	// bind to root command
	rootCmd.AddCommand(startCmd)
	// add flags to sub command
	startCmd.Flags().StringVarP(&configFilename, "config", "c", "", "server config file")

	// required flags
	//nolint
	startCmd.MarkFlagRequired("config")

}

var startCmd = &cobra.Command{
	Use:   "server",
	Short: "Starts the monitoring server",
	// Long: "",
	Run: runServer,
}

// runServer cobra server handler
func runServer(cmd *cobra.Command, args []string) {

	// Zap Logger
	var logger *zap.Logger
	var err error
	if Debug {
		logger, err = zap.NewDevelopment()
		if err != nil {
			log.Fatal(err)
		}
	} else {
		// override time format
		zapConfig := zap.NewProductionEncoderConfig()
		zapConfig.EncodeTime = zapcore.ISO8601TimeEncoder
		consoleEncoder := zapcore.NewJSONEncoder(zapConfig)

		// default writer for logger
		consoleDebugging := zapcore.Lock(os.Stdout)
		consoleErrors := zapcore.Lock(os.Stderr)

		// set log level to writer
		core := zapcore.NewTee(
			// zapcore.NewCore(consoleEncoder, consoleDebugging, zap.DebugLevel),
			zapcore.NewCore(consoleEncoder, consoleDebugging, zap.InfoLevel),
			zapcore.NewCore(consoleEncoder, consoleErrors, zap.WarnLevel),
			zapcore.NewCore(consoleEncoder, consoleErrors, zap.ErrorLevel),
		)

		// add function caller and stack trace on error
		logger = zap.New(core, zap.AddCaller(), zap.AddStacktrace(zapcore.ErrorLevel))
	}
	//nolint
	defer logger.Sync()

	// parse config
	config, err := internal.ParseConfig(configFilename)
	if err != nil {
		logger.Error("parsing config", zap.Error(err))
		os.Exit(1)
	}

	// validate config
	if !internal.ValidateConfig(config) {
		logger.Error("validating config", zap.Error(errors.New("Validation Error")))
		os.Exit(1)
	}

	// Create new Sever
	s, err := server.NewVaultMonitorServer(logger, config)
	if err != nil {
		logger.Error("Error create Vault Monitoring Server", zap.Error(err))
		os.Exit(1)
	}

	// start server
	err = s.StartServer()
	if err != nil {
		logger.Error("Error Starting Vault Monitoring Server", zap.Error(err))
		os.Exit(1)
	}

}
