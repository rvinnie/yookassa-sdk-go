// Package yoosettings describes all the necessary entities for working with YooMoney Settings.
package yoosettings

type SettingsStatus string

const (
	// Enabled – store is signed up for YooMoney and it can accept payments.
	Enabled SettingsStatus = "enabled"

	// Disabled – store can't accept payments (not signed up yet, closed, or temporarily unavailable).
	Disabled SettingsStatus = "disabled"
)
