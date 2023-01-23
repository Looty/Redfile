package config

import (
	"os"
)

func SetGlobalEnvVariable(config string) {
	for e := range f.Env {
		setEnv(e, f.Env[e])
	}
}

func SetCmdEnv(envs map[string]string) {
	for e := range envs {
		setEnv(e, envs[e])
	}
}

func setEnv(key string, value string) {
	os.Setenv(key, value)
}
