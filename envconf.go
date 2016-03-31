// Package envconf is for getting config values from the environment
package envconf

import (
	"flag"
	"fmt"
	"os"
)

// ConfigFromEnv parses the given flagset and fills them with configs from the environment
func ConfigFromEnv(set *flag.FlagSet) error {
	var err error
	set.VisitAll(func(f *flag.Flag) {
		if err != nil {
			return
		}

		val := os.Getenv(f.Name)
		if val == "" {
			err = fmt.Errorf("env variable %q missing", f.Name)
			return
		}
		if ferr := f.Value.Set(val); ferr != nil {
			err = fmt.Errorf("failed to set flag %q with value %q", f.Name, val)
		}
	})

	return err
}
