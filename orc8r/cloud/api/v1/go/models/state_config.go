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

// StateConfig State configuration
//
// swagger:model state_config
type StateConfig struct {

	// Sync interval in seconds
	// Example: 60
	// Minimum: 3
	SyncInterval uint32 `json:"sync_interval,omitempty"`
}

// Validate validates this state config
func (m *StateConfig) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSyncInterval(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *StateConfig) validateSyncInterval(formats strfmt.Registry) error {
	if swag.IsZero(m.SyncInterval) { // not required
		return nil
	}

	if err := validate.MinimumUint("sync_interval", "body", uint64(m.SyncInterval), 3, false); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this state config based on context it is used
func (m *StateConfig) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *StateConfig) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *StateConfig) UnmarshalBinary(b []byte) error {
	var res StateConfig
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
