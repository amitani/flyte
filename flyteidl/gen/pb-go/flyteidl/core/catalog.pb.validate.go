// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: flyteidl/core/catalog.proto

package core

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

// define the regex for a UUID once up-front
var _catalog_uuidPattern = regexp.MustCompile("^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}$")

// Validate checks the field values on CatalogArtifactTag with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *CatalogArtifactTag) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for ArtifactId

	// no validation rules for Name

	return nil
}

// CatalogArtifactTagValidationError is the validation error returned by
// CatalogArtifactTag.Validate if the designated constraints aren't met.
type CatalogArtifactTagValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CatalogArtifactTagValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CatalogArtifactTagValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CatalogArtifactTagValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CatalogArtifactTagValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CatalogArtifactTagValidationError) ErrorName() string {
	return "CatalogArtifactTagValidationError"
}

// Error satisfies the builtin error interface
func (e CatalogArtifactTagValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCatalogArtifactTag.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CatalogArtifactTagValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CatalogArtifactTagValidationError{}

// Validate checks the field values on CatalogMetadata with the rules defined
// in the proto definition for this message. If any rules are violated, an
// error is returned.
func (m *CatalogMetadata) Validate() error {
	if m == nil {
		return nil
	}

	if v, ok := interface{}(m.GetDatasetId()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return CatalogMetadataValidationError{
				field:  "DatasetId",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if v, ok := interface{}(m.GetArtifactTag()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return CatalogMetadataValidationError{
				field:  "ArtifactTag",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	switch m.SourceExecution.(type) {

	case *CatalogMetadata_SourceTaskExecution:

		if v, ok := interface{}(m.GetSourceTaskExecution()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return CatalogMetadataValidationError{
					field:  "SourceTaskExecution",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	return nil
}

// CatalogMetadataValidationError is the validation error returned by
// CatalogMetadata.Validate if the designated constraints aren't met.
type CatalogMetadataValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CatalogMetadataValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CatalogMetadataValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CatalogMetadataValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CatalogMetadataValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CatalogMetadataValidationError) ErrorName() string { return "CatalogMetadataValidationError" }

// Error satisfies the builtin error interface
func (e CatalogMetadataValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCatalogMetadata.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CatalogMetadataValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CatalogMetadataValidationError{}

// Validate checks the field values on CatalogReservation with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *CatalogReservation) Validate() error {
	if m == nil {
		return nil
	}

	return nil
}

// CatalogReservationValidationError is the validation error returned by
// CatalogReservation.Validate if the designated constraints aren't met.
type CatalogReservationValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CatalogReservationValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CatalogReservationValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CatalogReservationValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CatalogReservationValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CatalogReservationValidationError) ErrorName() string {
	return "CatalogReservationValidationError"
}

// Error satisfies the builtin error interface
func (e CatalogReservationValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCatalogReservation.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CatalogReservationValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CatalogReservationValidationError{}
