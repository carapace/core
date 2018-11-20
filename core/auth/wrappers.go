package auth

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/carapace/core/api/v0/proto"
	"github.com/carapace/core/core"
	"github.com/carapace/core/pkg/responses"
	"github.com/carapace/core/pkg/v0"
)

type wrapped struct {
	infoService   func() (*v0.Info, error)
	configService func(ctx context.Context, config *v0.Config, tx *sql.Tx) (*v0.Response, error)
}

func (s *wrapped) ConfigService(ctx context.Context, config *v0.Config, tx *sql.Tx) (*v0.Response, error) {
	return s.configService(ctx, config, tx)
}

func (s *wrapped) InfoService() (*v0.Info, error) {
	return s.infoService()
}

// Signed only verifies that the config object's signatures are all correctly signed
func (auth *Manager) Signed(service core.APIService) core.APIService {
	return &wrapped{
		infoService: service.InfoService,
		configService: func(ctx context.Context, config *v0.Config, tx *sql.Tx) (*v0.Response, error) {
			ok, wrongSig, err := auth.Check(config)
			if err != nil {
				return nil, err
			}

			if !ok {
				return v0_handler.WriteMSG(v0.Code_BadRequest, fmt.Sprintf("incorrect signature for: %s", wrongSig)), nil
			}
			return service.ConfigService(ctx, config, tx)
		},
	}
}

// Root requires the config message to be signed by a quorum amount of owners.
//
// Root errors if the node does not have any owners
func (auth *Manager) Root(service core.APIService) core.APIService {
	return &wrapped{
		infoService: service.InfoService,
		configService: func(ctx context.Context, config *v0.Config, tx *sql.Tx) (*v0.Response, error) {
			ok, wrongSig, err := auth.Check(config)
			if err != nil {
				return nil, err
			}

			if !ok {
				return response.MSG(v0.Code_BadRequest, fmt.Sprintf("incorrect signature for: %s", wrongSig)), nil
			}

			have, err := auth.HaveOwners()
			if err != nil {
				return nil, err
			}

			if !have {
				return response.MSG(v0.Code_UnAuthorized, fmt.Sprintf("node does not have owners yet")), nil
			}

			root, err := auth.GrantRoot(config.Witness)
			if err != nil {
				return response.Err(err), nil
			}

			if !root {
				return response.MSG(v0.Code_UnAuthorized, fmt.Sprintf("cannot grant root: %s", err.Error())), nil
			}
			return service.ConfigService(ctx, config, tx)
		},
	}
}

// RootOrBackupOrNoOwners requires that the config obj is signed by root quorum, or the Node has no owners yet.
//
// RootOrBackupOrNoOwners allows for the use of recovery keys
func (auth *Manager) RootOrBackupOrNoOwners(service core.APIService) core.APIService {
	return &wrapped{
		infoService: service.InfoService,
		configService: func(ctx context.Context, config *v0.Config, tx *sql.Tx) (*v0.Response, error) {
			ok, wrongSig, err := auth.Check(config)
			if err != nil {
				return response.MSG(v0.Code_BadRequest, err.Error()), nil
			}

			if !ok {
				return response.MSG(v0.Code_BadRequest, fmt.Sprintf("incorrect signature for: %s", wrongSig)), nil
			}

			have, err := auth.HaveOwners()
			if err != nil {
				return nil, err
			}

			if !have {
				return service.ConfigService(ctx, config, tx)
			}

			root, err := auth.GrantRoot(config.Witness)
			if err != nil {
				return response.Err(err), nil
			}

			if !root {
				root, err = auth.GrantBackupRoot(config.Witness)
			}

			if err != nil {
				return response.Err(err), nil
			}

			if !root {
				return response.MSG(v0.Code_UnAuthorized, fmt.Sprintf("cannot grant root: insufficient quorum")), nil
			}
			return service.ConfigService(ctx, config, tx)
		},
	}
}
