package utils

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

func ReadFromYaml(path string, target any) error {
	content, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}
	err = yaml.Unmarshal(content, target)
	if err != nil {
		return err
	}
	return nil
}
