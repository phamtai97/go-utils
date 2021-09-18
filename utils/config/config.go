package config

import (
	"encoding/json"
	"flag"
	"io/ioutil"
	"reflect"

	"github.com/phamtai97/go-utils/utils/logger"
	"go.uber.org/zap"
	"gopkg.in/yaml.v3"
)

// LoadYaml loads configuration from specific config path.
func LoadYaml(config interface{}, configPath string) error {
	buf, err := ioutil.ReadFile(configPath)
	if err != nil {
		return err
	}

	return yaml.Unmarshal(buf, config)
}

// LoadYamlByFlag loads the configuration from the env path.
func LoadYamlByFlag(config interface{}, flagPath string) error {
	var filePath string
	flag.StringVar(&filePath, flagPath, "config.yaml", "Path of config file")
	flag.Parse()

	return LoadYaml(config, filePath)
}

// LoadJson loads configuration from specific config path.
func LoadJson(config interface{}, configPath string) error {
	buf, err := ioutil.ReadFile(configPath)
	if err != nil {
		return err
	}

	return json.Unmarshal(buf, config)
}

// LoadJsonlByFlag loads the configuration from the env path.
func LoadJsonlByFlag(config interface{}, flagPath string) error {
	var filePath string
	flag.StringVar(&filePath, flagPath, "config.json", "Path of config file")
	flag.Parse()

	return LoadJson(config, filePath)
}

// Print to config
func Print(config interface{}, omittedKeys ...string) error {
	configMap, err := toMap(config)
	if err != nil {
		return err
	}

	if len(omittedKeys) == 0 {
		print(configMap)
		return nil
	}

	replaceHotKey(configMap, omittedKeys...)

	print(configMap)
	return nil
}

func print(configMap map[string]interface{}) {
	logger.Info("Print out application configuration", zap.Any("Configuration", configMap))
}

func replaceHotKey(configMap map[string]interface{}, omittedKeys ...string) {
	for _, key := range omittedKeys {
		delete(configMap, key)
	}

	for _, v := range configMap {
		if reflect.ValueOf(v).Kind() == reflect.Map {
			replaceHotKey(v.(map[string]interface{}), omittedKeys...)
		}
	}
}

func toMap(config interface{}) (map[string]interface{}, error) {
	var res map[string]interface{}
	data, err := json.Marshal(config)
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(data, &res); err != nil {
		return nil, err
	}

	return res, nil
}
