// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: storage/oauth2.proto

package storage

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

// Validate checks the field values on OAuth2Client with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *OAuth2Client) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on OAuth2Client with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in OAuth2ClientMultiError, or
// nil if none found.
func (m *OAuth2Client) ValidateAll() error {
	return m.validate(true)
}

func (m *OAuth2Client) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	// no validation rules for Secret

	if all {
		switch v := interface{}(m.GetCreatedAt()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, OAuth2ClientValidationError{
					field:  "CreatedAt",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, OAuth2ClientValidationError{
					field:  "CreatedAt",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetCreatedAt()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return OAuth2ClientValidationError{
				field:  "CreatedAt",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetUpdatedAt()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, OAuth2ClientValidationError{
					field:  "UpdatedAt",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, OAuth2ClientValidationError{
					field:  "UpdatedAt",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetUpdatedAt()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return OAuth2ClientValidationError{
				field:  "UpdatedAt",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	// no validation rules for OwnerId

	for idx, item := range m.GetCollaborators() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, OAuth2ClientValidationError{
						field:  fmt.Sprintf("Collaborators[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, OAuth2ClientValidationError{
						field:  fmt.Sprintf("Collaborators[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return OAuth2ClientValidationError{
					field:  fmt.Sprintf("Collaborators[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	// no validation rules for Version

	// no validation rules for Name

	// no validation rules for Description

	if len(errors) > 0 {
		return OAuth2ClientMultiError(errors)
	}

	return nil
}

// OAuth2ClientMultiError is an error wrapping multiple validation errors
// returned by OAuth2Client.ValidateAll() if the designated constraints aren't met.
type OAuth2ClientMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m OAuth2ClientMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m OAuth2ClientMultiError) AllErrors() []error { return m }

// OAuth2ClientValidationError is the validation error returned by
// OAuth2Client.Validate if the designated constraints aren't met.
type OAuth2ClientValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e OAuth2ClientValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e OAuth2ClientValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e OAuth2ClientValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e OAuth2ClientValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e OAuth2ClientValidationError) ErrorName() string { return "OAuth2ClientValidationError" }

// Error satisfies the builtin error interface
func (e OAuth2ClientValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sOAuth2Client.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = OAuth2ClientValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = OAuth2ClientValidationError{}

// Validate checks the field values on OAuth2AuthorizationCode with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *OAuth2AuthorizationCode) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on OAuth2AuthorizationCode with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// OAuth2AuthorizationCodeMultiError, or nil if none found.
func (m *OAuth2AuthorizationCode) ValidateAll() error {
	return m.validate(true)
}

func (m *OAuth2AuthorizationCode) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	// no validation rules for ClientId

	// no validation rules for UserId

	// no validation rules for RedirectUri

	if all {
		switch v := interface{}(m.GetCreatedAt()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, OAuth2AuthorizationCodeValidationError{
					field:  "CreatedAt",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, OAuth2AuthorizationCodeValidationError{
					field:  "CreatedAt",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetCreatedAt()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return OAuth2AuthorizationCodeValidationError{
				field:  "CreatedAt",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetUpdatedAt()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, OAuth2AuthorizationCodeValidationError{
					field:  "UpdatedAt",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, OAuth2AuthorizationCodeValidationError{
					field:  "UpdatedAt",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetUpdatedAt()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return OAuth2AuthorizationCodeValidationError{
				field:  "UpdatedAt",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetExpiresAt()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, OAuth2AuthorizationCodeValidationError{
					field:  "ExpiresAt",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, OAuth2AuthorizationCodeValidationError{
					field:  "ExpiresAt",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetExpiresAt()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return OAuth2AuthorizationCodeValidationError{
				field:  "ExpiresAt",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	// no validation rules for Version

	// no validation rules for CodeChallenge

	// no validation rules for CodeChallengeMethod

	if len(errors) > 0 {
		return OAuth2AuthorizationCodeMultiError(errors)
	}

	return nil
}

// OAuth2AuthorizationCodeMultiError is an error wrapping multiple validation
// errors returned by OAuth2AuthorizationCode.ValidateAll() if the designated
// constraints aren't met.
type OAuth2AuthorizationCodeMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m OAuth2AuthorizationCodeMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m OAuth2AuthorizationCodeMultiError) AllErrors() []error { return m }

// OAuth2AuthorizationCodeValidationError is the validation error returned by
// OAuth2AuthorizationCode.Validate if the designated constraints aren't met.
type OAuth2AuthorizationCodeValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e OAuth2AuthorizationCodeValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e OAuth2AuthorizationCodeValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e OAuth2AuthorizationCodeValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e OAuth2AuthorizationCodeValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e OAuth2AuthorizationCodeValidationError) ErrorName() string {
	return "OAuth2AuthorizationCodeValidationError"
}

// Error satisfies the builtin error interface
func (e OAuth2AuthorizationCodeValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sOAuth2AuthorizationCode.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = OAuth2AuthorizationCodeValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = OAuth2AuthorizationCodeValidationError{}

// Validate checks the field values on OAuth2Client_Collaborator with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *OAuth2Client_Collaborator) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on OAuth2Client_Collaborator with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// OAuth2Client_CollaboratorMultiError, or nil if none found.
func (m *OAuth2Client_Collaborator) ValidateAll() error {
	return m.validate(true)
}

func (m *OAuth2Client_Collaborator) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	// no validation rules for Role

	if len(errors) > 0 {
		return OAuth2Client_CollaboratorMultiError(errors)
	}

	return nil
}

// OAuth2Client_CollaboratorMultiError is an error wrapping multiple validation
// errors returned by OAuth2Client_Collaborator.ValidateAll() if the
// designated constraints aren't met.
type OAuth2Client_CollaboratorMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m OAuth2Client_CollaboratorMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m OAuth2Client_CollaboratorMultiError) AllErrors() []error { return m }

// OAuth2Client_CollaboratorValidationError is the validation error returned by
// OAuth2Client_Collaborator.Validate if the designated constraints aren't met.
type OAuth2Client_CollaboratorValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e OAuth2Client_CollaboratorValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e OAuth2Client_CollaboratorValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e OAuth2Client_CollaboratorValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e OAuth2Client_CollaboratorValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e OAuth2Client_CollaboratorValidationError) ErrorName() string {
	return "OAuth2Client_CollaboratorValidationError"
}

// Error satisfies the builtin error interface
func (e OAuth2Client_CollaboratorValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sOAuth2Client_Collaborator.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = OAuth2Client_CollaboratorValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = OAuth2Client_CollaboratorValidationError{}
