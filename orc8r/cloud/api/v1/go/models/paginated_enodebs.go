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

// PaginatedEnodebs Page of eNodeBs
//
// swagger:model paginated_enodebs
type PaginatedEnodebs struct {

	// enodebs
	// Required: true
	Enodebs map[string]*ENODEB `json:"enodebs"`

	// page token
	// Required: true
	PageToken *PageToken `json:"page_token"`

	// Estimated total number of eNodeBs
	// Example: 10
	// Required: true
	TotalCount uint64 `json:"total_count"`
}

// Validate validates this paginated enodebs
func (m *PaginatedEnodebs) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEnodebs(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePageToken(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTotalCount(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PaginatedEnodebs) validateEnodebs(formats strfmt.Registry) error {

	if err := validate.Required("enodebs", "body", m.Enodebs); err != nil {
		return err
	}

	for k := range m.Enodebs {

		if err := validate.Required("enodebs"+"."+k, "body", m.Enodebs[k]); err != nil {
			return err
		}
		if val, ok := m.Enodebs[k]; ok {
			if err := val.Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("enodebs" + "." + k)
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("enodebs" + "." + k)
				}
				return err
			}
		}

	}

	return nil
}

func (m *PaginatedEnodebs) validatePageToken(formats strfmt.Registry) error {

	if err := validate.Required("page_token", "body", m.PageToken); err != nil {
		return err
	}

	if err := validate.Required("page_token", "body", m.PageToken); err != nil {
		return err
	}

	if m.PageToken != nil {
		if err := m.PageToken.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("page_token")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("page_token")
			}
			return err
		}
	}

	return nil
}

func (m *PaginatedEnodebs) validateTotalCount(formats strfmt.Registry) error {

	if err := validate.Required("total_count", "body", uint64(m.TotalCount)); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this paginated enodebs based on the context it is used
func (m *PaginatedEnodebs) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateEnodebs(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidatePageToken(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PaginatedEnodebs) contextValidateEnodebs(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.Required("enodebs", "body", m.Enodebs); err != nil {
		return err
	}

	for k := range m.Enodebs {

		if val, ok := m.Enodebs[k]; ok {
			if err := val.ContextValidate(ctx, formats); err != nil {
				return err
			}
		}

	}

	return nil
}

func (m *PaginatedEnodebs) contextValidatePageToken(ctx context.Context, formats strfmt.Registry) error {

	if m.PageToken != nil {
		if err := m.PageToken.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("page_token")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("page_token")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *PaginatedEnodebs) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PaginatedEnodebs) UnmarshalBinary(b []byte) error {
	var res PaginatedEnodebs
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
