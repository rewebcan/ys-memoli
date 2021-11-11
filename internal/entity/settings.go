package entity

import (
	validation "github.com/go-ozzo/ozzo-validation"
)

// StoreSettingRequest represents the request to store the settings.
type StoreSettingRequest struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

func (ssr *StoreSettingRequest) Validate() error {
	return validation.ValidateStruct(ssr,
		validation.Field(&ssr.Key, validation.Required),
		validation.Field(&ssr.Value, validation.Required),
	)
}

type SettingsResponse struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}
