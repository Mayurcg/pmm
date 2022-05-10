// Code generated by go-swagger; DO NOT EDIT.

package gmodels

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// GlobalConfig GlobalConfig defines configuration parameters that are valid globally
// unless overwritten.
//
// swagger:model GlobalConfig
type GlobalConfig struct {

	// ops genie API key file
	OpsGenieAPIKeyFile string `json:"opsgenie_api_key_file,omitempty"`

	// SMTP auth identity
	SMTPAuthIdentity string `json:"smtp_auth_identity,omitempty"`

	// SMTP auth username
	SMTPAuthUsername string `json:"smtp_auth_username,omitempty"`

	// SMTP from
	SMTPFrom string `json:"smtp_from,omitempty"`

	// SMTP hello
	SMTPHello string `json:"smtp_hello,omitempty"`

	// SMTP require TLS
	SMTPRequireTLS bool `json:"smtp_require_tls,omitempty"`

	// slack API URL file
	SlackAPIURLFile string `json:"slack_api_url_file,omitempty"`

	// we chat API corp ID
	WeChatAPICorpID string `json:"wechat_api_corp_id,omitempty"`

	// http config
	HTTPConfig *HTTPClientConfig `json:"http_config,omitempty"`

	// opsgenie api key
	OpsgenieAPIKey Secret `json:"opsgenie_api_key,omitempty"`

	// opsgenie api url
	OpsgenieAPIURL *URL `json:"opsgenie_api_url,omitempty"`

	// pagerduty url
	PagerdutyURL *URL `json:"pagerduty_url,omitempty"`

	// resolve timeout
	ResolveTimeout Duration `json:"resolve_timeout,omitempty"`

	// slack api url
	SlackAPIURL *SecretURL `json:"slack_api_url,omitempty"`

	// smtp auth password
	SMTPAuthPassword Secret `json:"smtp_auth_password,omitempty"`

	// smtp auth secret
	SMTPAuthSecret Secret `json:"smtp_auth_secret,omitempty"`

	// smtp smarthost
	SMTPSmarthost *HostPort `json:"smtp_smarthost,omitempty"`

	// victorops api key
	VictoropsAPIKey Secret `json:"victorops_api_key,omitempty"`

	// victorops api url
	VictoropsAPIURL *URL `json:"victorops_api_url,omitempty"`

	// wechat api secret
	WechatAPISecret Secret `json:"wechat_api_secret,omitempty"`

	// wechat api url
	WechatAPIURL *URL `json:"wechat_api_url,omitempty"`
}

// Validate validates this global config
func (m *GlobalConfig) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateHTTPConfig(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOpsgenieAPIKey(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOpsgenieAPIURL(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePagerdutyURL(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateResolveTimeout(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSlackAPIURL(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSMTPAuthPassword(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSMTPAuthSecret(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSMTPSmarthost(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVictoropsAPIKey(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVictoropsAPIURL(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateWechatAPISecret(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateWechatAPIURL(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GlobalConfig) validateHTTPConfig(formats strfmt.Registry) error {
	if swag.IsZero(m.HTTPConfig) { // not required
		return nil
	}

	if m.HTTPConfig != nil {
		if err := m.HTTPConfig.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("http_config")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("http_config")
			}
			return err
		}
	}

	return nil
}

func (m *GlobalConfig) validateOpsgenieAPIKey(formats strfmt.Registry) error {
	if swag.IsZero(m.OpsgenieAPIKey) { // not required
		return nil
	}

	if err := m.OpsgenieAPIKey.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("opsgenie_api_key")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("opsgenie_api_key")
		}
		return err
	}

	return nil
}

func (m *GlobalConfig) validateOpsgenieAPIURL(formats strfmt.Registry) error {
	if swag.IsZero(m.OpsgenieAPIURL) { // not required
		return nil
	}

	if m.OpsgenieAPIURL != nil {
		if err := m.OpsgenieAPIURL.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("opsgenie_api_url")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("opsgenie_api_url")
			}
			return err
		}
	}

	return nil
}

func (m *GlobalConfig) validatePagerdutyURL(formats strfmt.Registry) error {
	if swag.IsZero(m.PagerdutyURL) { // not required
		return nil
	}

	if m.PagerdutyURL != nil {
		if err := m.PagerdutyURL.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("pagerduty_url")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("pagerduty_url")
			}
			return err
		}
	}

	return nil
}

func (m *GlobalConfig) validateResolveTimeout(formats strfmt.Registry) error {
	if swag.IsZero(m.ResolveTimeout) { // not required
		return nil
	}

	if err := m.ResolveTimeout.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("resolve_timeout")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("resolve_timeout")
		}
		return err
	}

	return nil
}

func (m *GlobalConfig) validateSlackAPIURL(formats strfmt.Registry) error {
	if swag.IsZero(m.SlackAPIURL) { // not required
		return nil
	}

	if m.SlackAPIURL != nil {
		if err := m.SlackAPIURL.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("slack_api_url")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("slack_api_url")
			}
			return err
		}
	}

	return nil
}

