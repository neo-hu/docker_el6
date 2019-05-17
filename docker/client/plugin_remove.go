package client 

import (
	"context"
	"net/url"

	"github.com/neo-hu/docker_el6/docker/api/types"
)

// PluginRemove removes a plugin
func (cli *Client) PluginRemove(ctx context.Context, name string, options types.PluginRemoveOptions) error {
	query := url.Values{}
	if options.Force {
		query.Set("force", "1")
	}

	resp, err := cli.delete(ctx, "/plugins/"+name, query, nil)
	defer ensureReaderClosed(resp)
	return wrapResponseError(err, resp, "plugin", name)
}
