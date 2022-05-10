// Code generated by go-swagger; DO NOT EDIT.

package gmodels

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// PostableAPIReceiver postable Api receiver
//
// swagger:model PostableApiReceiver
type PostableAPIReceiver struct {

	// email configs
	EmailConfigs []*EmailConfig `json:"email_configs"`

	// grafana managed receivers
	GrafanaManagedReceivers []*PostableGrafanaReceiver `json:"grafana_managed_receiver_configs"`

	// A unique identifier for this receiver.
	Name string `json:"name,omitempty"`

	// ops genie configs
	OpsGenieConfigs []*OpsGenieConfig `json:"opsgenie_configs"`

	// pagerduty configs
	PagerdutyConfigs []*PagerdutyConfig `json:"pagerduty_configs"`

	// pushover configs
	PushoverConfigs []*PushoverConfig `json:"pushover_configs"`

	// s n s configs
	SNSConfigs []*SNSConfig `json:"sns_configs"`

	// slack configs
	SlackConfigs []*SlackConfig `json:"slack_configs"`

	// victor ops configs
	VictorOpsConfigs []*VictorOpsConfig `json:"victorops_configs"`

	// webhook configs
	WebhookConfigs []*WebhookConfig `json:"webhook_configs"`

	// wechat configs
	WechatConfigs []*WechatConfig `json:"wechat_configs"`
}

