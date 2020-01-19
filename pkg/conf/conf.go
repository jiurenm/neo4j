package conf

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

type Configs struct {
	Neo4j neo4j `yaml:"neo4j"`
}

type neo4j struct {
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
}

func ConfigFile() string {
	return "configs/conf.yaml"
}

func New(path string) (*Configs, error) {
	if file, err := ioutil.ReadFile(path); err != nil {
		return nil, err
	} else {
		var conf Configs
		if err := yaml.Unmarshal(file, &conf); err != nil {
			return nil, err
		}
		return &conf, nil
	}
}
