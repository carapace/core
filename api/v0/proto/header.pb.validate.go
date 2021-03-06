// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: header.proto

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

// Validate checks the field values on Header with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *Header) Validate() error {
	if m == nil {
		return nil
	}

	if utf8.RuneCountInString(m.GetApiVersion()) < 2 {
		return HeaderValidationError{
			field:  "ApiVersion",
			reason: "value length must be at least 2 runes",
		}
	}

	if !_Header_ApiVersion_Pattern.MatchString(m.GetApiVersion()) {
		return HeaderValidationError{
			field:  "ApiVersion",
			reason: "value does not match regex pattern \"^v\"",
		}
	}

	if utf8.RuneCountInString(m.GetKind()) < 1 {
		return HeaderValidationError{
			field:  "Kind",
			reason: "value length must be at least 1 runes",
		}
	}

	if m.GetIncrement() <= 0 {
		return HeaderValidationError{
			field:  "Increment",
			reason: "value must be greater than 0",
		}
	}

	return nil
}

// HeaderValidationError is the validation error returned by Header.Validate if
// the designated constraints aren't met.
type HeaderValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e HeaderValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e HeaderValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e HeaderValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e HeaderValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e HeaderValidationError) ErrorName() string { return "HeaderValidationError" }

// Error satisfies the builtin error interface
func (e HeaderValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sHeader.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = HeaderValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = HeaderValidationError{}

var _Header_ApiVersion_Pattern = regexp.MustCompile("^v")
