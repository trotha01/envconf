// Package envconf is for getting config values from the environment
package envconf

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

// ConfigFromEnv parses the given flagset and fills them with configs from the environment
func ConfigFromEnv(set *flag.FlagSet) error {
	var errs []string
	set.VisitAll(func(f *flag.Flag) {
		val := os.Getenv(f.Name)
		if val == "" {
			errs = append(errs, fmt.Sprintf("env variable %s missing", f.Name))
			return
		}
		if err := f.Value.Set(val); err != nil {
			errs = append(errs, fmt.Sprintf("failed to set flag %q with value %q - %s\n", f.Name, val, err))

		}
	})

	if len(errs) != 0 {
		return fmt.Errorf(strings.Join(errs, ", "))
	}
	return nil
}
