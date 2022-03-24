// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// Cbsd cbsd
//
// swagger:model cbsd
type Cbsd struct {

	// capabilities
	// Required: true
	Capabilities Capabilities `json:"capabilities"`

	// id of cbsd in SAS
	// Example: some_cbsd_id
	CbsdID string `json:"cbsd_id,omitempty"`

	// fcc id
	// Example: some_fcc_id
	// Required: true
	// Min Length: 1
	FccID string `json:"fcc_id"`

	// frequency preferences
	// Required: true
	FrequencyPreferences FrequencyPreferences `json:"frequency_preferences"`

	// grant
	Grant *Grant `json:"grant,omitempty"`

	// database id of cbsd
	// Required: true
	ID int64 `json:"id"`

	// false if cbsd have not contacted DP for certain amount of time
	// Required: true
	IsActive bool `json:"is_active"`

	// serial number
	// Example: some_serial_number
	// Required: true
	// Min Length: 1
	SerialNumber string `json:"serial_number"`

	// state of cbsd in SAS
	// Required: true
	// Enum: [unregistered registered]
	State string `json:"state"`

	// user id
	// Example: some_user_id
	// Required: true
	// Min Length: 1
	UserID string `json:"user_id"`
}

// Validate validates this cbsd
func (m *Cbsd) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCapabilities(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFccID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFrequencyPreferences(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateGrant(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIsActive(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSerialNumber(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateState(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUserID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Cbsd) validateCapabilities(formats strfmt.Registry) error {

	if err := m.Capabilities.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("capabilities")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("capabilities")
		}
		return err
	}

	return nil
}

func (m *Cbsd) validateFccID(formats strfmt.Registry) error {

	if err := validate.RequiredString("fcc_id", "body", m.FccID); err != nil {
		return err
	}

	if err := validate.MinLength("fcc_id", "body", m.FccID, 1); err != nil {
		return err
	}

	return nil
}

func (m *Cbsd) validateFrequencyPreferences(formats strfmt.Registry) error {

	if err := m.FrequencyPreferences.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("frequency_preferences")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("frequency_preferences")
		}
		return err
	}

	return nil
}

func (m *Cbsd) validateGrant(formats strfmt.Registry) error {
	if swag.IsZero(m.Grant) { // not required
		return nil
	}

	if m.Grant != nil {
		if err := m.Grant.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("grant")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("grant")
			}
			return err
		}
	}

	return nil
}

func (m *Cbsd) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", int64(m.ID)); err != nil {
		return err
	}

	return nil
}

func (m *Cbsd) validateIsActive(formats strfmt.Registry) error {

	if err := validate.Required("is_active", "body", bool(m.IsActive)); err != nil {
		return err
	}

	return nil
}

func (m *Cbsd) validateSerialNumber(formats strfmt.Registry) error {

	if err := validate.RequiredString("serial_number", "body", m.SerialNumber); err != nil {
		return err
	}

	if err := validate.MinLength("serial_number", "body", m.SerialNumber, 1); err != nil {
		return err
	}

	return nil
}

var cbsdTypeStatePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["unregistered","registered"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		cbsdTypeStatePropEnum = append(cbsdTypeStatePropEnum, v)
	}
}

const (

	// CbsdStateUnregistered captures enum value "unregistered"
	CbsdStateUnregistered string = "unregistered"

	// CbsdStateRegistered captures enum value "registered"
	CbsdStateRegistered string = "registered"
)

// prop value enum
func (m *Cbsd) validateStateEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, cbsdTypeStatePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *Cbsd) validateState(formats strfmt.Registry) error {

	if err := validate.RequiredString("state", "body", m.State); err != nil {
		return err
	}

	// value enum
	if err := m.validateStateEnum("state", "body", m.State); err != nil {
		return err
	}

	return nil
}

func (m *Cbsd) validateUserID(formats strfmt.Registry) error {

	if err := validate.RequiredString("user_id", "body", m.UserID); err != nil {
		return err
	}

	if err := validate.MinLength("user_id", "body", m.UserID, 1); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this cbsd based on the context it is used
func (m *Cbsd) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateCapabilities(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateFrequencyPreferences(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateGrant(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Cbsd) contextValidateCapabilities(ctx context.Context, formats strfmt.Registry) error {

	if err := m.Capabilities.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("capabilities")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("capabilities")
		}
		return err
	}

	return nil
}

func (m *Cbsd) contextValidateFrequencyPreferences(ctx context.Context, formats strfmt.Registry) error {

	if err := m.FrequencyPreferences.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("frequency_preferences")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("frequency_preferences")
		}
		return err
	}

	return nil
}

func (m *Cbsd) contextValidateGrant(ctx context.Context, formats strfmt.Registry) error {

	if m.Grant != nil {
		if err := m.Grant.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("grant")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("grant")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Cbsd) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Cbsd) UnmarshalBinary(b []byte) error {
	var res Cbsd
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
