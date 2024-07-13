// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: order/v1/order.proto

package __order

import (
	"bytes"
	"errors"
	"fmt"
	"net"
	"net/mail"
	"net/url"
	"regexp"
	"sort"
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
	_ = sort.Sort
)

// Validate checks the field values on AddOrderRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *AddOrderRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on AddOrderRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// AddOrderRequestMultiError, or nil if none found.
func (m *AddOrderRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *AddOrderRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for OrderId

	// no validation rules for AddresseeId

	if all {
		switch v := interface{}(m.GetShelfLife()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, AddOrderRequestValidationError{
					field:  "ShelfLife",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, AddOrderRequestValidationError{
					field:  "ShelfLife",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetShelfLife()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return AddOrderRequestValidationError{
				field:  "ShelfLife",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if m.Weight != nil {
		// no validation rules for Weight
	}

	if m.Price != nil {
		// no validation rules for Price
	}

	if m.WrapType != nil {
		// no validation rules for WrapType
	}

	if len(errors) > 0 {
		return AddOrderRequestMultiError(errors)
	}

	return nil
}

// AddOrderRequestMultiError is an error wrapping multiple validation errors
// returned by AddOrderRequest.ValidateAll() if the designated constraints
// aren't met.
type AddOrderRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m AddOrderRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m AddOrderRequestMultiError) AllErrors() []error { return m }

// AddOrderRequestValidationError is the validation error returned by
// AddOrderRequest.Validate if the designated constraints aren't met.
type AddOrderRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e AddOrderRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e AddOrderRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e AddOrderRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e AddOrderRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e AddOrderRequestValidationError) ErrorName() string { return "AddOrderRequestValidationError" }

// Error satisfies the builtin error interface
func (e AddOrderRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sAddOrderRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = AddOrderRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = AddOrderRequestValidationError{}

// Validate checks the field values on ReturnToDelivererRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *ReturnToDelivererRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ReturnToDelivererRequest with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// ReturnToDelivererRequestMultiError, or nil if none found.
func (m *ReturnToDelivererRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *ReturnToDelivererRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for OrderId

	if len(errors) > 0 {
		return ReturnToDelivererRequestMultiError(errors)
	}

	return nil
}

// ReturnToDelivererRequestMultiError is an error wrapping multiple validation
// errors returned by ReturnToDelivererRequest.ValidateAll() if the designated
// constraints aren't met.
type ReturnToDelivererRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ReturnToDelivererRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ReturnToDelivererRequestMultiError) AllErrors() []error { return m }

// ReturnToDelivererRequestValidationError is the validation error returned by
// ReturnToDelivererRequest.Validate if the designated constraints aren't met.
type ReturnToDelivererRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ReturnToDelivererRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ReturnToDelivererRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ReturnToDelivererRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ReturnToDelivererRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ReturnToDelivererRequestValidationError) ErrorName() string {
	return "ReturnToDelivererRequestValidationError"
}

// Error satisfies the builtin error interface
func (e ReturnToDelivererRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sReturnToDelivererRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ReturnToDelivererRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ReturnToDelivererRequestValidationError{}

// Validate checks the field values on GiveToAddresseeRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *GiveToAddresseeRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GiveToAddresseeRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// GiveToAddresseeRequestMultiError, or nil if none found.
func (m *GiveToAddresseeRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *GiveToAddresseeRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for OrderId

	// no validation rules for AddresseeId

	if len(errors) > 0 {
		return GiveToAddresseeRequestMultiError(errors)
	}

	return nil
}

// GiveToAddresseeRequestMultiError is an error wrapping multiple validation
// errors returned by GiveToAddresseeRequest.ValidateAll() if the designated
// constraints aren't met.
type GiveToAddresseeRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GiveToAddresseeRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GiveToAddresseeRequestMultiError) AllErrors() []error { return m }

// GiveToAddresseeRequestValidationError is the validation error returned by
// GiveToAddresseeRequest.Validate if the designated constraints aren't met.
type GiveToAddresseeRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GiveToAddresseeRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GiveToAddresseeRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GiveToAddresseeRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GiveToAddresseeRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GiveToAddresseeRequestValidationError) ErrorName() string {
	return "GiveToAddresseeRequestValidationError"
}

// Error satisfies the builtin error interface
func (e GiveToAddresseeRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGiveToAddresseeRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GiveToAddresseeRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GiveToAddresseeRequestValidationError{}

// Validate checks the field values on ListRefundRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *ListRefundRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ListRefundRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// ListRefundRequestMultiError, or nil if none found.
func (m *ListRefundRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *ListRefundRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if m.PageLen != nil {
		// no validation rules for PageLen
	}

	if m.PageNumber != nil {
		// no validation rules for PageNumber
	}

	if len(errors) > 0 {
		return ListRefundRequestMultiError(errors)
	}

	return nil
}

// ListRefundRequestMultiError is an error wrapping multiple validation errors
// returned by ListRefundRequest.ValidateAll() if the designated constraints
// aren't met.
type ListRefundRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ListRefundRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ListRefundRequestMultiError) AllErrors() []error { return m }

// ListRefundRequestValidationError is the validation error returned by
// ListRefundRequest.Validate if the designated constraints aren't met.
type ListRefundRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ListRefundRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ListRefundRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ListRefundRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ListRefundRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ListRefundRequestValidationError) ErrorName() string {
	return "ListRefundRequestValidationError"
}

