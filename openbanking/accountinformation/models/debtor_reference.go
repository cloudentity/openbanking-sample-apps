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

// DebtorReference A reference value provided by the PSU to the PISP while setting up the scheduled payment.
//
// swagger:model DebtorReference
type DebtorReference string

// Validate validates this debtor reference
func (m DebtorReference) Validate(formats strfmt.Registry) error {
	var res []error

	if err := validate.MinLength("", "body", string(m), 1); err != nil {
		return err
	}

	if err := validate.MaxLength("", "body", string(m), 35); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this debtor reference based on context it is used
func (m DebtorReference) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
