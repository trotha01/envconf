# EnvConf

This is a minimalist library for getting a config from the environment

# Example
```
set := flag.NewFlagSet("service1", flag.ExitOnError)

commHost := set.String("SERVICE1_HOST", "", "host to use for server")
commPort := set.String("SERVICE1_PORT", "", "port to use for server")
incomingSocket := set.String("SERVICE1_LOG_FILE", "", "path to log file")

err := envconf.ConfigFromEnv(set)
if err != nil {
    log.Fatal("error getting config: %s", err.Error())
}
```
