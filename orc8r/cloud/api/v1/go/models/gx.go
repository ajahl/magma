// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// Gx gx configuration
//
// swagger:model gx
type Gx struct {

	// disable gx
	// Example: false
	DisableGx *bool `json:"disableGx,omitempty"`

	// overwrite apn
	OverwriteAPN string `json:"overwrite_apn,omitempty"`

	// server
	Server *DiameterClientConfigs `json:"server,omitempty"`

	// servers
	Servers []*DiameterClientConfigs `json:"servers"`

	// virtual apn rules
	VirtualAPNRules []*VirtualAPNRule `json:"virtual_apn_rules"`
}

// Validate validates this gx
func (m *Gx) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateServer(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateServers(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVirtualAPNRules(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Gx) validateServer(formats strfmt.Registry) error {
	if swag.IsZero(m.Server) { // not required
		return nil
	}

	if m.Server != nil {
		if err := m.Server.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("server")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("server")
			}
			return err
		}
	}

	return nil
}

func (m *Gx) validateServers(formats strfmt.Registry) error {
	if swag.IsZero(m.Servers) { // not required
		return nil
	}

	for i := 0; i < len(m.Servers); i++ {
		if swag.IsZero(m.Servers[i]) { // not required
			continue
		}

		if m.Servers[i] != nil {
			if err := m.Servers[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("servers" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("servers" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Gx) validateVirtualAPNRules(formats strfmt.Registry) error {
	if swag.IsZero(m.VirtualAPNRules) { // not required
		return nil
	}

	for i := 0; i < len(m.VirtualAPNRules); i++ {
		if swag.IsZero(m.VirtualAPNRules[i]) { // not required
			continue
		}

		if m.VirtualAPNRules[i] != nil {
			if err := m.VirtualAPNRules[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("virtual_apn_rules" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("virtual_apn_rules" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this gx based on the context it is used
func (m *Gx) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateServer(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateServers(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateVirtualAPNRules(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Gx) contextValidateServer(ctx context.Context, formats strfmt.Registry) error {

	if m.Server != nil {
		if err := m.Server.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("server")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("server")
			}
			return err
		}
	}

	return nil
}

func (m *Gx) contextValidateServers(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Servers); i++ {

		if m.Servers[i] != nil {
			if err := m.Servers[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("servers" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("servers" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Gx) contextValidateVirtualAPNRules(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.VirtualAPNRules); i++ {

		if m.VirtualAPNRules[i] != nil {
			if err := m.VirtualAPNRules[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("virtual_apn_rules" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("virtual_apn_rules" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *Gx) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Gx) UnmarshalBinary(b []byte) error {
	var res Gx
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
