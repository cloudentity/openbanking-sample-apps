// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/validate"
)

// OBAccountStatus1Code Specifies the status of account resource in code form.
//
// swagger:model OBAccountStatus1Code
type OBAccountStatus1Code string

func NewOBAccountStatus1Code(value OBAccountStatus1Code) *OBAccountStatus1Code {
	v := value
	return &v
}

const (

	// OBAccountStatus1CodeDeleted captures enum value "Deleted"
	OBAccountStatus1CodeDeleted OBAccountStatus1Code = "Deleted"

	// OBAccountStatus1CodeDisabled captures enum value "Disabled"
	OBAccountStatus1CodeDisabled OBAccountStatus1Code = "Disabled"

	// OBAccountStatus1CodeEnabled captures enum value "Enabled"
	OBAccountStatus1CodeEnabled OBAccountStatus1Code = "Enabled"

	// OBAccountStatus1CodePending captures enum value "Pending"
	OBAccountStatus1CodePending OBAccountStatus1Code = "Pending"

	// OBAccountStatus1CodeProForma captures enum value "ProForma"
	OBAccountStatus1CodeProForma OBAccountStatus1Code = "ProForma"
)

// for schema
var oBAccountStatus1CodeEnum []interface{}

func init() {
	var res []OBAccountStatus1Code
	if err := json.Unmarshal([]byte(`["Deleted","Disabled","Enabled","Pending","ProForma"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		oBAccountStatus1CodeEnum = append(oBAccountStatus1CodeEnum, v)
	}
}

func (m OBAccountStatus1Code) validateOBAccountStatus1CodeEnum(path, location string, value OBAccountStatus1Code) error {
	if err := validate.EnumCase(path, location, value, oBAccountStatus1CodeEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this o b account status1 code
func (m OBAccountStatus1Code) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateOBAccountStatus1CodeEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this o b account status1 code based on context it is used
func (m OBAccountStatus1Code) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
