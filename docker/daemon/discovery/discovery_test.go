package discovery

import (
	"fmt"
	"testing"
	"time"

	"github.com/gotestyourself/gotestyourself/assert"
	is "github.com/gotestyourself/gotestyourself/assert/cmp"
)

func TestDiscoveryOptsErrors(t *testing.T) {
	var testcases = []struct {
		doc  string
		opts map[string]string
	}{
		{
			doc:  "discovery.ttl < discovery.heartbeat",
			opts: map[string]string{"discovery.heartbeat": "10", "discovery.ttl": "5"},
		},
		{
			doc:  "discovery.ttl == discovery.heartbeat",
			opts: map[string]string{"discovery.heartbeat": "10", "discovery.ttl": "10"},
		},
		{
			doc:  "negative discovery.heartbeat",
			opts: map[string]string{"discovery.heartbeat": "-10", "discovery.ttl": "10"},
		},
		{
			doc:  "negative discovery.ttl",
			opts: map[string]string{"discovery.heartbeat": "10", "discovery.ttl": "-10"},
		},
		{
			doc:  "invalid discovery.heartbeat",
			opts: map[string]string{"discovery.heartbeat": "invalid"},
		},
		{
			doc:  "invalid discovery.ttl",
			opts: map[string]string{"discovery.ttl": "invalid"},
		},
	}

	for _, testcase := range testcases {
		_, _, err := discoveryOpts(testcase.opts)
		assert.Check(t, is.ErrorContains(err, ""), testcase.doc)
	}
}

func TestDiscoveryOpts(t *testing.T) {
	clusterOpts := map[string]string{"discovery.heartbeat": "10", "discovery.ttl": "20"}
	heartbeat, ttl, err := discoveryOpts(clusterOpts)
	assert.NilError(t, err)
	assert.Check(t, is.Equal(10*time.Second, heartbeat))
	assert.Check(t, is.Equal(20*time.Second, ttl))

	clusterOpts = map[string]string{"discovery.heartbeat": "10"}
	heartbeat, ttl, err = discoveryOpts(clusterOpts)
	assert.NilError(t, err)
	assert.Check(t, is.Equal(10*time.Second, heartbeat))
	assert.Check(t, is.Equal(10*defaultDiscoveryTTLFactor*time.Second, ttl))

	clusterOpts = map[string]string{"discovery.ttl": "30"}
	heartbeat, ttl, err = discoveryOpts(clusterOpts)
	assert.NilError(t, err)

	if ttl != 30*time.Second {
		t.Fatalf("TTL - Expected : %v, Actual : %v", 30*time.Second, ttl)
	}

	expected := 30 * time.Second / defaultDiscoveryTTLFactor
	if heartbeat != expected {
		t.Fatalf("Heartbeat - Expected : %v, Actual : %v", expected, heartbeat)
	}

	discoveryTTL := fmt.Sprintf("%d", defaultDiscoveryTTLFactor-1)
	clusterOpts = map[string]string{"discovery.ttl": discoveryTTL}
	heartbeat, _, err = discoveryOpts(clusterOpts)
	if err == nil && heartbeat == 0 {
		t.Fatal("discovery.heartbeat must be positive")
	}

	clusterOpts = map[string]string{}
	heartbeat, ttl, err = discoveryOpts(clusterOpts)
	if err != nil {
		t.Fatal(err)
	}

	if heartbeat != defaultDiscoveryHeartbeat {
		t.Fatalf("Heartbeat - Expected : %v, Actual : %v", defaultDiscoveryHeartbeat, heartbeat)
	}

	expected = defaultDiscoveryHeartbeat * defaultDiscoveryTTLFactor
	if ttl != expected {
		t.Fatalf("TTL - Expected : %v, Actual : %v", expected, ttl)
	}
}
