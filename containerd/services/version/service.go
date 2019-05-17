package version

import (
	api "github.com/neo-hu/docker_el6/containerd/api/services/version/v1"
	"github.com/neo-hu/docker_el6/containerd/plugin"
	ctrdversion "github.com/neo-hu/docker_el6/containerd/version"
	ptypes "github.com/gogo/protobuf/types"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

var _ api.VersionServer = &service{}

func init() {
	plugin.Register(&plugin.Registration{
		Type:   plugin.GRPCPlugin,
		ID:     "version",
		InitFn: initFunc,
	})
}

func initFunc(ic *plugin.InitContext) (interface{}, error) {
	return &service{}, nil
}

type service struct {
}

func (s *service) Register(server *grpc.Server) error {
	api.RegisterVersionServer(server, s)
	return nil
}

func (s *service) Version(ctx context.Context, _ *ptypes.Empty) (*api.VersionResponse, error) {
	return &api.VersionResponse{
		Version:  ctrdversion.Version,
		Revision: ctrdversion.Revision,
	}, nil
}
