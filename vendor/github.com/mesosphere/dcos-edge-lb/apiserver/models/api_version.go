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

// APIVersion API version
// swagger:model APIVersion
type APIVersion string

const (
	// APIVersionV1 captures enum value "V1"
	APIVersionV1 APIVersion = "V1"
	// APIVersionV2 captures enum value "V2"
	APIVersionV2 APIVersion = "V2"
)

// for schema
var apiVersionEnum []interface{}

func init() {
	var res []APIVersion
	if err := json.Unmarshal([]byte(`["V1","V2"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		apiVersionEnum = append(apiVersionEnum, v)
	}
}

func (m APIVersion) validateAPIVersionEnum(path, location string, value APIVersion) error {
	if err := validate.Enum(path, location, value, apiVersionEnum); err != nil {
		return err
	}
	return nil
}

// Validate validates this API version
func (m APIVersion) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateAPIVersionEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
