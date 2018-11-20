package router

import (
	"context"
	"database/sql"
	"github.com/carapace/core/core/mocks"
	"github.com/golang/mock/gomock"
	"testing"

	"github.com/carapace/core/api/v0/proto"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestRouter_Register(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	service := mock.NewMockAPIService(ctrl)

	rt := New()
	service.EXPECT().InfoService().Return(&v0.Info{ApiVersion: "v0", Kinds: []string{"ownerSet"}}, nil).Times(1)
	assert.NotPanics(t, func() {
		rt.Register(service)
	})

	// verify that a double registration fails
	service.EXPECT().InfoService().Return(&v0.Info{ApiVersion: "v0", Kinds: []string{"ownerSet"}}, nil).Times(1)
	assert.Panics(t, func() {
		rt.Register(service)
	})
}

func TestV0_Route(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	service := mock.NewMockAPIService(ctrl)

	rt := New()
	service.EXPECT().InfoService().Return(&v0.Info{ApiVersion: "v0", Kinds: []string{"walletSet"}}, nil).Times(1)
	rt.Register(service)

	tcs := []struct {
		apiVersion string
		kind       string

		prep *gomock.Call
		res  *v0.Response
		err  error
		desc string
	}{
		{
			prep:       service.EXPECT().ConfigService(gomock.Any(), gomock.Any(), gomock.Any()).Return(nil, nil),
			apiVersion: "v0",
			kind:       "walletSet",

			res:  nil,
			err:  nil,
			desc: "correct apiVersion and kind should return nil, nil",
		},
		{
			apiVersion: "v1",
			kind:       "walletSet",

			res:  nil,
			err:  errors.New("unregistered apiVersion"),
			desc: "incorrect apiVersion and correct kind should return nil, unregistered apiVersion",
		},
		{
			apiVersion: "v0",
			kind:       "replicationSet",

			res:  nil,
			err:  errors.New("unregistered kind"),
			desc: "correct apiVersion and incorrect kind should return nil, unregistered kind",
		},
		{
			apiVersion: "v1",
			kind:       "replicationSet",

			res:  nil,
			err:  errors.New("unregistered apiVersion"),
			desc: "incorrect apiVersion and incorrect kind should return nil, unregistered apiVersion",
		},
	}

	for _, tc := range tcs {
		res, err := rt.Route(context.Background(), &v0.Config{Header: &v0.Header{ApiVersion: tc.apiVersion, Kind: tc.kind}}, &sql.Tx{})

		if tc.err == nil {
			require.NoError(t, err, tc.desc)
		} else {
			assert.EqualError(t, err, tc.err.Error(), tc.desc)
		}
		assert.Equal(t, tc.res, res, tc.desc)
	}
}
