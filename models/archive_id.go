// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
)

// ArchiveID archive uid
//
// swagger:model ArchiveID
type ArchiveID int64

// Validate validates this archive ID
func (m ArchiveID) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this archive ID based on context it is used
func (m ArchiveID) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}