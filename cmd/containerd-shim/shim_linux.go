package main

import (
	"github.com/neo-hu/docker_el6/containerd/sys"
	"os"
	"os/signal"

	"github.com/neo-hu/docker_el6/containerd/reaper"
	"github.com/neo-hu/docker_el6/go-runc"
	"github.com/stevvooe/ttrpc"
	"golang.org/x/sys/unix"
)

// setupSignals creates a new signal handler for all signals and sets the shim as a
// sub-reaper so that the container processes are reparented
func setupSignals() (chan os.Signal, error) {
	signals := make(chan os.Signal, 32)
	signal.Notify(signals, unix.SIGTERM, unix.SIGINT, unix.SIGCHLD, unix.SIGPIPE)
	// make sure runc is setup to use the monitor
	// for waiting on processes
	runc.Monitor = reaper.Default
	// set the shim as the subreaper for all orphaned processes created by the container
	// Linux 3.4
	if err := sys.SetSubreaper(1); err != nil {
		return nil, err
	}
	return signals, nil
}

func newServer() (*ttrpc.Server, error) {
	return ttrpc.NewServer(ttrpc.WithServerHandshaker(ttrpc.UnixSocketRequireSameUser()))
}
