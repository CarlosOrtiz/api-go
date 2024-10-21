package config

type BasicResponse struct {
	Success bool `json:"success"`
	Detail  any  `json:"detail"`
	Message any  `json:"message,omitempty"`
}
