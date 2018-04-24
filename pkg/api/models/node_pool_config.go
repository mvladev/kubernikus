// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// NodePoolConfig node pool config
// swagger:model NodePoolConfig
type NodePoolConfig struct {

	// repair
	Repair bool `json:"repair,omitempty"`

	// upgrade
	Upgrade bool `json:"upgrade,omitempty"`
}

// Validate validates this node pool config
func (m *NodePoolConfig) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *NodePoolConfig) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NodePoolConfig) UnmarshalBinary(b []byte) error {
	var res NodePoolConfig
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}