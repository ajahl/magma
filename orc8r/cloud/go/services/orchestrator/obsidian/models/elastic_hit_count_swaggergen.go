// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
)

// ElasticHitCount elastic hit count
//
// swagger:model elastic_hit_count
type ElasticHitCount float64

// Validate validates this elastic hit count
func (m ElasticHitCount) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this elastic hit count based on context it is used
func (m ElasticHitCount) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
