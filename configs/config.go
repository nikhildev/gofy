package configs

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"os"
	"regexp"
	"strings"
)

func LoadAppConfig() (AppConfig, error) {
	var appConfig AppConfig
	return appConfig, Load("application.yaml", &appConfig)
}

func Load(file string, out interface{}) error {
	err := loadYaml(file, out)

	if err != nil {
		return err
	}

	// if override yaml exists, we override the config with the override data
	fileNameWithoutExt := strings.Split(file, ".yaml")[0]
	if _, err := os.Stat(fileNameWithoutExt + "-override.yaml"); err == nil {
		err := loadYaml("application-override.yaml", out)
		if err != nil {
			return err
		}
	}

	return err
}

func loadYaml(file string, out interface{}) error {
	data, err := os.ReadFile(file)

	if err != nil {
		return err
	}

	if err = readYamlConfig(data, out); err != nil {
		return err
	}

	return nil
}

var envVariableRegexp = regexp.MustCompile(`\${.+}`)

func readYamlConfig(data []byte, out interface{}) error {
	result := envVariableRegexp.ReplaceAllStringFunc(string(data), func(str string) string {
		split := strings.SplitN(str, ":", 2) // return max 2 substrings (so it will be split only on first match)
		hasDefaultValue := len(split) > 1    // if we can split on :, a default value is present

		envVariable := split[0]

		if hasDefaultValue {
			// if defaultValue is present, it was split half way through and we need to add a bracket at the end
			envVariable = fmt.Sprintf("%s}", envVariable)
		}

		expanded := os.ExpandEnv(envVariable)

		if len(expanded) > 0 || !hasDefaultValue { // no default value we can return empty value anyway
			return expanded
		}

		defaultValue := split[1]
		defaultValue = defaultValue[0 : len(defaultValue)-1] // remove bracket

		return defaultValue
	})

	return yaml.Unmarshal([]byte(result), out)
}
