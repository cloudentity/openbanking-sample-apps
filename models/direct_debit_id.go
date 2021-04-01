// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/validate"
)

// DirectDebitID A unique and immutable identifier used to identify the direct debit resource. This identifier has no meaning to the account owner.
//
// swagger:model DirectDebitId
type DirectDebitID string

// Validate validates this direct debit Id
func (m DirectDebitID) Validate(formats strfmt.Registry) error {
	var res []error

	if err := validate.MinLength("", "body", string(m), 1); err != nil {
		return err
	}

	if err := validate.MaxLength("", "body", string(m), 40); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this direct debit Id based on context it is used
func (m DirectDebitID) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
