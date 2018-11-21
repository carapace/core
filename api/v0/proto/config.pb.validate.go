// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: config.proto

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

	// no validation rules for ApiVersion

	// no validation rules for Kind

	// no validation rules for Increment

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

// Validate checks the field values on Witness with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *Witness) Validate() error {
	if m == nil {
		return nil
	}

	for idx, item := range m.GetSignatures() {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return WitnessValidationError{
					field:  fmt.Sprintf("Signatures[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	return nil
}

// WitnessValidationError is the validation error returned by Witness.Validate
// if the designated constraints aren't met.
type WitnessValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e WitnessValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e WitnessValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e WitnessValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e WitnessValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e WitnessValidationError) ErrorName() string { return "WitnessValidationError" }

// Error satisfies the builtin error interface
func (e WitnessValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sWitness.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = WitnessValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = WitnessValidationError{}

// Validate checks the field values on Signature with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *Signature) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for R

	// no validation rules for S

	switch m.Key.(type) {

	case *Signature_PrimaryPublicKey:
		// no validation rules for PrimaryPublicKey

	case *Signature_RecoveryPublicKey:
		// no validation rules for RecoveryPublicKey

	}

	return nil
}

// SignatureValidationError is the validation error returned by
// Signature.Validate if the designated constraints aren't met.
type SignatureValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e SignatureValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e SignatureValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e SignatureValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e SignatureValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e SignatureValidationError) ErrorName() string { return "SignatureValidationError" }

// Error satisfies the builtin error interface
func (e SignatureValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sSignature.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = SignatureValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = SignatureValidationError{}

// Validate checks the field values on Config with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *Config) Validate() error {
	if m == nil {
		return nil
	}

	if m.GetHeader() == nil {
		return ConfigValidationError{
			field:  "Header",
			reason: "value is required",
		}
	}

	if v, ok := interface{}(m.GetHeader()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return ConfigValidationError{
				field:  "Header",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if m.GetWitness() == nil {
		return ConfigValidationError{
			field:  "Witness",
			reason: "value is required",
		}
	}

	if v, ok := interface{}(m.GetWitness()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return ConfigValidationError{
				field:  "Witness",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if m.GetSpec() == nil {
		return ConfigValidationError{
			field:  "Spec",
			reason: "value is required",
		}
	}

	if v, ok := interface{}(m.GetSpec()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return ConfigValidationError{
				field:  "Spec",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// ConfigValidationError is the validation error returned by Config.Validate if
// the designated constraints aren't met.
type ConfigValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ConfigValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ConfigValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ConfigValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ConfigValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ConfigValidationError) ErrorName() string { return "ConfigValidationError" }

// Error satisfies the builtin error interface
func (e ConfigValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sConfig.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ConfigValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ConfigValidationError{}