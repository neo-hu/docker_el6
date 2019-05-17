package daemon

import (
	"time"

	"github.com/neo-hu/docker_el6/docker/api/types/container"
	"github.com/neo-hu/docker_el6/docker/libcontainerd"
	specs "github.com/opencontainers/runtime-spec/specs-go"
)

func toContainerdResources(resources container.Resources) *libcontainerd.Resources {
	var r libcontainerd.Resources

	r.BlockIO = &specs.LinuxBlockIO{
		Weight: &resources.BlkioWeight,
	}

	shares := uint64(resources.CPUShares)
	r.CPU = &specs.LinuxCPU{
		Shares: &shares,
		Cpus:   resources.CpusetCpus,
		Mems:   resources.CpusetMems,
	}

	var (
		period uint64
		quota  int64
	)
	if resources.NanoCPUs != 0 {
		period = uint64(100 * time.Millisecond / time.Microsecond)
		quota = resources.NanoCPUs * int64(period) / 1e9
	}
	if quota == 0 && resources.CPUQuota != 0 {
		quota = resources.CPUQuota
	}
	if period == 0 && resources.CPUPeriod != 0 {
		period = uint64(resources.CPUPeriod)
	}

	r.CPU.Period = &period
	r.CPU.Quota = &quota

	r.Memory = &specs.LinuxMemory{
		Limit:       &resources.Memory,
		Reservation: &resources.MemoryReservation,
		Kernel:      &resources.KernelMemory,
	}

	if resources.MemorySwap > 0 {
		r.Memory.Swap = &resources.MemorySwap
	}

	return &r
}