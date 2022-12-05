package proc

import (
	"fmt"
	"strings"
)

type Envs map[string]string

func Merge(e, f Envs) Envs {
	g := make(Envs)
	for k, v := range e {
		g[k] = v
	}
	for k, v := range f {
		g[k] = v
	}
	return g
}

func updatedEnvFrom(newValues Envs, oldEnvs []string) []string {
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

func parseEnv(envs []string) Envs {
	envMap := make(Envs)
	for _, kv := range envs {
		parts := strings.SplitN(kv, "=", 2)
		if len(parts) == 2 {
			envMap[parts[0]] = parts[1]
		}
	}
	return envMap
}
