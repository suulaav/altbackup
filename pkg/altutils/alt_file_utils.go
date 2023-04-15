package altutils

import (
	"github.com/suulaav/altbackup/constants"
	"gopkg.in/yaml.v3"
	"os"
)

type NestedMap map[string]interface{}

func ReadYaml(path string) *NestedMap {
	yamlMap := make(NestedMap)
	if path == "" {
		panic("invalid File path")
	}
	CheckExtension(path, constants.YamlExtension)
	yamlFile, err := os.ReadFile(path)
	CheckError(err)
	err = yaml.Unmarshal(yamlFile, &yamlMap)
	return &yamlMap
}

func ReadFile(path string) []byte {
	if path == "" {
		panic("invalid File path")
	}
	file, err := os.ReadFile(path)
	CheckError(err)
	return file
}
