package model

import (
	"github.com/go-openapi/errors"
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// Event event
// swagger:model event
type Event struct {

	// contenttype
	Contenttype Contenttype `json:"contenttype,omitempty"`

	// data
	Data interface{} `json:"data,omitempty"`

	// extensions
	Extensions Extensions `json:"extensions,omitempty"`

	// id
	// Required: true
	ID ID `json:"id"`

	// source
	// Required: true
	Source Source `json:"source"`

	// specversion
	// Required: true
	Specversion Specversion `json:"specversion"`

	// time
	// Format: date-time
	Time Time `json:"time,omitempty"`

	// type
	// Required: true
	Type Type `json:"type"`

	// triggeredID
	// Required: true
	TriggeredID string `json:"triggeredid"`

	// shkeptncontext
	// Required: true
	ShkeptnContext string `json:"shkeptncontext"`

	// GitCommitID
	// Required: false
	GitcommitID string `json:"gitcommitid"`
}

// Validate validates this event
func (m *Event) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateContenttype(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSource(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSpecversion(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Event) validateContenttype(formats strfmt.Registry) error {

	if swag.IsZero(m.Contenttype) { // not required
		return nil
	}

	if err := m.Contenttype.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("contenttype")
		}
		return err
	}

	return nil
}

func (m *Event) validateID(formats strfmt.Registry) error {

	if err := m.ID.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("id")
		}
		return err
	}

	return nil
}

func (m *Event) validateSource(formats strfmt.Registry) error {

	if err := m.Source.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("source")
		}
		return err
	}

	return nil
}

func (m *Event) validateSpecversion(formats strfmt.Registry) error {

	if err := m.Specversion.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("specversion")
		}
		return err
	}

	return nil
}

func (m *Event) validateTime(formats strfmt.Registry) error {

	if swag.IsZero(m.Time) { // not required
		return nil
	}

	if err := m.Time.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("time")
		}
		return err
	}

	return nil
}

func (m *Event) validateType(formats strfmt.Registry) error {

	if err := m.Type.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("type")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Event) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Event) UnmarshalBinary(b []byte) error {
	var res Event
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
