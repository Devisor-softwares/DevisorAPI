package config

import (
    "gopkg.in/yaml.v3"
    "io/ioutil"
)

type Route struct {
    Method  string `yaml:"method"`
    Path    string `yaml:"path"`
    Handler string `yaml:"handler"`
    Auth    bool   `yaml:"auth"`
}

type RoutesConfig struct {
    Routes []Route `yaml:"routes"`
}

func LoadRoutes(file string) (*RoutesConfig, error) {
    data, err := ioutil.ReadFile(file)
    if err != nil {
        return nil, err
    }
    var cfg RoutesConfig
    if err := yaml.Unmarshal(data, &cfg); err != nil {
        return nil, err
    }
    return &cfg, nil
}
