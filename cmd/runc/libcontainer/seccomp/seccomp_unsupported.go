// +build !linux !cgo !seccomp

package seccomp

import (
	"errors"

	"github.com/neo-hu/docker_el6/cmd/runc/libcontainer/configs"
)

var ErrSeccompNotEnabled = errors.New("seccomp: config provided but seccomp not supported")

// InitSeccomp does nothing because seccomp is not supported.
func InitSeccomp(config *configs.Seccomp) error {
	if config != nil {
		return ErrSeccompNotEnabled
	}
	return nil
}

// IsEnabled returns false, because it is not supported.
func IsEnabled() bool {
	return false
}
