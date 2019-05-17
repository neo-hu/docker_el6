#!/usr/bin/env bash

go run -tags "no_btrfs" main.go main_unix.go config_linux.go config.go builtins_linux.go builtins.go  publish.go --config containerd.toml -l debug $@