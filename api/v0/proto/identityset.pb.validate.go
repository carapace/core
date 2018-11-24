// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: identityset.proto

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

// Validate checks the field values on Identity with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *Identity) Validate() error {
	if m == nil {
		return nil
	}

	if utf8.RuneCountInString(m.GetName()) < 1 {
		return IdentityValidationError{
			field:  "Name",
			reason: "value length must be at least 1 runes",
		}
	}

	if _, ok := _Identity_Asset_NotInLookup[m.GetAsset()]; ok {
		return IdentityValidationError{
			field:  "Asset",
			reason: "value must not be in list [0]",
		}
	}

	if _, ok := Asset_name[int32(m.GetAsset())]; !ok {
		return IdentityValidationError{
			field:  "Asset",
			reason: "value must be one of the defined enum values",
		}
	}

	if len(m.GetAccess()) < 1 {
		return IdentityValidationError{
			field:  "Access",
			reason: "value must contain at least 1 item(s)",
		}
	}

	for idx, item := range m.GetAccess() {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return IdentityValidationError{
					field:  fmt.Sprintf("Access[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	switch m.State.(type) {

	case *Identity_Enabled:

		if v, ok := interface{}(m.GetEnabled()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return IdentityValidationError{
					field:  "Enabled",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	case *Identity_Disabled:

		if v, ok := interface{}(m.GetDisabled()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return IdentityValidationError{
					field:  "Disabled",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	return nil
}

// IdentityValidationError is the validation error returned by
// Identity.Validate if the designated constraints aren't met.
type IdentityValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e IdentityValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e IdentityValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e IdentityValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e IdentityValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e IdentityValidationError) ErrorName() string { return "IdentityValidationError" }

// Error satisfies the builtin error interface
func (e IdentityValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sIdentity.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = IdentityValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = IdentityValidationError{}

var _Identity_Asset_NotInLookup = map[Asset]struct{}{
	0: {},
}
