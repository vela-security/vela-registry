package registry

import "github.com/vela-security/vela-public/assert"

func WithEnv(env assert.Environment) {
	env.Error("not support registry with linux")
}
