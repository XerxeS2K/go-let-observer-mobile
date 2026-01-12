//go:build android || mobile
// +build android mobile

package main

type FocusField int

const (
	FocusHost FocusField = iota
	FocusPort
	FocusPassword
)

type RCONSetupUI struct {
	Host     string
	Port     string
	Password string
	Focus    FocusField
}

func NewRCONSetupUI() *RCONSetupUI {
	return &RCONSetupUI{
		Host:  "",
		Port:  "8080",
		Focus: FocusHost,
	}
}

func (ui *RCONSetupUI) ToConfig() RCONConfig {
	return RCONConfig{
		Host:     ui.Host,
		Port:     ui.Port,
		Password: ui.Password,
	}
}
