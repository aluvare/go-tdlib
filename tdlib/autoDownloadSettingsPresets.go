// AUTOGENERATED - DO NOT EDIT

package tdlib

// AutoDownloadSettingsPresets Contains auto-download settings presets for the user
type AutoDownloadSettingsPresets struct {
	tdCommon
	Low    *AutoDownloadSettings `json:"low"`    // Preset with lowest settings; supposed to be used by default when roaming
	Medium *AutoDownloadSettings `json:"medium"` // Preset with medium settings; supposed to be used by default when using mobile data
	High   *AutoDownloadSettings `json:"high"`   // Preset with highest settings; supposed to be used by default when connected on Wi-Fi
}

// MessageType return the string telegram-type of AutoDownloadSettingsPresets
func (autoDownloadSettingsPresets *AutoDownloadSettingsPresets) MessageType() string {
	return "autoDownloadSettingsPresets"
}

// NewAutoDownloadSettingsPresets creates a new AutoDownloadSettingsPresets
//
// @param low Preset with lowest settings; supposed to be used by default when roaming
// @param medium Preset with medium settings; supposed to be used by default when using mobile data
// @param high Preset with highest settings; supposed to be used by default when connected on Wi-Fi
func NewAutoDownloadSettingsPresets(low *AutoDownloadSettings, medium *AutoDownloadSettings, high *AutoDownloadSettings) *AutoDownloadSettingsPresets {
	autoDownloadSettingsPresetsTemp := AutoDownloadSettingsPresets{
		tdCommon: tdCommon{Type: "autoDownloadSettingsPresets"},
		Low:      low,
		Medium:   medium,
		High:     high,
	}

	return &autoDownloadSettingsPresetsTemp
}
