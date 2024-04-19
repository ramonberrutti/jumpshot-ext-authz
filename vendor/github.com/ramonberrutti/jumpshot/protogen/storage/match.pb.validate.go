// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: storage/match.proto

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

	commonv1 "github.com/ramonberrutti/jumpshot/protogen/common/v1"
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

	_ = commonv1.TournamentMatch_State(0)
)

// Validate checks the field values on TournamentMatch with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *TournamentMatch) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on TournamentMatch with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// TournamentMatchMultiError, or nil if none found.
func (m *TournamentMatch) ValidateAll() error {
	return m.validate(true)
}

func (m *TournamentMatch) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	// no validation rules for TournamentId

	// no validation rules for PhaseId

	// no validation rules for MatchId

	for idx, item := range m.GetParticipants() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, TournamentMatchValidationError{
						field:  fmt.Sprintf("Participants[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, TournamentMatchValidationError{
						field:  fmt.Sprintf("Participants[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return TournamentMatchValidationError{
					field:  fmt.Sprintf("Participants[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	// no validation rules for Version

	// no validation rules for Synced

	// no validation rules for State

	if all {
		switch v := interface{}(m.GetStartTime()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, TournamentMatchValidationError{
					field:  "StartTime",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, TournamentMatchValidationError{
					field:  "StartTime",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetStartTime()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return TournamentMatchValidationError{
				field:  "StartTime",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetEndTime()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, TournamentMatchValidationError{
					field:  "EndTime",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, TournamentMatchValidationError{
					field:  "EndTime",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetEndTime()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return TournamentMatchValidationError{
				field:  "EndTime",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	for idx, item := range m.GetMessages() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, TournamentMatchValidationError{
						field:  fmt.Sprintf("Messages[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, TournamentMatchValidationError{
						field:  fmt.Sprintf("Messages[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return TournamentMatchValidationError{
					field:  fmt.Sprintf("Messages[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if all {
		switch v := interface{}(m.GetCreateTime()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, TournamentMatchValidationError{
					field:  "CreateTime",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, TournamentMatchValidationError{
					field:  "CreateTime",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetCreateTime()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return TournamentMatchValidationError{
				field:  "CreateTime",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetUpdateTime()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, TournamentMatchValidationError{
					field:  "UpdateTime",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, TournamentMatchValidationError{
					field:  "UpdateTime",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetUpdateTime()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return TournamentMatchValidationError{
				field:  "UpdateTime",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetMatchConfig()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, TournamentMatchValidationError{
					field:  "MatchConfig",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, TournamentMatchValidationError{
					field:  "MatchConfig",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetMatchConfig()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return TournamentMatchValidationError{
				field:  "MatchConfig",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	// no validation rules for ParticipantsMode

	if len(errors) > 0 {
		return TournamentMatchMultiError(errors)
	}

	return nil
}

// TournamentMatchMultiError is an error wrapping multiple validation errors
// returned by TournamentMatch.ValidateAll() if the designated constraints
// aren't met.
type TournamentMatchMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m TournamentMatchMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m TournamentMatchMultiError) AllErrors() []error { return m }

// TournamentMatchValidationError is the validation error returned by
// TournamentMatch.Validate if the designated constraints aren't met.
type TournamentMatchValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e TournamentMatchValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e TournamentMatchValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e TournamentMatchValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e TournamentMatchValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e TournamentMatchValidationError) ErrorName() string { return "TournamentMatchValidationError" }

// Error satisfies the builtin error interface
func (e TournamentMatchValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sTournamentMatch.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = TournamentMatchValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = TournamentMatchValidationError{}

// Validate checks the field values on TournamentMatch_Messages with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *TournamentMatch_Messages) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on TournamentMatch_Messages with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// TournamentMatch_MessagesMultiError, or nil if none found.
func (m *TournamentMatch_Messages) ValidateAll() error {
	return m.validate(true)
}

func (m *TournamentMatch_Messages) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	// no validation rules for CreateBy

	// no validation rules for TeamBy

	// no validation rules for Message

	if len(errors) > 0 {
		return TournamentMatch_MessagesMultiError(errors)
	}

	return nil
}

// TournamentMatch_MessagesMultiError is an error wrapping multiple validation
// errors returned by TournamentMatch_Messages.ValidateAll() if the designated
// constraints aren't met.
type TournamentMatch_MessagesMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m TournamentMatch_MessagesMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m TournamentMatch_MessagesMultiError) AllErrors() []error { return m }

// TournamentMatch_MessagesValidationError is the validation error returned by
// TournamentMatch_Messages.Validate if the designated constraints aren't met.
type TournamentMatch_MessagesValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e TournamentMatch_MessagesValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e TournamentMatch_MessagesValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e TournamentMatch_MessagesValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e TournamentMatch_MessagesValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e TournamentMatch_MessagesValidationError) ErrorName() string {
	return "TournamentMatch_MessagesValidationError"
}

// Error satisfies the builtin error interface
func (e TournamentMatch_MessagesValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sTournamentMatch_Messages.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = TournamentMatch_MessagesValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = TournamentMatch_MessagesValidationError{}
