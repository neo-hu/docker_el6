package storage

import (
	"path/filepath"
	"testing"

	// Does not require root but flag must be defined for snapshot tests

	_ "github.com/neo-hu/docker_el6/containerd/testutil"
)

func TestMetastore(t *testing.T) {
	MetaStoreSuite(t, "Metastore", func(root string) (*MetaStore, error) {
		return NewMetaStore(filepath.Join(root, "metadata.db"))
	})
}

func BenchmarkSuite(b *testing.B) {
	Benchmarks(b, "BoltDBBench", func(root string) (*MetaStore, error) {
		return NewMetaStore(filepath.Join(root, "metadata.db"))
	})
}
