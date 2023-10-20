package configSys

import "github.com/BurntSushi/toml"

func Load(path string) (Config, error) {
	var config Config
	_, err := toml.DecodeFile(path, &config)

	if err != nil {
		return Config{}, err
	}

	return config, err
}
