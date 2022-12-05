package proc

import (
	"fmt"
	"strings"
)

type Env map[string]string

func Merge(e, f Env) Env {
	g := make(Env)
	for k, v := range e {
		g[k] = v
	}
	for k, v := range f {
		g[k] = v
	}
	return g
}

func updatedEnvFrom(newValues Env, oldEnvs []string) []string {
	envMap := parseEnv(oldEnvs)
	for k, v := range newValues {
		envMap[k] = v
	}
	var envs []string
	for k, v := range envMap {
		envs = append(envs, fmt.Sprintf("%s=%s", k, v))
	}
	return envs
}

func parseEnv(envs []string) Env {
	envMap := make(Env)
	for _, kv := range envs {
		parts := strings.SplitN(kv, "=", 2)
		if len(parts) == 2 {
			envMap[parts[0]] = parts[1]
		}
	}
	return envMap
}
