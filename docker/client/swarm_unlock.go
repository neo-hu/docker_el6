package client 

import (
	"context"

	"github.com/neo-hu/docker_el6/docker/api/types/swarm"
)

// SwarmUnlock unlocks locked swarm.
func (cli *Client) SwarmUnlock(ctx context.Context, req swarm.UnlockRequest) error {
	serverResp, err := cli.post(ctx, "/swarm/unlock", nil, req, nil)
	ensureReaderClosed(serverResp)
	return err
}
