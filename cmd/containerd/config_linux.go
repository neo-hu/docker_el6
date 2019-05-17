package main

import (
	"github.com/neo-hu/docker_el6/containerd/defaults"
	"github.com/neo-hu/docker_el6/containerd/server"
)

func defaultConfig() *server.Config {
	return &server.Config{
		Root:  defaults.DefaultRootDir,
		State: defaults.DefaultStateDir,
		GRPC: server.GRPCConfig{
			Address: defaults.DefaultAddress,
		},
		Debug: server.Debug{
			Level:   "info",
			Address: defaults.DefaultDebugAddress,
		},
	}
}
