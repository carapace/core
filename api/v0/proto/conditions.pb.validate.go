// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: conditions.proto

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

// Validate checks the field values on AuthLevelGreaterArg with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *AuthLevelGreaterArg) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Level

	return nil
}

// AuthLevelGreaterArgValidationError is the validation error returned by
// AuthLevelGreaterArg.Validate if the designated constraints aren't met.
type AuthLevelGreaterArgValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e AuthLevelGreaterArgValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e AuthLevelGreaterArgValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e AuthLevelGreaterArgValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e AuthLevelGreaterArgValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e AuthLevelGreaterArgValidationError) ErrorName() string {
	return "AuthLevelGreaterArgValidationError"
}

// Error satisfies the builtin error interface
func (e AuthLevelGreaterArgValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sAuthLevelGreaterArg.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = AuthLevelGreaterArgValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = AuthLevelGreaterArgValidationError{}

// Validate checks the field values on UserOwnsArg with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *UserOwnsArg) Validate() error {
	if m == nil {
		return nil
	}

	return nil
}

// UserOwnsArgValidationError is the validation error returned by
// UserOwnsArg.Validate if the designated constraints aren't met.
type UserOwnsArgValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UserOwnsArgValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UserOwnsArgValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UserOwnsArgValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UserOwnsArgValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UserOwnsArgValidationError) ErrorName() string { return "UserOwnsArgValidationError" }

// Error satisfies the builtin error interface
func (e UserOwnsArgValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUserOwnsArg.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UserOwnsArgValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UserOwnsArgValidationError{}

// Validate checks the field values on InSetsArg with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *InSetsArg) Validate() error {
	if m == nil {
		return nil
	}

	return nil
}

// InSetsArgValidationError is the validation error returned by
// InSetsArg.Validate if the designated constraints aren't met.
type InSetsArgValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e InSetsArgValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e InSetsArgValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e InSetsArgValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e InSetsArgValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e InSetsArgValidationError) ErrorName() string { return "InSetsArgValidationError" }

// Error satisfies the builtin error interface
func (e InSetsArgValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sInSetsArg.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = InSetsArgValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = InSetsArgValidationError{}

// Validate checks the field values on MultiSigArg with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *MultiSigArg) Validate() error {
	if m == nil {
		return nil
	}

	return nil
}

// MultiSigArgValidationError is the validation error returned by
// MultiSigArg.Validate if the designated constraints aren't met.
type MultiSigArgValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e MultiSigArgValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e MultiSigArgValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e MultiSigArgValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e MultiSigArgValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e MultiSigArgValidationError) ErrorName() string { return "MultiSigArgValidationError" }

// Error satisfies the builtin error interface
func (e MultiSigArgValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sMultiSigArg.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = MultiSigArgValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = MultiSigArgValidationError{}
