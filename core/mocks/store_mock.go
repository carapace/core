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
	}
	Users *MockUserStore
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
		}{OwnerSet: NewMockOwnerSet(controller), UserSet: NewMockUserSet(controller), Config: NewMockConfigManager(controller)},
		Users: NewMockUserStore(controller),
	}

	return &core.Store{
			DB: db,
			Sets: &core.Sets{
				OwnerSet: ctrl.Sets.OwnerSet,
				UserSet:  ctrl.Sets.UserSet,
				Config:   ctrl.Sets.Config,
			},
			Users: ctrl.Users,
		}, ctrl, func() {
			db.Close()
		}
}
