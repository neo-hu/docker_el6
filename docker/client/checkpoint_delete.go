package client

import (
	"context"
	"net/url"

	"github.com/neo-hu/docker_el6/docker/api/types"
)

// CheckpointDelete deletes the checkpoint with the given name from the given container
func (cli *Client) CheckpointDelete(ctx context.Context, containerID string, options types.CheckpointDeleteOptions) error {
	query := url.Values{}
	if options.CheckpointDir != "" {
		query.Set("dir", options.CheckpointDir)
	}

	resp, err := cli.delete(ctx, "/containers/"+containerID+"/checkpoints/"+options.CheckpointID, query, nil)
	ensureReaderClosed(resp)
	return err
}