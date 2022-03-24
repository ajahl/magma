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

// Capabilities capabilities
//
// swagger:model capabilities
type Capabilities struct {

	// antenna gain
	// Example: 15
	// Required: true
	AntennaGain *float64 `json:"antenna_gain"`

	// max tx power available on cbsd
	// Example: 30
	// Required: true
	MaxPower *float64 `json:"max_power"`

	// min tx power available on cbsd
	// Required: true
	MinPower *float64 `json:"min_power"`

	// number of antennas
	// Example: 2
	// Required: true
	// Minimum: 1
	NumberOfAntennas int64 `json:"number_of_antennas"`
}

// Validate validates this capabilities
func (m *Capabilities) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAntennaGain(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMaxPower(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMinPower(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNumberOfAntennas(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Capabilities) validateAntennaGain(formats strfmt.Registry) error {

	if err := validate.Required("antenna_gain", "body", m.AntennaGain); err != nil {
		return err
	}

	return nil
}

func (m *Capabilities) validateMaxPower(formats strfmt.Registry) error {

	if err := validate.Required("max_power", "body", m.MaxPower); err != nil {
		return err
	}

	return nil
}

func (m *Capabilities) validateMinPower(formats strfmt.Registry) error {

	if err := validate.Required("min_power", "body", m.MinPower); err != nil {
		return err
	}

	return nil
}

func (m *Capabilities) validateNumberOfAntennas(formats strfmt.Registry) error {

	if err := validate.Required("number_of_antennas", "body", int64(m.NumberOfAntennas)); err != nil {
		return err
	}

	if err := validate.MinimumInt("number_of_antennas", "body", m.NumberOfAntennas, 1, false); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this capabilities based on context it is used
func (m *Capabilities) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *Capabilities) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Capabilities) UnmarshalBinary(b []byte) error {
	var res Capabilities
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
