// +build linux

package fifo

import (
	"fmt"
	"os"
	"sync"
	"syscall"

	"github.com/pkg/errors"
)

const O_PATH = 010000000

type handle struct {
	f         *os.File
	fd        uintptr
	dev       uint64
	ino       uint64
	closeOnce sync.Once
	name      string
}

func getHandle(fn string) (*handle, error) {
	f, err := os.OpenFile(fn, O_PATH|syscall.SOCK_NONBLOCK, 0)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to open %v with O_PATH", fn)
	}

	var (
		stat syscall.Stat_t
		fd   = f.Fd()
	)
	if err := syscall.Fstat(int(fd), &stat); err != nil {
		f.Close()
		return nil, errors.Wrapf(err, "failed to stat handle %v", fd)
	}

	h := &handle{
		f:    f,
		name: fn,
		dev:  uint64(stat.Dev),
		ino:  stat.Ino,
		fd:   fd,
	}

	// check /proc just in case
	if _, err := os.Stat(h.procPath()); err != nil {
		f.Close()
		return nil, errors.Wrapf(err, "couldn't stat %v", h.procPath())
	}

	return h, nil
}

func (h *handle) procPath() string {
	return fmt.Sprintf("/proc/self/fd/%d", h.fd)
}

func (h *handle) Name() string {
	return h.name
}

func (h *handle) Path() (string, error) {
	var stat syscall.Stat_t
	if err := syscall.Stat(h.procPath(), &stat); err != nil {
		return "", errors.Wrapf(err, "path %v could not be statted", h.procPath())
	}
	if uint64(stat.Dev) != h.dev || stat.Ino != h.ino {
		return "", errors.Errorf("failed to verify handle %v/%v %v/%v", stat.Dev, h.dev, stat.Ino, h.ino)
	}
	return h.procPath(), nil
}

func (h *handle) Close() error {
	h.closeOnce.Do(func() {
		h.f.Close()
	})
	return nil
}
