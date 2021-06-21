// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: api/ocp-resume-api/ocp-resume-api.proto

package ocp_resume_api

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

// Validate checks the field values on CreateResumeV1Request with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *CreateResumeV1Request) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for DocumentId

	return nil
}

// CreateResumeV1RequestValidationError is the validation error returned by
// CreateResumeV1Request.Validate if the designated constraints aren't met.
type CreateResumeV1RequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CreateResumeV1RequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CreateResumeV1RequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CreateResumeV1RequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CreateResumeV1RequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CreateResumeV1RequestValidationError) ErrorName() string {
	return "CreateResumeV1RequestValidationError"
}

// Error satisfies the builtin error interface
func (e CreateResumeV1RequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCreateResumeV1Request.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CreateResumeV1RequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CreateResumeV1RequestValidationError{}

// Validate checks the field values on CreateResumeV1Response with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *CreateResumeV1Response) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for ResumeId

	return nil
}

// CreateResumeV1ResponseValidationError is the validation error returned by
// CreateResumeV1Response.Validate if the designated constraints aren't met.
type CreateResumeV1ResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CreateResumeV1ResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CreateResumeV1ResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CreateResumeV1ResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CreateResumeV1ResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CreateResumeV1ResponseValidationError) ErrorName() string {
	return "CreateResumeV1ResponseValidationError"
}

// Error satisfies the builtin error interface
func (e CreateResumeV1ResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCreateResumeV1Response.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CreateResumeV1ResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CreateResumeV1ResponseValidationError{}

// Validate checks the field values on DescribeResumeV1Request with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *DescribeResumeV1Request) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for ResumeId

	return nil
}

// DescribeResumeV1RequestValidationError is the validation error returned by
// DescribeResumeV1Request.Validate if the designated constraints aren't met.
type DescribeResumeV1RequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DescribeResumeV1RequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DescribeResumeV1RequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DescribeResumeV1RequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DescribeResumeV1RequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DescribeResumeV1RequestValidationError) ErrorName() string {
	return "DescribeResumeV1RequestValidationError"
}

// Error satisfies the builtin error interface
func (e DescribeResumeV1RequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDescribeResumeV1Request.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DescribeResumeV1RequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DescribeResumeV1RequestValidationError{}

// Validate checks the field values on DescribeResumeV1Response with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *DescribeResumeV1Response) Validate() error {
	if m == nil {
		return nil
	}

	if v, ok := interface{}(m.GetResume()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return DescribeResumeV1ResponseValidationError{
				field:  "Resume",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// DescribeResumeV1ResponseValidationError is the validation error returned by
// DescribeResumeV1Response.Validate if the designated constraints aren't met.
type DescribeResumeV1ResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DescribeResumeV1ResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DescribeResumeV1ResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DescribeResumeV1ResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DescribeResumeV1ResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DescribeResumeV1ResponseValidationError) ErrorName() string {
	return "DescribeResumeV1ResponseValidationError"
}

// Error satisfies the builtin error interface
func (e DescribeResumeV1ResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDescribeResumeV1Response.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DescribeResumeV1ResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DescribeResumeV1ResponseValidationError{}

// Validate checks the field values on ListResumesV1Request with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *ListResumesV1Request) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Limit

	// no validation rules for Offset

	return nil
}

// ListResumesV1RequestValidationError is the validation error returned by
// ListResumesV1Request.Validate if the designated constraints aren't met.
type ListResumesV1RequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ListResumesV1RequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ListResumesV1RequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ListResumesV1RequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ListResumesV1RequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ListResumesV1RequestValidationError) ErrorName() string {
	return "ListResumesV1RequestValidationError"
}

// Error satisfies the builtin error interface
func (e ListResumesV1RequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sListResumesV1Request.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ListResumesV1RequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ListResumesV1RequestValidationError{}

// Validate checks the field values on ListResumesV1Response with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *ListResumesV1Response) Validate() error {
	if m == nil {
		return nil
	}

	for idx, item := range m.GetResumes() {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return ListResumesV1ResponseValidationError{
					field:  fmt.Sprintf("Resumes[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	return nil
}

// ListResumesV1ResponseValidationError is the validation error returned by
// ListResumesV1Response.Validate if the designated constraints aren't met.
type ListResumesV1ResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ListResumesV1ResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ListResumesV1ResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ListResumesV1ResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ListResumesV1ResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ListResumesV1ResponseValidationError) ErrorName() string {
	return "ListResumesV1ResponseValidationError"
}

// Error satisfies the builtin error interface
func (e ListResumesV1ResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sListResumesV1Response.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ListResumesV1ResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ListResumesV1ResponseValidationError{}

// Validate checks the field values on RemoveResumeV1Request with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *RemoveResumeV1Request) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for ResumeId

	return nil
}

// RemoveResumeV1RequestValidationError is the validation error returned by
// RemoveResumeV1Request.Validate if the designated constraints aren't met.
type RemoveResumeV1RequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e RemoveResumeV1RequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e RemoveResumeV1RequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e RemoveResumeV1RequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e RemoveResumeV1RequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e RemoveResumeV1RequestValidationError) ErrorName() string {
	return "RemoveResumeV1RequestValidationError"
}

// Error satisfies the builtin error interface
func (e RemoveResumeV1RequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sRemoveResumeV1Request.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = RemoveResumeV1RequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = RemoveResumeV1RequestValidationError{}

// Validate checks the field values on RemoveResumeV1Response with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *RemoveResumeV1Response) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Found

	return nil
}

// RemoveResumeV1ResponseValidationError is the validation error returned by
// RemoveResumeV1Response.Validate if the designated constraints aren't met.
type RemoveResumeV1ResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e RemoveResumeV1ResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e RemoveResumeV1ResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e RemoveResumeV1ResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e RemoveResumeV1ResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e RemoveResumeV1ResponseValidationError) ErrorName() string {
	return "RemoveResumeV1ResponseValidationError"
}

// Error satisfies the builtin error interface
func (e RemoveResumeV1ResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sRemoveResumeV1Response.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = RemoveResumeV1ResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = RemoveResumeV1ResponseValidationError{}

// Validate checks the field values on Resume with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *Resume) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Id

	// no validation rules for DocumentId

	return nil
}

// ResumeValidationError is the validation error returned by Resume.Validate if
// the designated constraints aren't met.
type ResumeValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ResumeValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ResumeValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ResumeValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ResumeValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ResumeValidationError) ErrorName() string { return "ResumeValidationError" }

// Error satisfies the builtin error interface
func (e ResumeValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sResume.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ResumeValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ResumeValidationError{}