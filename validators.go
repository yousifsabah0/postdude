package postdude

import (
	"errors"
	"net/url"
	"strings"

	"github.com/spf13/cobra"
)

func ValidateArg(config *Config) func(cmd *cobra.Command, args []string) error {
	return func(cmd *cobra.Command, args []string) error {
		if len(args) != 1 {
			return errors.New("only single URL needed to be called")
		}

		u, err := url.Parse(args[0])
		if err != nil {
			return err
		}

		config.URL = *u
		return nil
	}
}

func ValidateOption(config *Config, headers []string) func(cmd *cobra.Command, args []string) error {
	return func(cmd *cobra.Command, args []string) error {
		for _, header := range headers {
			if name, value, found := strings.Cut(header, ":"); found {
				config.Headers.Add(name, value)
			} else {
				return errors.New("not valid header separator, headers should be separated by ','")
			}
		}
		return nil
	}
}
