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

// GettableAlertSilencer gettable alert silencer
//
// swagger:model gettable_alert_silencer
type GettableAlertSilencer struct {
	AlertSilencer

	// id
	// Required: true
	ID *string `json:"id"`

	// status
	// Required: true
	Status *AlertSilenceStatus `json:"status"`

	// updated at
	// Required: true
	UpdatedAt *string `json:"updatedAt"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *GettableAlertSilencer) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 AlertSilencer
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.AlertSilencer = aO0

	// AO1
	var dataAO1 struct {
		ID *string `json:"id"`

		Status *AlertSilenceStatus `json:"status"`

		UpdatedAt *string `json:"updatedAt"`
	}
	if err := swag.ReadJSON(raw, &dataAO1); err != nil {
		return err
	}

	m.ID = dataAO1.ID

	m.Status = dataAO1.Status

	m.UpdatedAt = dataAO1.UpdatedAt

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m GettableAlertSilencer) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	aO0, err := swag.WriteJSON(m.AlertSilencer)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)
	var dataAO1 struct {
		ID *string `json:"id"`

		Status *AlertSilenceStatus `json:"status"`

		UpdatedAt *string `json:"updatedAt"`
	}

	dataAO1.ID = m.ID

	dataAO1.Status = m.Status

	dataAO1.UpdatedAt = m.UpdatedAt

	jsonDataAO1, errAO1 := swag.WriteJSON(dataAO1)
	if errAO1 != nil {
		return nil, errAO1
	}
	_parts = append(_parts, jsonDataAO1)
	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this gettable alert silencer
func (m *GettableAlertSilencer) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with AlertSilencer
	if err := m.AlertSilencer.Validate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUpdatedAt(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GettableAlertSilencer) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

func (m *GettableAlertSilencer) validateStatus(formats strfmt.Registry) error {

	if err := validate.Required("status", "body", m.Status); err != nil {
		return err
	}

	if m.Status != nil {
		if err := m.Status.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("status")
			}
			return err
		}
	}

	return nil
}

func (m *GettableAlertSilencer) validateUpdatedAt(formats strfmt.Registry) error {

	if err := validate.Required("updatedAt", "body", m.UpdatedAt); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this gettable alert silencer based on the context it is used
func (m *GettableAlertSilencer) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with AlertSilencer
	if err := m.AlertSilencer.ContextValidate(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateStatus(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GettableAlertSilencer) contextValidateStatus(ctx context.Context, formats strfmt.Registry) error {

	if m.Status != nil {
		if err := m.Status.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("status")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *GettableAlertSilencer) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GettableAlertSilencer) UnmarshalBinary(b []byte) error {
	var res GettableAlertSilencer
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
