// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: access.proto

package v0

import (
	"bytes"
	"errors"
	"fmt"
	"net"
	"net/mail"
	"net/url"
	"regexp"
	"strings"
	"time"
	"unicode/utf8"

	"github.com/golang/protobuf/ptypes"
)

// ensure the imports are used
var (
	_ = bytes.MinRead
	_ = errors.New("")
	_ = fmt.Print
	_ = utf8.UTFMax
	_ = (*regexp.Regexp)(nil)
	_ = (*strings.Reader)(nil)
	_ = net.IPv4len
	_ = time.Duration(0)
	_ = (*url.URL)(nil)
	_ = (*mail.Address)(nil)
	_ = ptypes.DynamicAny{}
)

// Validate checks the field values on AccessProtocol with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *AccessProtocol) Validate() error {
	if m == nil {
		return nil
	}

	switch m.Method.(type) {

	case *AccessProtocol_AuthLevel:
		// no validation rules for AuthLevel

	case *AccessProtocol_User:
		// no validation rules for User

	case *AccessProtocol_UserSet:
		// no validation rules for UserSet

	}

	return nil
}

// AccessProtocolValidationError is the validation error returned by
// AccessProtocol.Validate if the designated constraints aren't met.
type AccessProtocolValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e AccessProtocolValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e AccessProtocolValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e AccessProtocolValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e AccessProtocolValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e AccessProtocolValidationError) ErrorName() string { return "AccessProtocolValidationError" }

// Error satisfies the builtin error interface
func (e AccessProtocolValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sAccessProtocol.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = AccessProtocolValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = AccessProtocolValidationError{}
