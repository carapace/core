package ownerset

import (
	"context"
	"fmt"
	"sync"
	"testing"

	"github.com/pkg/errors"

	"github.com/carapace/core/internal/v0"
	"github.com/golang/protobuf/proto"

	"github.com/golang/protobuf/ptypes"
	"github.com/golang/protobuf/ptypes/any"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/carapace/core/api/v0/proto"
	"github.com/golang/mock/gomock"
)

// genAny is used to ignore the error return of marshal.Any when declaring test
// cases
func genAny(t *testing.T, spec proto.Message) *any.Any {
	a, err := ptypes.MarshalAny(spec)
	require.NoError(t, err)
	return a
}

func TestHandler_Handle(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockObj := mock.NewMockAuthenticator(mockCtrl)

	handler := Handler{
		auth: mockObj,
		mu:   &sync.RWMutex{},
	}

	tcs := []struct {
		config *v0.Config

		err      error
		response *v0.Response
		desc     string
	}{
		{
			config:   &v0.Config{Header: &v0.Header{Kind: "OWNERSET"}},
			err:      v0_handler.ErrIncorrectKind,
			response: nil,
			desc:     "no owners and incorrect kind should return v0_handler.ErrIncorrectKind",
		},
	}

	for _, tc := range tcs {
		res, err := handler.Handle(context.Background(), tc.config)
		assert.Equal(t, tc.response, res)
		if err != nil {
			assert.EqualError(t, err, tc.err.Error(), tc.desc)
		}
	}
}

func TestHandler_newOwnerSet(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockObj := mock.NewMockAuthenticator(mockCtrl)

	handler := Handler{
		auth: mockObj,
	}

	tcs := []struct {
		config          *v0.Config
		CheckSignatures *gomock.Call
		SetOwners       *gomock.Call

		err      error
		response *v0.Response
		desc     string
	}{
		{
			config:          &v0.Config{Witness: &v0.Witness{}, Spec: genAny(t, &v0.OwnerSet{})},
			CheckSignatures: mockObj.EXPECT().CheckSignatures(&v0.Witness{}).Return(false, "", errors.New("err")),

			err:      errors.New("err"),
			response: nil,
			desc:     "checksignatures returning an error should return an error",
		},
		{
			config:          &v0.Config{Witness: &v0.Witness{}, Spec: genAny(t, &v0.OwnerSet{})},
			CheckSignatures: mockObj.EXPECT().CheckSignatures(&v0.Witness{}).Return(false, "signature", nil),

			err:      nil,
			response: v0_handler.WriteMSG(v0.Code_BadRequest, fmt.Sprintf("incorrect signature by %s", "signature")),
			desc:     "checksignatures finding an invalid signature should return Code_Forbidded",
		},
		{
			config: &v0.Config{
				Witness: &v0.Witness{Signatures: map[string][]byte{"correct key": {}}},
				Spec:    genAny(t, &v0.OwnerSet{Owners: []*v0.Owner{{PrimaryPublicKey: "incorrect key", Name: "Jaap"}}})},
			CheckSignatures: mockObj.EXPECT().CheckSignatures(gomock.Any()).Return(true, "", nil),

			err:      nil,
			response: v0_handler.WriteMSG(v0.Code_BadRequest, fmt.Sprintf("not all owners signed the set: %s", "Jaap")),
			desc:     "missing signature returns Code_Forbidded",
		},
		{
			config: &v0.Config{
				Witness: &v0.Witness{Signatures: map[string][]byte{"correct key": {}}},
				Spec:    genAny(t, &v0.OwnerSet{Owners: []*v0.Owner{{PrimaryPublicKey: "correct key", Name: "Jaap"}}})},
			CheckSignatures: mockObj.EXPECT().CheckSignatures(gomock.Any()).Return(true, "", nil),
			SetOwners:       mockObj.EXPECT().SetOwners(context.Background(), gomock.Any()).Return(nil).Times(1),
			err:             nil,
			response:        nil,
			desc:            "correct signatures returns nil, nil",
		},
	}

	for _, tc := range tcs {
		res, err := handler.newOwnerSet(context.Background(), tc.config)
		if err != nil {
			assert.EqualError(t, err, tc.err.Error(), tc.desc)
		}
		assert.Equal(t, tc.response, res, tc.desc)
	}
}

func TestHandler_adjustOwnerSet(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockObj := mock.NewMockAuthenticator(mockCtrl)

	handler := Handler{
		auth: mockObj,
	}

	tcs := []struct {
		config          *v0.Config
		grantRoot       *gomock.Call
		grantRootBackup *gomock.Call
		SetOwners       *gomock.Call

		err      error
		response *v0.Response
		desc     string
	}{
		{
			config:    &v0.Config{Witness: &v0.Witness{}, Spec: genAny(t, &v0.OwnerSet{})},
			grantRoot: mockObj.EXPECT().GrantRoot(&v0.Witness{}).Return(true, nil).Times(1),
			SetOwners: mockObj.EXPECT().SetOwners(context.Background(), &v0.OwnerSet{}).Return(nil),

			err:      nil,
			response: nil,
			desc:     "granting direct root should set owners",
		},
		{
			config:          &v0.Config{Witness: &v0.Witness{}, Spec: genAny(t, &v0.OwnerSet{})},
			grantRoot:       mockObj.EXPECT().GrantRoot(&v0.Witness{}).Return(false, v0_handler.ErrBackupKeyPresent).Times(1),
			grantRootBackup: mockObj.EXPECT().GrantBackupRoot(&v0.Witness{}).Return(true, nil).Times(1),
			SetOwners:       mockObj.EXPECT().SetOwners(context.Background(), &v0.OwnerSet{}).Return(nil),

			err:      nil,
			response: nil,
			desc:     "granting backup root should set owners",
		},
		{
			config:          &v0.Config{Witness: &v0.Witness{}, Spec: genAny(t, &v0.OwnerSet{})},
			grantRoot:       mockObj.EXPECT().GrantRoot(&v0.Witness{}).Return(false, v0_handler.ErrBackupKeyPresent).Times(1),
			grantRootBackup: mockObj.EXPECT().GrantBackupRoot(&v0.Witness{}).Return(false, nil).Times(1),

			err:      nil,
			response: v0_handler.WriteMSG(v0.Code_Forbidded, "insufficient quorum for root op"),
			desc:     "granting no root and no backup root should return code forbidden",
		},
	}

	for _, tc := range tcs {
		res, err := handler.adjustOwnerSet(context.Background(), tc.config)
		assert.Equal(t, tc.err, err, tc.desc)
		assert.Equal(t, tc.response, res, tc.desc)
	}
}
