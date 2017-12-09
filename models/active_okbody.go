// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// ActiveOKBody active o k body
// swagger:model activeOKBody
type ActiveOKBody struct {
	BasicReport

	ModelReport

	DeltaReport

	SeriesDateInt
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *ActiveOKBody) UnmarshalJSON(raw []byte) error {

	var aO0 BasicReport
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.BasicReport = aO0

	var aO1 ModelReport
	if err := swag.ReadJSON(raw, &aO1); err != nil {
		return err
	}
	m.ModelReport = aO1

	var aO2 DeltaReport
	if err := swag.ReadJSON(raw, &aO2); err != nil {
		return err
	}
	m.DeltaReport = aO2

	var aO3 SeriesDateInt
	if err := swag.ReadJSON(raw, &aO3); err != nil {
		return err
	}
	m.SeriesDateInt = aO3

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m ActiveOKBody) MarshalJSON() ([]byte, error) {
	var _parts [][]byte

	aO0, err := swag.WriteJSON(m.BasicReport)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)

	aO1, err := swag.WriteJSON(m.ModelReport)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO1)

	aO2, err := swag.WriteJSON(m.DeltaReport)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO2)

	aO3, err := swag.WriteJSON(m.SeriesDateInt)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO3)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this active o k body
func (m *ActiveOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.BasicReport.Validate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.ModelReport.Validate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.DeltaReport.Validate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.SeriesDateInt.Validate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *ActiveOKBody) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ActiveOKBody) UnmarshalBinary(b []byte) error {
	var res ActiveOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
