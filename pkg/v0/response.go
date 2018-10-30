package v0_handler

import (
	"github.com/carapace/core/api/v0/proto"
)

func writeResponse(code v0.Code, msg string, err error) *v0.Response {
	errmsg := ""
	if err != nil {
		errmsg = err.Error()
	}

	return &v0.Response{
		Code: code,
		MSG:  msg,
		Err:  errmsg,
	}
}

func WriteErr(err error) *v0.Response {
	return writeResponse(v0.Code_Internal, "", err)
}

func WriteMSG(code v0.Code, msg string) *v0.Response {
	return writeResponse(code, msg, nil)
}
