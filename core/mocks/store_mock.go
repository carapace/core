package mock

import (
	"github.com/carapace/core/core"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"gopkg.in/DATA-DOG/go-sqlmock.v1"
	"testing"
)

type StoreController struct {
	DB   sqlmock.Sqlmock
	Sets struct {
		OwnerSet *MockOwnerSet
		UserSet  *MockUserSet
		Config   *MockConfigManager
		Identity *MockIdentitySet
	}
	Users    *MockUserStore
	Policies *MockPolicyManager
}

func NewStoreMock(t *testing.T, controller *gomock.Controller) (*core.Store, *StoreController, func()) {
	db, mockdb, err := sqlmock.New()
	require.NoError(t, err)

	ctrl := &StoreController{
		DB: mockdb,
		Sets: struct {
			OwnerSet *MockOwnerSet
			UserSet  *MockUserSet
			Config   *MockConfigManager
			Identity *MockIdentitySet
		}{
			OwnerSet: NewMockOwnerSet(controller),
			UserSet:  NewMockUserSet(controller),
			Config:   NewMockConfigManager(controller),
			Identity: NewMockIdentitySet(controller),
		},
		Users:    NewMockUserStore(controller),
		Policies: NewMockPolicyManager(controller),
	}

	return &core.Store{
			DB: db,
			Sets: &core.Sets{
				OwnerSet: ctrl.Sets.OwnerSet,
				UserSet:  ctrl.Sets.UserSet,
				Config:   ctrl.Sets.Config,
				Identity: ctrl.Sets.Identity,
			},
			Users:    ctrl.Users,
			Policies: ctrl.Policies,
		}, ctrl, func() {
			db.Close()
		}
}
