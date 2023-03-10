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

// FilesForArchive files for archive
//
// swagger:model FilesForArchive
type FilesForArchive []*FilesForArchiveItems0

// Validate validates this files for archive
func (m FilesForArchive) Validate(formats strfmt.Registry) error {
	var res []error

	for i := 0; i < len(m); i++ {
		if swag.IsZero(m[i]) { // not required
			continue
		}

		if m[i] != nil {
			if err := m[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName(strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName(strconv.Itoa(i))
				}
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validate this files for archive based on the context it is used
func (m FilesForArchive) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	for i := 0; i < len(m); i++ {

		if m[i] != nil {
			if err := m[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName(strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName(strconv.Itoa(i))
				}
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// FilesForArchiveItems0 files for archive items0
//
// swagger:model FilesForArchiveItems0
type FilesForArchiveItems0 struct {

	// files
	Files []*FilesForArchiveItems0FilesItems0 `json:"files"`

	// id
	ID ArchiveID `json:"id,omitempty"`
}

// Validate validates this files for archive items0
func (m *FilesForArchiveItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateFiles(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *FilesForArchiveItems0) validateFiles(formats strfmt.Registry) error {
	if swag.IsZero(m.Files) { // not required
		return nil
	}

	for i := 0; i < len(m.Files); i++ {
		if swag.IsZero(m.Files[i]) { // not required
			continue
		}

		if m.Files[i] != nil {
			if err := m.Files[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("files" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("files" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *FilesForArchiveItems0) validateID(formats strfmt.Registry) error {
	if swag.IsZero(m.ID) { // not required
		return nil
	}

	if err := m.ID.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("id")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("id")
		}
		return err
	}

	return nil
}

// ContextValidate validate this files for archive items0 based on the context it is used
func (m *FilesForArchiveItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateFiles(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateID(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *FilesForArchiveItems0) contextValidateFiles(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Files); i++ {

		if m.Files[i] != nil {
			if err := m.Files[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("files" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("files" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *FilesForArchiveItems0) contextValidateID(ctx context.Context, formats strfmt.Registry) error {

	if err := m.ID.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("id")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("id")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *FilesForArchiveItems0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *FilesForArchiveItems0) UnmarshalBinary(b []byte) error {
	var res FilesForArchiveItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// FilesForArchiveItems0FilesItems0 files for archive items0 files items0
//
// swagger:model FilesForArchiveItems0FilesItems0
type FilesForArchiveItems0FilesItems0 struct {

	// name
	Name FileName `json:"name,omitempty"`

	// new name
	NewName FileName `json:"new-name,omitempty"`

	// new path
	NewPath FilePath `json:"new-path,omitempty"`

	// path
	Path FilePath `json:"path,omitempty"`
}

// Validate validates this files for archive items0 files items0
func (m *FilesForArchiveItems0FilesItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNewName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNewPath(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePath(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *FilesForArchiveItems0FilesItems0) validateName(formats strfmt.Registry) error {
	if swag.IsZero(m.Name) { // not required
		return nil
	}

	if err := m.Name.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("name")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("name")
		}
		return err
	}

	return nil
}

func (m *FilesForArchiveItems0FilesItems0) validateNewName(formats strfmt.Registry) error {
	if swag.IsZero(m.NewName) { // not required
		return nil
	}

	if err := m.NewName.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("new-name")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("new-name")
		}
		return err
	}

	return nil
}

func (m *FilesForArchiveItems0FilesItems0) validateNewPath(formats strfmt.Registry) error {
	if swag.IsZero(m.NewPath) { // not required
		return nil
	}

	if err := m.NewPath.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("new-path")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("new-path")
		}
		return err
	}

	return nil
}

func (m *FilesForArchiveItems0FilesItems0) validatePath(formats strfmt.Registry) error {
	if swag.IsZero(m.Path) { // not required
		return nil
	}

	if err := m.Path.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("path")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("path")
		}
		return err
	}

	return nil
}

// ContextValidate validate this files for archive items0 files items0 based on the context it is used
func (m *FilesForArchiveItems0FilesItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateName(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateNewName(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateNewPath(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidatePath(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *FilesForArchiveItems0FilesItems0) contextValidateName(ctx context.Context, formats strfmt.Registry) error {

	if err := m.Name.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("name")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("name")
		}
		return err
	}

	return nil
}

func (m *FilesForArchiveItems0FilesItems0) contextValidateNewName(ctx context.Context, formats strfmt.Registry) error {

	if err := m.NewName.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("new-name")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("new-name")
		}
		return err
	}

	return nil
}

func (m *FilesForArchiveItems0FilesItems0) contextValidateNewPath(ctx context.Context, formats strfmt.Registry) error {

	if err := m.NewPath.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("new-path")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("new-path")
		}
		return err
	}

	return nil
}

func (m *FilesForArchiveItems0FilesItems0) contextValidatePath(ctx context.Context, formats strfmt.Registry) error {

	if err := m.Path.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("path")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("path")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *FilesForArchiveItems0FilesItems0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *FilesForArchiveItems0FilesItems0) UnmarshalBinary(b []byte) error {
	var res FilesForArchiveItems0FilesItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
