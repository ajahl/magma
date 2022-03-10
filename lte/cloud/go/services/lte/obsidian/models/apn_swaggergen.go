// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// Apn apn
//
// swagger:model apn
type Apn struct {

	// apn configuration
	// Required: true
	ApnConfiguration *ApnConfiguration `json:"apn_configuration"`

	// apn name
	// Required: true
	ApnName ApnName `json:"apn_name" magma_alt_name:"service_selection"`
}

// Validate validates this apn
func (m *Apn) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateApnConfiguration(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateApnName(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Apn) validateApnConfiguration(formats strfmt.Registry) error {

	if err := validate.Required("apn_configuration", "body", m.ApnConfiguration); err != nil {
		return err
	}

	if m.ApnConfiguration != nil {
		if err := m.ApnConfiguration.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("apn_configuration")
			}
			return err
		}
	}

	return nil
}

func (m *Apn) validateApnName(formats strfmt.Registry) error {

	if err := validate.Required("apn_name", "body", ApnName(m.ApnName)); err != nil {
		return err
	}

	if err := m.ApnName.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("apn_name")
		}
		return err
	}

	return nil
}

// ContextValidate validate this apn based on the context it is used
func (m *Apn) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateApnConfiguration(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateApnName(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Apn) contextValidateApnConfiguration(ctx context.Context, formats strfmt.Registry) error {

	if m.ApnConfiguration != nil {
		if err := m.ApnConfiguration.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("apn_configuration")
			}
			return err
		}
	}

	return nil
}

func (m *Apn) contextValidateApnName(ctx context.Context, formats strfmt.Registry) error {

	if err := m.ApnName.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("apn_name")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Apn) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Apn) UnmarshalBinary(b []byte) error {
	var res Apn
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