// Validate validates this postable Api receiver
func (m *PostableAPIReceiver) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEmailConfigs(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateGrafanaManagedReceivers(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOpsGenieConfigs(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePagerdutyConfigs(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePushoverConfigs(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSNSConfigs(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSlackConfigs(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVictorOpsConfigs(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateWebhookConfigs(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateWechatConfigs(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PostableAPIReceiver) validateEmailConfigs(formats strfmt.Registry) error {
	if swag.IsZero(m.EmailConfigs) { // not required
		return nil
	}

	for i := 0; i < len(m.EmailConfigs); i++ {
		if swag.IsZero(m.EmailConfigs[i]) { // not required
			continue
		}

		if m.EmailConfigs[i] != nil {
			if err := m.EmailConfigs[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("email_configs" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("email_configs" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *PostableAPIReceiver) validateGrafanaManagedReceivers(formats strfmt.Registry) error {
	if swag.IsZero(m.GrafanaManagedReceivers) { // not required
		return nil
	}

	for i := 0; i < len(m.GrafanaManagedReceivers); i++ {
		if swag.IsZero(m.GrafanaManagedReceivers[i]) { // not required
			continue
		}

		if m.GrafanaManagedReceivers[i] != nil {
			if err := m.GrafanaManagedReceivers[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("grafana_managed_receiver_configs" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("grafana_managed_receiver_configs" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *PostableAPIReceiver) validateOpsGenieConfigs(formats strfmt.Registry) error {
	if swag.IsZero(m.OpsGenieConfigs) { // not required
		return nil
	}

	for i := 0; i < len(m.OpsGenieConfigs); i++ {
		if swag.IsZero(m.OpsGenieConfigs[i]) { // not required
			continue
		}

		if m.OpsGenieConfigs[i] != nil {
			if err := m.OpsGenieConfigs[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("opsgenie_configs" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("opsgenie_configs" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *PostableAPIReceiver) validatePagerdutyConfigs(formats strfmt.Registry) error {
	if swag.IsZero(m.PagerdutyConfigs) { // not required
		return nil
	}

	for i := 0; i < len(m.PagerdutyConfigs); i++ {
		if swag.IsZero(m.PagerdutyConfigs[i]) { // not required
			continue
		}

		if m.PagerdutyConfigs[i] != nil {
			if err := m.PagerdutyConfigs[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("pagerduty_configs" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("pagerduty_configs" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *PostableAPIReceiver) validatePushoverConfigs(formats strfmt.Registry) error {
	if swag.IsZero(m.PushoverConfigs) { // not required
		return nil
	}

	for i := 0; i < len(m.PushoverConfigs); i++ {
		if swag.IsZero(m.PushoverConfigs[i]) { // not required
			continue
		}

		if m.PushoverConfigs[i] != nil {
			if err := m.PushoverConfigs[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("pushover_configs" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("pushover_configs" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *PostableAPIReceiver) validateSNSConfigs(formats strfmt.Registry) error {
	if swag.IsZero(m.SNSConfigs) { // not required
		return nil
	}

	for i := 0; i < len(m.SNSConfigs); i++ {
		if swag.IsZero(m.SNSConfigs[i]) { // not required
			continue
		}

		if m.SNSConfigs[i] != nil {
			if err := m.SNSConfigs[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("sns_configs" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("sns_configs" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *PostableAPIReceiver) validateSlackConfigs(formats strfmt.Registry) error {
	if swag.IsZero(m.SlackConfigs) { // not required
		return nil
	}

	for i := 0; i < len(m.SlackConfigs); i++ {
		if swag.IsZero(m.SlackConfigs[i]) { // not required
			continue
		}

		if m.SlackConfigs[i] != nil {
			if err := m.SlackConfigs[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("slack_configs" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("slack_configs" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *PostableAPIReceiver) validateVictorOpsConfigs(formats strfmt.Registry) error {
	if swag.IsZero(m.VictorOpsConfigs) { // not required
		return nil
	}

	for i := 0; i < len(m.VictorOpsConfigs); i++ {
		if swag.IsZero(m.VictorOpsConfigs[i]) { // not required
			continue
		}

		if m.VictorOpsConfigs[i] != nil {
			if err := m.VictorOpsConfigs[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("victorops_configs" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("victorops_configs" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *PostableAPIReceiver) validateWebhookConfigs(formats strfmt.Registry) error {
	if swag.IsZero(m.WebhookConfigs) { // not required
		return nil
	}

	for i := 0; i < len(m.WebhookConfigs); i++ {
		if swag.IsZero(m.WebhookConfigs[i]) { // not required
			continue
		}

		if m.WebhookConfigs[i] != nil {
			if err := m.WebhookConfigs[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("webhook_configs" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("webhook_configs" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *PostableAPIReceiver) validateWechatConfigs(formats strfmt.Registry) error {
	if swag.IsZero(m.WechatConfigs) { // not required
		return nil
	}

	for i := 0; i < len(m.WechatConfigs); i++ {
		if swag.IsZero(m.WechatConfigs[i]) { // not required
			continue
		}

		if m.WechatConfigs[i] != nil {
			if err := m.WechatConfigs[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("wechat_configs" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("wechat_configs" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this postable Api receiver based on the context it is used
func (m *PostableAPIReceiver) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateEmailConfigs(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateGrafanaManagedReceivers(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateOpsGenieConfigs(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidatePagerdutyConfigs(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidatePushoverConfigs(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSNSConfigs(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSlackConfigs(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateVictorOpsConfigs(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateWebhookConfigs(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateWechatConfigs(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PostableAPIReceiver) contextValidateEmailConfigs(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.EmailConfigs); i++ {

		if m.EmailConfigs[i] != nil {
			if err := m.EmailConfigs[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("email_configs" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("email_configs" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *PostableAPIReceiver) contextValidateGrafanaManagedReceivers(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.GrafanaManagedReceivers); i++ {

		if m.GrafanaManagedReceivers[i] != nil {
			if err := m.GrafanaManagedReceivers[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("grafana_managed_receiver_configs" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("grafana_managed_receiver_configs" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *PostableAPIReceiver) contextValidateOpsGenieConfigs(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.OpsGenieConfigs); i++ {

		if m.OpsGenieConfigs[i] != nil {
			if err := m.OpsGenieConfigs[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("opsgenie_configs" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("opsgenie_configs" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *PostableAPIReceiver) contextValidatePagerdutyConfigs(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.PagerdutyConfigs); i++ {

		if m.PagerdutyConfigs[i] != nil {
			if err := m.PagerdutyConfigs[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("pagerduty_configs" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("pagerduty_configs" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *PostableAPIReceiver) contextValidatePushoverConfigs(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.PushoverConfigs); i++ {

		if m.PushoverConfigs[i] != nil {
			if err := m.PushoverConfigs[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("pushover_configs" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("pushover_configs" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *PostableAPIReceiver) contextValidateSNSConfigs(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.SNSConfigs); i++ {

		if m.SNSConfigs[i] != nil {
			if err := m.SNSConfigs[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("sns_configs" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("sns_configs" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *PostableAPIReceiver) contextValidateSlackConfigs(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.SlackConfigs); i++ {

		if m.SlackConfigs[i] != nil {
			if err := m.SlackConfigs[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("slack_configs" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("slack_configs" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *PostableAPIReceiver) contextValidateVictorOpsConfigs(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.VictorOpsConfigs); i++ {

		if m.VictorOpsConfigs[i] != nil {
			if err := m.VictorOpsConfigs[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("victorops_configs" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("victorops_configs" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *PostableAPIReceiver) contextValidateWebhookConfigs(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.WebhookConfigs); i++ {

		if m.WebhookConfigs[i] != nil {
			if err := m.WebhookConfigs[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("webhook_configs" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("webhook_configs" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *PostableAPIReceiver) contextValidateWechatConfigs(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.WechatConfigs); i++ {

		if m.WechatConfigs[i] != nil {
			if err := m.WechatConfigs[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("wechat_configs" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("wechat_configs" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *PostableAPIReceiver) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PostableAPIReceiver) UnmarshalBinary(b []byte) error {
	var res PostableAPIReceiver
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
