package main

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	// logLevel contains the value of the log_level flag.
	logLevel string

	// logMode contains the value of the log_mode flag.
	logMode string
)

func rootCmd() *cobra.Command {
	var rootCmd = &cobra.Command{
		Use:   "jumpshot",
		Short: "JumpShot V1",
		PersistentPreRunE: func(_ *cobra.Command, _ []string) error {
			logger, err := parseLogMode(logMode, logLevel)
			if err != nil {
				return err
			}
			zap.ReplaceGlobals(logger)
			return nil
		},
	}

	rootCmd.PersistentFlags().StringVar(&logLevel, "log_level", "", "Zap log level. Empty means default.")
	rootCmd.PersistentFlags().StringVar(&logMode, "log_mode", "dev", "Zap log mode. 'dev' and 'prod' are supported.")

	rootCmd.AddCommand(serveCmd())
	return rootCmd
}

// Parse log level.
func parseLogLevel(level string) (*zapcore.Level, error) {
	if level == "" {
		return nil, fmt.Errorf("empty log level")
	}

	var l zapcore.Level
	if err := l.Set(level); err != nil {
		return nil, fmt.Errorf("error parsing log level %q: %v", level, err)
	}

	return &l, nil
}

// Create a new logger by the flags.
func parseLogMode(mode string, level string) (*zap.Logger, error) {
	var options []zap.Option
	if l, err := parseLogLevel(level); err == nil {
		options = append(options, zap.AddStacktrace(l))
	}

	switch strings.ToLower(mode) {
	case "prod":
		return zap.NewProduction(options...)
	case "dev":
		return zap.NewDevelopment(options...)
	default:
		return nil, fmt.Errorf("invalid log mode: %v", mode)
	}
}
