package auth

import (
	"context"
	"fmt"
	"github.com/carapace/core/api/v0/proto"
	"github.com/carapace/core/core"
	"github.com/carapace/core/pkg/responses"
)

type wrapped struct {
	infoService   func() (*v0.Info, error)
	configService func(ctx context.Context, config *v0.Config) (*v0.Response, error)
}

func (s *wrapped) ConfigService(ctx context.Context, config *v0.Config) (*v0.Response, error) {
	return s.configService(ctx, config)
}

func (s *wrapped) InfoService() (*v0.Info, error) {
	return s.infoService()
}

// Signed only verifies that the config object's signatures are all correctly signed
func (auth *Manager) Signed(service core.APIService) core.APIService {
	return &wrapped{
		infoService: service.InfoService,
		configService: func(ctx context.Context, config *v0.Config) (*v0.Response, error) {
			ok, wrongSig, err := auth.Check(config)
			if err != nil {
				return nil, err
			}

			if !ok {
				return response.MSG(v0.Code_BadRequest, fmt.Sprintf("incorrect signature for: %s", wrongSig)), nil
			}
			return service.ConfigService(ctx, config)
		},
	}
}

// Root requires the config message to be signed by a quorum amount of owners.
//
// Root errors if the node does not have any owners
func (auth *Manager) Root(service core.APIService) core.APIService {
	return &wrapped{
		infoService: service.InfoService,
		configService: func(ctx context.Context, config *v0.Config) (*v0.Response, error) {
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
			return service.ConfigService(ctx, config)
		},
	}
}

// RootOrBackupOrNoOwners requires that the config obj is signed by root quorum, or the Node has no owners yet.
//
// RootOrBackupOrNoOwners allows for the use of recovery keys
func (auth *Manager) RootOrBackupOrNoOwners(service core.APIService) core.APIService {
	return &wrapped{
		infoService: service.InfoService,
		configService: func(ctx context.Context, config *v0.Config) (*v0.Response, error) {
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
				return service.ConfigService(ctx, config)
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
			return service.ConfigService(ctx, config)
		},
	}
}

// RegularAuth allows for defining three requirements to authenticate/authorize.
//
// It requires:
// 	1. Signatures are correct
// 	2. Node has owners
//  3. minLevel < 0
//  4. minSignees < 0
//  5. maxSignees < 0
//  6. minSignees < len(witness.Signatures) < maxSignees
//  7. the sum of the signees AuthLevel > minLevel
//
// This function will panic during startup if it is not properly configured.
func (auth *Manager) RegularAuth(minLevel int32, minSignees uint8, maxSignees uint8, service core.APIService) core.APIService {

	// this code is evaluated on startup, not during authentication
	if minLevel < 0 {
		panic("auth.Wrapper: RegularAuth: minLevel must be gte 0")
	}

	return &wrapped{
		infoService: service.InfoService,
		configService: func(ctx context.Context, config *v0.Config) (*v0.Response, error) {

			if uint8(len(config.Witness.Signatures)) < minSignees {
				return response.MSG(
					v0.Code_BadRequest,
					fmt.Sprintf("at least %v signees are needed for this operation", minSignees)), nil
			}

			if uint8(len(config.Witness.Signatures)) > maxSignees {
				return response.MSG(
					v0.Code_BadRequest,
					fmt.Sprintf("no more than %v signees may be used for this operation", minSignees)), nil
			}

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
				return response.MSG(v0.Code_NotImplemented, "An OwnerSet must first be provided"), nil
			}

			tx := core.TXFromContext(ctx)

			// compute the total auth level
			var totalAuthLevel int32
			for _, user := range config.Witness.Signatures {
				user, err := auth.Store.Users.Get(tx, user.GetPrimaryPublicKey())
				if err != nil {
					return response.Err(err), nil
				}
				totalAuthLevel += user.AuthLevel
			}
			if totalAuthLevel < minLevel {
				return response.MSG(
					v0.Code_BadRequest,
					fmt.Sprintf("the minimum combined AuthLevel required for this operation is %v", minLevel)), nil
			}
			return service.ConfigService(ctx, config)
		},
	}

}
