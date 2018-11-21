// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: ownerset.proto

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

// Validate checks the field values on OwnerSet with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *OwnerSet) Validate() error {
	if m == nil {
		return nil
	}

	if m.GetQuorum() <= 0 {
		return OwnerSetValidationError{
			field:  "Quorum",
			reason: "value must be greater than 0",
		}
	}

	if len(m.GetOwners()) < 1 {
		return OwnerSetValidationError{
			field:  "Owners",
			reason: "value must contain at least 1 item(s)",
		}
	}

	for idx, item := range m.GetOwners() {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return OwnerSetValidationError{
					field:  fmt.Sprintf("Owners[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	return nil
}

// OwnerSetValidationError is the validation error returned by
// OwnerSet.Validate if the designated constraints aren't met.
type OwnerSetValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e OwnerSetValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e OwnerSetValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e OwnerSetValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e OwnerSetValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e OwnerSetValidationError) ErrorName() string { return "OwnerSetValidationError" }

// Error satisfies the builtin error interface
func (e OwnerSetValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sOwnerSet.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = OwnerSetValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = OwnerSetValidationError{}