// Error satisfies the builtin error interface
func (e ListRefundRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sListRefundRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ListRefundRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ListRefundRequestValidationError{}

// Validate checks the field values on ListOrderRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *ListOrderRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ListOrderRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// ListOrderRequestMultiError, or nil if none found.
func (m *ListOrderRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *ListOrderRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for ClientId

	if m.CountOrders != nil {
		// no validation rules for CountOrders
	}

	if len(errors) > 0 {
		return ListOrderRequestMultiError(errors)
	}

	return nil
}

// ListOrderRequestMultiError is an error wrapping multiple validation errors
// returned by ListOrderRequest.ValidateAll() if the designated constraints
// aren't met.
type ListOrderRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ListOrderRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ListOrderRequestMultiError) AllErrors() []error { return m }

// ListOrderRequestValidationError is the validation error returned by
// ListOrderRequest.Validate if the designated constraints aren't met.
type ListOrderRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ListOrderRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ListOrderRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ListOrderRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ListOrderRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ListOrderRequestValidationError) ErrorName() string { return "ListOrderRequestValidationError" }

// Error satisfies the builtin error interface
func (e ListOrderRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sListOrderRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ListOrderRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ListOrderRequestValidationError{}

// Validate checks the field values on ReturnFromAddresseeRequest with the
// rules defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *ReturnFromAddresseeRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ReturnFromAddresseeRequest with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// ReturnFromAddresseeRequestMultiError, or nil if none found.
func (m *ReturnFromAddresseeRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *ReturnFromAddresseeRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return ReturnFromAddresseeRequestMultiError(errors)
	}

	return nil
}

// ReturnFromAddresseeRequestMultiError is an error wrapping multiple
// validation errors returned by ReturnFromAddresseeRequest.ValidateAll() if
// the designated constraints aren't met.
type ReturnFromAddresseeRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ReturnFromAddresseeRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ReturnFromAddresseeRequestMultiError) AllErrors() []error { return m }

// ReturnFromAddresseeRequestValidationError is the validation error returned
// by ReturnFromAddresseeRequest.Validate if the designated constraints aren't met.
type ReturnFromAddresseeRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ReturnFromAddresseeRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ReturnFromAddresseeRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ReturnFromAddresseeRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ReturnFromAddresseeRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ReturnFromAddresseeRequestValidationError) ErrorName() string {
	return "ReturnFromAddresseeRequestValidationError"
}

// Error satisfies the builtin error interface
func (e ReturnFromAddresseeRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sReturnFromAddresseeRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ReturnFromAddresseeRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ReturnFromAddresseeRequestValidationError{}

// Validate checks the field values on ListOrderResponse with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *ListOrderResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ListOrderResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// ListOrderResponseMultiError, or nil if none found.
func (m *ListOrderResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *ListOrderResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return ListOrderResponseMultiError(errors)
	}

	return nil
}

// ListOrderResponseMultiError is an error wrapping multiple validation errors
// returned by ListOrderResponse.ValidateAll() if the designated constraints
// aren't met.
type ListOrderResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ListOrderResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ListOrderResponseMultiError) AllErrors() []error { return m }

// ListOrderResponseValidationError is the validation error returned by
// ListOrderResponse.Validate if the designated constraints aren't met.
type ListOrderResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ListOrderResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ListOrderResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ListOrderResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ListOrderResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ListOrderResponseValidationError) ErrorName() string {
	return "ListOrderResponseValidationError"
}

// Error satisfies the builtin error interface
func (e ListOrderResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sListOrderResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ListOrderResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ListOrderResponseValidationError{}

// Validate checks the field values on ListRefundResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *ListRefundResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ListRefundResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// ListRefundResponseMultiError, or nil if none found.
func (m *ListRefundResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *ListRefundResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return ListRefundResponseMultiError(errors)
	}

	return nil
}

// ListRefundResponseMultiError is an error wrapping multiple validation errors
// returned by ListRefundResponse.ValidateAll() if the designated constraints
// aren't met.
type ListRefundResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ListRefundResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ListRefundResponseMultiError) AllErrors() []error { return m }

// ListRefundResponseValidationError is the validation error returned by
// ListRefundResponse.Validate if the designated constraints aren't met.
type ListRefundResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ListRefundResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ListRefundResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ListRefundResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ListRefundResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ListRefundResponseValidationError) ErrorName() string {
	return "ListRefundResponseValidationError"
}

// Error satisfies the builtin error interface
func (e ListRefundResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sListRefundResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ListRefundResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ListRefundResponseValidationError{}
