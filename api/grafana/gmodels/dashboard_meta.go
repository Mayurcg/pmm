// Code generated by go-swagger; DO NOT EDIT.

package gmodels

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// DashboardMeta dashboard meta
//
// swagger:model DashboardMeta
type DashboardMeta struct {

	// can admin
	CanAdmin bool `json:"canAdmin,omitempty"`

	// can delete
	CanDelete bool `json:"canDelete,omitempty"`

	// can edit
	CanEdit bool `json:"canEdit,omitempty"`

	// can save
	CanSave bool `json:"canSave,omitempty"`

	// can star
	CanStar bool `json:"canStar,omitempty"`

	// created
	// Format: date-time
	Created strfmt.DateTime `json:"created,omitempty"`

	// created by
	CreatedBy string `json:"createdBy,omitempty"`

	// expires
	// Format: date-time
	Expires strfmt.DateTime `json:"expires,omitempty"`

	// folder Id
	FolderID int64 `json:"folderId,omitempty"`

	// folder title
	FolderTitle string `json:"folderTitle,omitempty"`

	// folder Uid
	FolderUID string `json:"folderUid,omitempty"`

	// folder Url
	FolderURL string `json:"folderUrl,omitempty"`

	// has Acl
	HasACL bool `json:"hasAcl,omitempty"`

	// is folder
	IsFolder bool `json:"isFolder,omitempty"`

	// is home
	IsHome bool `json:"isHome,omitempty"`

	// is snapshot
	IsSnapshot bool `json:"isSnapshot,omitempty"`

	// is starred
	IsStarred bool `json:"isStarred,omitempty"`

	// provisioned
	Provisioned bool `json:"provisioned,omitempty"`

	// provisioned external Id
	ProvisionedExternalID string `json:"provisionedExternalId,omitempty"`

	// slug
	Slug string `json:"slug,omitempty"`

	// type
	Type string `json:"type,omitempty"`

	// updated
	// Format: date-time
	Updated strfmt.DateTime `json:"updated,omitempty"`

	// updated by
	UpdatedBy string `json:"updatedBy,omitempty"`

	// Url
	URL string `json:"url,omitempty"`

	// version
	Version int64 `json:"version,omitempty"`

	// annotations permissions
	AnnotationsPermissions *AnnotationPermission `json:"annotationsPermissions,omitempty"`
}

// Validate validates this dashboard meta
func (m *DashboardMeta) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCreated(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateExpires(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUpdated(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAnnotationsPermissions(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DashboardMeta) validateCreated(formats strfmt.Registry) error {
	if swag.IsZero(m.Created) { // not required
		return nil
	}

	if err := validate.FormatOf("created", "body", "date-time", m.Created.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *DashboardMeta) validateExpires(formats strfmt.Registry) error {
	if swag.IsZero(m.Expires) { // not required
		return nil
	}

	if err := validate.FormatOf("expires", "body", "date-time", m.Expires.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *DashboardMeta) validateUpdated(formats strfmt.Registry) error {
	if swag.IsZero(m.Updated) { // not required
		return nil
	}

	if err := validate.FormatOf("updated", "body", "date-time", m.Updated.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *DashboardMeta) validateAnnotationsPermissions(formats strfmt.Registry) error {
	if swag.IsZero(m.AnnotationsPermissions) { // not required
		return nil
	}

	if m.AnnotationsPermissions != nil {
		if err := m.AnnotationsPermissions.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("annotationsPermissions")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("annotationsPermissions")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this dashboard meta based on the context it is used
func (m *DashboardMeta) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAnnotationsPermissions(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DashboardMeta) contextValidateAnnotationsPermissions(ctx context.Context, formats strfmt.Registry) error {

	if m.AnnotationsPermissions != nil {
		if err := m.AnnotationsPermissions.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("annotationsPermissions")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("annotationsPermissions")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *DashboardMeta) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DashboardMeta) UnmarshalBinary(b []byte) error {
	var res DashboardMeta
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
