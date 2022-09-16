// Package postdude contains functions to run GET/POST/PUT/DELETE requests from terminal/cmd.
package postdude

import (
	"net/http"
	"os"

	"github.com/spf13/cobra"
)

// New creates a new command line program.
func New() *cobra.Command {
	config := &Config{
		Headers:            map[string][]string{},
		ResponseBodyOutput: os.Stdout,
		ControlOutput:      os.Stdout,
	}
	headers := make([]string, 0, 255)

	postdude := &cobra.Command{
		Use:     "pd URL",
		Short:   "HTTP client within your terminal",
		Long:    "Postdude is an HTTP client that helps you start an HTTP requests's within your terminal.",
		Args:    ValidateArg(config),
		PreRunE: ValidateOption(config, headers),
		RunE: func(cmd *cobra.Command, args []string) error {
			return Run(config)
		},
	}

	// Set flags.
	postdude.PersistentFlags().StringSliceVarP(&headers, "headers", "e", nil, `Setup custom headers to be sent with the request.\nSeparated by ','.\n Example: HeaderName: TheContent, OtherHeader: TheContent`)
	postdude.PersistentFlags().StringVarP(&config.UserAgent, "user-agent", "u", "pd", "Setup the user-agent to be used for the request.")
	postdude.PersistentFlags().StringVarP(&config.Data, "data", "d", "", "The data to be sent with the request.")
	postdude.PersistentFlags().StringVarP(&config.Method, "method", "m", http.MethodGet, "Set the HTTP method, default to GET.")
	postdude.PersistentFlags().BoolVarP(&config.Insecure, "insecure", "i", false, "Set secure connections over HTTPS.")

	return postdude
}
