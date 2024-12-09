package config

import (
	"os"
	"strconv"
	"strings"
)

type Config struct {
	Postgres Postgres `yaml:"postgres"`
}

var conf *Config

func NewConfig() *Config {
	if conf != nil {
		return conf
	}

	config := Config{}
	conf = &config

	return conf
}

func loadStringProp(property string) string {
	if isEnvMask(property) {
		return loadEnv(property)
	}

	return property
}

func loadIntProp(property string) int {
	if isEnvMask(property) {
		property = loadEnv(property)
	}

	res, err := strconv.Atoi(property)
	if err != nil {
		panic(err)
	}

	return res
}

func isEnvMask(property string) bool {
	property = strings.TrimSpace(property)

	return strings.HasPrefix(property, "${") && strings.HasSuffix(property, "}")
}

func loadEnv(property string) string {
	return os.Getenv(property[2 : len(property)-1])
}