func (m *GlobalConfig) validateSMTPAuthPassword(formats strfmt.Registry) error {
	if swag.IsZero(m.SMTPAuthPassword) { // not required
		return nil
	}

	if err := m.SMTPAuthPassword.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("smtp_auth_password")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("smtp_auth_password")
		}
		return err
	}

	return nil
}

func (m *GlobalConfig) validateSMTPAuthSecret(formats strfmt.Registry) error {
	if swag.IsZero(m.SMTPAuthSecret) { // not required
		return nil
	}

	if err := m.SMTPAuthSecret.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("smtp_auth_secret")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("smtp_auth_secret")
		}
		return err
	}

	return nil
}

func (m *GlobalConfig) validateSMTPSmarthost(formats strfmt.Registry) error {
	if swag.IsZero(m.SMTPSmarthost) { // not required
		return nil
	}

	if m.SMTPSmarthost != nil {
		if err := m.SMTPSmarthost.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("smtp_smarthost")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("smtp_smarthost")
			}
			return err
		}
	}

	return nil
}

func (m *GlobalConfig) validateVictoropsAPIKey(formats strfmt.Registry) error {
	if swag.IsZero(m.VictoropsAPIKey) { // not required
		return nil
	}

	if err := m.VictoropsAPIKey.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("victorops_api_key")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("victorops_api_key")
		}
		return err
	}

	return nil
}

func (m *GlobalConfig) validateVictoropsAPIURL(formats strfmt.Registry) error {
	if swag.IsZero(m.VictoropsAPIURL) { // not required
		return nil
	}

	if m.VictoropsAPIURL != nil {
		if err := m.VictoropsAPIURL.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("victorops_api_url")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("victorops_api_url")
			}
			return err
		}
	}

	return nil
}

func (m *GlobalConfig) validateWechatAPISecret(formats strfmt.Registry) error {
	if swag.IsZero(m.WechatAPISecret) { // not required
		return nil
	}

	if err := m.WechatAPISecret.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("wechat_api_secret")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("wechat_api_secret")
		}
		return err
	}

	return nil
}

