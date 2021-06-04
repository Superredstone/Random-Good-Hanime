package RandomGoodHanime

import (
	"os"

	"gopkg.in/yaml.v3"
)

//Read config file in the start directory of the application
func ReadConfig() (Config, error) {
	if _, err := os.Stat("config.yml"); os.IsNotExist(err) {
		f, err := os.Create("config.yml")
		if err != nil {
			return Config{}, err
		}

		f.WriteString(StandardConfig)
	}

	f, err := os.Open("config.yml")
	if err != nil {
		return Config{}, err
	}
	defer f.Close()

	var cfg Config

	decoder := yaml.NewDecoder(f)

	err = decoder.Decode(&cfg)
	if err != nil {
		return Config{}, err
	}

	return cfg, nil
}
