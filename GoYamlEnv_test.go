package goyamlenv

import "testing"

type Configs struct {
	APP      Apps      `mapstructure:"app"`
	Endpoint Endpoints `mapstructure:"endpoint"`
}

type Apps struct {
	ENV         string `mapstructure:"env"`
	PORT        string `mapstructure:"port"`
	ContextPath string `mapstructure:"context_path"`
	HOST        string `mapstructure:"host"`
	RootPath    string `mapstructure:"root_path"`
}

type Endpoints struct {
	Classroom struct {
		HOST string `mapstructure:"host"`
		PATH string `mapstructure:"path"`
	} `mapstructure:"classroom"`
}

func TestXxx(t *testing.T) {
	env := Configs{}
	err := LoadConfig(&env)
	t.Log(err)
}
