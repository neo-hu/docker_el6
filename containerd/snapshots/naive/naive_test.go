package naive

import (
	"context"
	"runtime"
	"testing"

	"github.com/neo-hu/docker_el6/containerd/snapshots"
	"github.com/neo-hu/docker_el6/containerd/snapshots/testsuite"
	"github.com/neo-hu/docker_el6/containerd/testutil"
)

func newSnapshotter(ctx context.Context, root string) (snapshots.Snapshotter, func() error, error) {
	snapshotter, err := NewSnapshotter(root)
	if err != nil {
		return nil, nil, err
	}

	return snapshotter, func() error { return snapshotter.Close() }, nil
}

func TestNaive(t *testing.T) {
	if runtime.GOOS == "windows" {
		t.Skip("snapshotter not implemented on windows")
	}
	testutil.RequiresRoot(t)
	testsuite.SnapshotterSuite(t, "Naive", newSnapshotter)
}
