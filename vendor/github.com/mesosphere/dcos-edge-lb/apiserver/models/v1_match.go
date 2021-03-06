// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/validate"
)

// V1Match v1 match
// swagger:model V1Match
type V1Match string

const (
	// V1MatchEXACT captures enum value "EXACT"
	V1MatchEXACT V1Match = "EXACT"
	// V1MatchREGEX captures enum value "REGEX"
	V1MatchREGEX V1Match = "REGEX"
)

// for schema
var v1MatchEnum []interface{}

func init() {
	var res []V1Match
	if err := json.Unmarshal([]byte(`["EXACT","REGEX"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		v1MatchEnum = append(v1MatchEnum, v)
	}
}

func (m V1Match) validateV1MatchEnum(path, location string, value V1Match) error {
	if err := validate.Enum(path, location, value, v1MatchEnum); err != nil {
		return err
	}
	return nil
}

// Validate validates this v1 match
func (m V1Match) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateV1MatchEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
