// AUTOGENERATED - DO NOT EDIT

package tdlib

import (
	"encoding/json"
	"fmt"
)

// InputCredentials Contains information about the payment method chosen by the user
type InputCredentials interface {
	GetInputCredentialsEnum() InputCredentialsEnum
}

// InputCredentialsEnum Alias for abstract InputCredentials 'Sub-Classes', used as constant-enum here
type InputCredentialsEnum string

// InputCredentials enums
const (
	InputCredentialsSavedType      InputCredentialsEnum = "inputCredentialsSaved"
	InputCredentialsNewType        InputCredentialsEnum = "inputCredentialsNew"
	InputCredentialsAndroidPayType InputCredentialsEnum = "inputCredentialsAndroidPay"
	InputCredentialsApplePayType   InputCredentialsEnum = "inputCredentialsApplePay"
)

func unmarshalInputCredentials(rawMsg *json.RawMessage) (InputCredentials, error) {

	if rawMsg == nil {
		return nil, nil
	}
	var objMap map[string]interface{}
	err := json.Unmarshal(*rawMsg, &objMap)
	if err != nil {
		return nil, err
	}

	switch InputCredentialsEnum(objMap["@type"].(string)) {
	case InputCredentialsSavedType:
		var inputCredentialsSaved InputCredentialsSaved
		err := json.Unmarshal(*rawMsg, &inputCredentialsSaved)
		return &inputCredentialsSaved, err

	case InputCredentialsNewType:
		var inputCredentialsNew InputCredentialsNew
		err := json.Unmarshal(*rawMsg, &inputCredentialsNew)
		return &inputCredentialsNew, err

	case InputCredentialsAndroidPayType:
		var inputCredentialsAndroidPay InputCredentialsAndroidPay
		err := json.Unmarshal(*rawMsg, &inputCredentialsAndroidPay)
		return &inputCredentialsAndroidPay, err

	case InputCredentialsApplePayType:
		var inputCredentialsApplePay InputCredentialsApplePay
		err := json.Unmarshal(*rawMsg, &inputCredentialsApplePay)
		return &inputCredentialsApplePay, err

	default:
		return nil, fmt.Errorf("Error UnMarshaling, unknown type:" + objMap["@type"].(string))
	}
}

// InputCredentialsSaved Applies if a user chooses some previously saved payment credentials. To use their previously saved credentials, the user must have a valid temporary password
type InputCredentialsSaved struct {
	tdCommon
	SavedCredentialsID string `json:"saved_credentials_id"` // Identifier of the saved credentials
}

// MessageType return the string telegram-type of InputCredentialsSaved
func (inputCredentialsSaved *InputCredentialsSaved) MessageType() string {
	return "inputCredentialsSaved"
}

// NewInputCredentialsSaved creates a new InputCredentialsSaved
//
// @param savedCredentialsID Identifier of the saved credentials
func NewInputCredentialsSaved(savedCredentialsID string) *InputCredentialsSaved {
	inputCredentialsSavedTemp := InputCredentialsSaved{
		tdCommon:           tdCommon{Type: "inputCredentialsSaved"},
		SavedCredentialsID: savedCredentialsID,
	}

	return &inputCredentialsSavedTemp
}

// GetInputCredentialsEnum return the enum type of this object
func (inputCredentialsSaved *InputCredentialsSaved) GetInputCredentialsEnum() InputCredentialsEnum {
	return InputCredentialsSavedType
}

// InputCredentialsNew Applies if a user enters new credentials on a payment provider website
type InputCredentialsNew struct {
	tdCommon
	Data      string `json:"data"`       // Contains JSON-encoded data with a credential identifier from the payment provider
	AllowSave bool   `json:"allow_save"` // True, if the credential identifier can be saved on the server side
}

// MessageType return the string telegram-type of InputCredentialsNew
func (inputCredentialsNew *InputCredentialsNew) MessageType() string {
	return "inputCredentialsNew"
}

// NewInputCredentialsNew creates a new InputCredentialsNew
//
// @param data Contains JSON-encoded data with a credential identifier from the payment provider
// @param allowSave True, if the credential identifier can be saved on the server side
func NewInputCredentialsNew(data string, allowSave bool) *InputCredentialsNew {
	inputCredentialsNewTemp := InputCredentialsNew{
		tdCommon:  tdCommon{Type: "inputCredentialsNew"},
		Data:      data,
		AllowSave: allowSave,
	}

	return &inputCredentialsNewTemp
}

// GetInputCredentialsEnum return the enum type of this object
func (inputCredentialsNew *InputCredentialsNew) GetInputCredentialsEnum() InputCredentialsEnum {
	return InputCredentialsNewType
}

// InputCredentialsAndroidPay Applies if a user enters new credentials using Android Pay
type InputCredentialsAndroidPay struct {
	tdCommon
	Data string `json:"data"` // JSON-encoded data with the credential identifier
}

// MessageType return the string telegram-type of InputCredentialsAndroidPay
func (inputCredentialsAndroidPay *InputCredentialsAndroidPay) MessageType() string {
	return "inputCredentialsAndroidPay"
}

// NewInputCredentialsAndroidPay creates a new InputCredentialsAndroidPay
//
// @param data JSON-encoded data with the credential identifier
func NewInputCredentialsAndroidPay(data string) *InputCredentialsAndroidPay {
	inputCredentialsAndroidPayTemp := InputCredentialsAndroidPay{
		tdCommon: tdCommon{Type: "inputCredentialsAndroidPay"},
		Data:     data,
	}

	return &inputCredentialsAndroidPayTemp
}

// GetInputCredentialsEnum return the enum type of this object
func (inputCredentialsAndroidPay *InputCredentialsAndroidPay) GetInputCredentialsEnum() InputCredentialsEnum {
	return InputCredentialsAndroidPayType
}

// InputCredentialsApplePay Applies if a user enters new credentials using Apple Pay
type InputCredentialsApplePay struct {
	tdCommon
	Data string `json:"data"` // JSON-encoded data with the credential identifier
}

// MessageType return the string telegram-type of InputCredentialsApplePay
func (inputCredentialsApplePay *InputCredentialsApplePay) MessageType() string {
	return "inputCredentialsApplePay"
}

// NewInputCredentialsApplePay creates a new InputCredentialsApplePay
//
// @param data JSON-encoded data with the credential identifier
func NewInputCredentialsApplePay(data string) *InputCredentialsApplePay {
	inputCredentialsApplePayTemp := InputCredentialsApplePay{
		tdCommon: tdCommon{Type: "inputCredentialsApplePay"},
		Data:     data,
	}

	return &inputCredentialsApplePayTemp
}

// GetInputCredentialsEnum return the enum type of this object
func (inputCredentialsApplePay *InputCredentialsApplePay) GetInputCredentialsEnum() InputCredentialsEnum {
	return InputCredentialsApplePayType
}
