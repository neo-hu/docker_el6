package main

// register containerd builtins here
import (
	_ "github.com/neo-hu/docker_el6/containerd/diff/walking"
	_ "github.com/neo-hu/docker_el6/containerd/gc/scheduler"
	_ "github.com/neo-hu/docker_el6/containerd/services/containers"
	_ "github.com/neo-hu/docker_el6/containerd/services/content"
	_ "github.com/neo-hu/docker_el6/containerd/services/diff"
	_ "github.com/neo-hu/docker_el6/containerd/services/events"
	_ "github.com/neo-hu/docker_el6/containerd/services/healthcheck"
	_ "github.com/neo-hu/docker_el6/containerd/services/images"
	_ "github.com/neo-hu/docker_el6/containerd/services/introspection"
	_ "github.com/neo-hu/docker_el6/containerd/services/leases"
	_ "github.com/neo-hu/docker_el6/containerd/services/namespaces"
	_ "github.com/neo-hu/docker_el6/containerd/services/snapshots"
	_ "github.com/neo-hu/docker_el6/containerd/services/tasks"
	_ "github.com/neo-hu/docker_el6/containerd/services/version"
)
