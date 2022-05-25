# cfg

Adds a Go module for loading configurations from YAML or Java properties files,
dotenv files, and environment variables. Returns a \*koanf.Koaf instance from
https://github.com/knadh/koanf.

You can load either a YAML or Java properties file per-instance, along with some
combination of dotenv files and environment variables.

The default configuration file path is `/etc/cyverse/de/configs/service.yml`.

The default dotenv file path is `/etc/cyverse/de/env/service.env`.

The default environment variable prefix is `DISCOENV_`.

The default file type is YAML and the strict merge is not enabled by default.

# Example

```Go
    import (
        "github.com/cyverse-de/go-mod/cfg"
    )

    var (
        configPath    = flag.String("config", cfg.DefaultConfigPath, "The path to the config file")
		dotEnvPath    = flag.String("dotenv-path", cfg.DefaultDotEnvPath, "The path to the env file to load")
        envPrefix     = flag.String("env-prefix", cfg.DefaultEnvPrefix, "The prefix to look for when setting configuration setting in environment variables")
    )

	k, err := cfg.Init(&cfg.Settings{
		EnvPrefix:   *envPrefix,
		ConfigPath:  *configPath,
		DotEnvPath:  *dotEnvPath,
		StrictMerge: false,
		FileType:    cfg.YAML,
	})
	if err != nil {
		log.Fatal(err)
	}

	dbURI := k.String("db.uri")
	if dbURI == "" {
		log.Fatal("db.uri must be set in the configuration file")
	}
```