func (m *GlobalConfig) validateWechatAPIURL(formats strfmt.Registry) error {
	if swag.IsZero(m.WechatAPIURL) { // not required
		return nil
	}

	if m.WechatAPIURL != nil {
		if err := m.WechatAPIURL.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("wechat_api_url")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("wechat_api_url")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this global config based on the context it is used
func (m *GlobalConfig) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateHTTPConfig(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateOpsgenieAPIKey(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateOpsgenieAPIURL(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidatePagerdutyURL(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateResolveTimeout(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSlackAPIURL(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSMTPAuthPassword(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSMTPAuthSecret(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSMTPSmarthost(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateVictoropsAPIKey(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateVictoropsAPIURL(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateWechatAPISecret(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateWechatAPIURL(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GlobalConfig) contextValidateHTTPConfig(ctx context.Context, formats strfmt.Registry) error {

	if m.HTTPConfig != nil {
		if err := m.HTTPConfig.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("http_config")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("http_config")
			}
			return err
		}
	}

	return nil
}

func (m *GlobalConfig) contextValidateOpsgenieAPIKey(ctx context.Context, formats strfmt.Registry) error {

	if err := m.OpsgenieAPIKey.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("opsgenie_api_key")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("opsgenie_api_key")
		}
		return err
	}

	return nil
}

func (m *GlobalConfig) contextValidateOpsgenieAPIURL(ctx context.Context, formats strfmt.Registry) error {

	if m.OpsgenieAPIURL != nil {
		if err := m.OpsgenieAPIURL.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("opsgenie_api_url")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("opsgenie_api_url")
			}
			return err
		}
	}

	return nil
}

func (m *GlobalConfig) contextValidatePagerdutyURL(ctx context.Context, formats strfmt.Registry) error {

	if m.PagerdutyURL != nil {
		if err := m.PagerdutyURL.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("pagerduty_url")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("pagerduty_url")
			}
			return err
		}
	}

	return nil
}

func (m *GlobalConfig) contextValidateResolveTimeout(ctx context.Context, formats strfmt.Registry) error {

	if err := m.ResolveTimeout.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("resolve_timeout")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("resolve_timeout")
		}
		return err
	}

	return nil
}

func (m *GlobalConfig) contextValidateSlackAPIURL(ctx context.Context, formats strfmt.Registry) error {

	if m.SlackAPIURL != nil {
		if err := m.SlackAPIURL.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("slack_api_url")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("slack_api_url")
			}
			return err
		}
	}

	return nil
}

func (m *GlobalConfig) contextValidateSMTPAuthPassword(ctx context.Context, formats strfmt.Registry) error {

	if err := m.SMTPAuthPassword.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("smtp_auth_password")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("smtp_auth_password")
		}
		return err
	}

	return nil
}

func (m *GlobalConfig) contextValidateSMTPAuthSecret(ctx context.Context, formats strfmt.Registry) error {

	if err := m.SMTPAuthSecret.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("smtp_auth_secret")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("smtp_auth_secret")
		}
		return err
	}

	return nil
}

func (m *GlobalConfig) contextValidateSMTPSmarthost(ctx context.Context, formats strfmt.Registry) error {

	if m.SMTPSmarthost != nil {
		if err := m.SMTPSmarthost.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("smtp_smarthost")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("smtp_smarthost")
			}
			return err
		}
	}

	return nil
}

func (m *GlobalConfig) contextValidateVictoropsAPIKey(ctx context.Context, formats strfmt.Registry) error {

	if err := m.VictoropsAPIKey.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("victorops_api_key")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("victorops_api_key")
		}
		return err
	}

	return nil
}

func (m *GlobalConfig) contextValidateVictoropsAPIURL(ctx context.Context, formats strfmt.Registry) error {

	if m.VictoropsAPIURL != nil {
		if err := m.VictoropsAPIURL.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("victorops_api_url")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("victorops_api_url")
			}
			return err
		}
	}

	return nil
}

func (m *GlobalConfig) contextValidateWechatAPISecret(ctx context.Context, formats strfmt.Registry) error {

	if err := m.WechatAPISecret.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("wechat_api_secret")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("wechat_api_secret")
		}
		return err
	}

	return nil
}

func (m *GlobalConfig) contextValidateWechatAPIURL(ctx context.Context, formats strfmt.Registry) error {

	if m.WechatAPIURL != nil {
		if err := m.WechatAPIURL.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("wechat_api_url")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("wechat_api_url")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *GlobalConfig) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GlobalConfig) UnmarshalBinary(b []byte) error {
	var res GlobalConfig
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
