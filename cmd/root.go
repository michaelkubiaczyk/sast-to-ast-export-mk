package cmd

import (
	"os"

	"github.com/checkmarxDev/ast-observability-library/pkg/aol"
	"github.com/checkmarxDev/ast-sast-export/internal"
	"github.com/checkmarxDev/ast-sast-export/internal/export"
	"github.com/checkmarxDev/ast-sast-export/internal/logging"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

const (
	userArg                      = "user"
	passArg                      = "pass"
	urlArg                       = "url"
	exportArg                    = "export"
	resultsProjectActiveSinceArg = "results-project-active-since"
	debugArg                     = "debug"
	verboseArg                   = "verbose"

	resultsProjectActiveSinceDefaultValue = 180
)

// productName is defined in Makefile and initialized during build
var productName string

// productVersion is defined in VERSION and initialized during build
var productVersion string

// productBuild is defined in Makefile and initialized during build
var productBuild string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   productName,
	Short: "Exports SAST data for importing in AST",
	Long: `Exports encrypted SAST data for importing in AST. Example usage:

cxsast_exporter --user username --pass password --url http://localhost

Produces a zip file containing the encrypted SAST data, e.g. cxsast_exporter-2021-09-10-15-42-35.zip
Also produces a log file with diagnostic information, e.g. cxsast_exporter-2021-09-10-15-42-35.log`,
	Run: func(cmd *cobra.Command, args []string) {
		// setup logging
		verbose, flagErr := cmd.Flags().GetBool(verboseArg)
		if flagErr != nil {
			panic(flagErr)
		}

		logFileWriter, err := logging.NewFileWriter(productName)
		if err != nil {
			panic(err)
		}
		defer logFileWriter.Close()

		levelWriter := logging.NewMultiLevelWriter(verbose, zerolog.InfoLevel, aol.GetNewConsoleWriter(), logFileWriter)

		aolErr := aol.Init(aol.InitOptions{
			ServiceName:       productName,
			LogLevel:          zerolog.LevelTraceValue,
			LogOutputStream:   &levelWriter,
			Version:           "",
			TelemetryEndpoint: "",
		})
		if aolErr != nil {
			panic(aolErr)
		}

		defer func() {
			if r := recover(); r != nil {
				log.Error().Msgf("execution failed: %v", r)
				os.Exit(1)
			}
		}()

		// start export
		allArgs := GetArgs(cmd, productName)
		internal.RunExport(&allArgs)
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

//nolint:gochecknoinits
func init() {
	resultsProjectActiveSinceUsage := "include only results from projects active in the last N days"

	rootCmd.Flags().StringP(userArg, "", "", "SAST username")
	rootCmd.Flags().StringP(passArg, "", "", "SAST password")
	rootCmd.Flags().StringP(urlArg, "", "", "SAST url")
	rootCmd.Flags().StringSliceP(exportArg, "", export.GetOptions(), "SAST export options")
	rootCmd.Flags().IntP(resultsProjectActiveSinceArg, "", resultsProjectActiveSinceDefaultValue, resultsProjectActiveSinceUsage)
	rootCmd.Flags().Bool(debugArg, false, "activate debug mode")
	rootCmd.Flags().BoolP(verboseArg, "v", false, "enable verbose logging to console")

	if err := rootCmd.MarkFlagRequired(userArg); err != nil {
		panic(err)
	}
	if err := rootCmd.MarkFlagRequired(passArg); err != nil {
		panic(err)
	}
	if err := rootCmd.MarkFlagRequired(urlArg); err != nil {
		panic(err)
	}
	if err := rootCmd.MarkFlagCustom(resultsProjectActiveSinceArg, resultsProjectActiveSinceUsage); err != nil {
		panic(err)
	}
}
