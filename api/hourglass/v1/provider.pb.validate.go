// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: api/hourglass/v1/provider.proto

package v1

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

	"google.golang.org/protobuf/types/known/anypb"
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
	_ = anypb.Any{}
)

// Validate checks the field values on ProviderRegisterRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *ProviderRegisterRequest) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Capacity

	// no validation rules for Name

	// no validation rules for MinClients

	return nil
}

// ProviderRegisterRequestValidationError is the validation error returned by
// ProviderRegisterRequest.Validate if the designated constraints aren't met.
type ProviderRegisterRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ProviderRegisterRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ProviderRegisterRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ProviderRegisterRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ProviderRegisterRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ProviderRegisterRequestValidationError) ErrorName() string {
	return "ProviderRegisterRequestValidationError"
}

// Error satisfies the builtin error interface
func (e ProviderRegisterRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sProviderRegisterRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ProviderRegisterRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ProviderRegisterRequestValidationError{}

// Validate checks the field values on ProviderInfo with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *ProviderInfo) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Id

	// no validation rules for Capacity

	// no validation rules for Name

	// no validation rules for MinClients

	return nil
}

// ProviderInfoValidationError is the validation error returned by
// ProviderInfo.Validate if the designated constraints aren't met.
type ProviderInfoValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ProviderInfoValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ProviderInfoValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ProviderInfoValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ProviderInfoValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ProviderInfoValidationError) ErrorName() string { return "ProviderInfoValidationError" }

// Error satisfies the builtin error interface
func (e ProviderInfoValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sProviderInfo.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ProviderInfoValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ProviderInfoValidationError{}

// Validate checks the field values on ProviderDisregisterRequest with the
// rules defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *ProviderDisregisterRequest) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Id

	// no validation rules for Name

	return nil
}

// ProviderDisregisterRequestValidationError is the validation error returned
// by ProviderDisregisterRequest.Validate if the designated constraints aren't met.
type ProviderDisregisterRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ProviderDisregisterRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ProviderDisregisterRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ProviderDisregisterRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ProviderDisregisterRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ProviderDisregisterRequestValidationError) ErrorName() string {
	return "ProviderDisregisterRequestValidationError"
}

// Error satisfies the builtin error interface
func (e ProviderDisregisterRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sProviderDisregisterRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ProviderDisregisterRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ProviderDisregisterRequestValidationError{}
