#!/usr/bin/env bash

DOCKER_NOWARN_KERNEL_VERSION=1 go run -p 4 -tags "exclude_graphdriver_btrfs" *.go --containerd /var/run/docker/containerd/docker-containerd.sock -b none -l debug
