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
	"github.com/go-openapi/validate"
)

// PaginatedCbsds Page of cbsds
//
// swagger:model paginated_cbsds
type PaginatedCbsds struct {

	// cbsds
	// Required: true
	// Read Only: true
	Cbsds []*Cbsd `json:"cbsds"`

	// Total number of cbsds
	// Example: 10
	// Required: true
	// Read Only: true
	TotalCount int64 `json:"total_count"`
}

// Validate validates this paginated cbsds
func (m *PaginatedCbsds) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCbsds(formats); err != nil {
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

func (m *PaginatedCbsds) validateCbsds(formats strfmt.Registry) error {

	if err := validate.Required("cbsds", "body", m.Cbsds); err != nil {
		return err
	}

	for i := 0; i < len(m.Cbsds); i++ {
		if swag.IsZero(m.Cbsds[i]) { // not required
			continue
		}

		if m.Cbsds[i] != nil {
			if err := m.Cbsds[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("cbsds" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *PaginatedCbsds) validateTotalCount(formats strfmt.Registry) error {

	if err := validate.Required("total_count", "body", int64(m.TotalCount)); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this paginated cbsds based on the context it is used
func (m *PaginatedCbsds) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateCbsds(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateTotalCount(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PaginatedCbsds) contextValidateCbsds(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "cbsds", "body", []*Cbsd(m.Cbsds)); err != nil {
		return err
	}

	for i := 0; i < len(m.Cbsds); i++ {

		if m.Cbsds[i] != nil {
			if err := m.Cbsds[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("cbsds" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *PaginatedCbsds) contextValidateTotalCount(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "total_count", "body", int64(m.TotalCount)); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *PaginatedCbsds) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PaginatedCbsds) UnmarshalBinary(b []byte) error {
	var res PaginatedCbsds
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
