// +build linux

package client

import (
	"os/exec"
	"syscall"

	"github.com/neo-hu/docker_el6/containerd/cgroups"
	"github.com/pkg/errors"
)

func getSysProcAttr() *syscall.SysProcAttr {
	return &syscall.SysProcAttr{
		Setpgid: true,
	}
}

func setCgroup(cgroupPath string, cmd *exec.Cmd) error {
	cg, err := cgroups.Load(cgroups.V1, cgroups.StaticPath(cgroupPath))
	if err != nil {
		return errors.Wrapf(err, "failed to load cgroup %s", cgroupPath)
	}
	if err := cg.Add(cgroups.Process{
		Pid: cmd.Process.Pid,
	}); err != nil {
		return errors.Wrapf(err, "failed to join cgroup %s", cgroupPath)
	}
	return nil
}
