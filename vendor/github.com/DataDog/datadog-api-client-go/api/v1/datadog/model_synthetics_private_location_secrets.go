/*
 * Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
 * This product includes software developed at Datadog (https://www.datadoghq.com/).
 * Copyright 2019-Present Datadog, Inc.
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package datadog

import (
	"encoding/json"
)

// SyntheticsPrivateLocationSecrets Secrets for the private location. Only present in the response when creating the private location.
type SyntheticsPrivateLocationSecrets struct {
	Authentication   *SyntheticsPrivateLocationSecretsAuthentication   `json:"authentication,omitempty"`
	ConfigDecryption *SyntheticsPrivateLocationSecretsConfigDecryption `json:"config_decryption,omitempty"`
}

// NewSyntheticsPrivateLocationSecrets instantiates a new SyntheticsPrivateLocationSecrets object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSyntheticsPrivateLocationSecrets() *SyntheticsPrivateLocationSecrets {
	this := SyntheticsPrivateLocationSecrets{}
	return &this
}

// NewSyntheticsPrivateLocationSecretsWithDefaults instantiates a new SyntheticsPrivateLocationSecrets object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSyntheticsPrivateLocationSecretsWithDefaults() *SyntheticsPrivateLocationSecrets {
	this := SyntheticsPrivateLocationSecrets{}
	return &this
}

// GetAuthentication returns the Authentication field value if set, zero value otherwise.
func (o *SyntheticsPrivateLocationSecrets) GetAuthentication() SyntheticsPrivateLocationSecretsAuthentication {
	if o == nil || o.Authentication == nil {
		var ret SyntheticsPrivateLocationSecretsAuthentication
		return ret
	}
	return *o.Authentication
}

// GetAuthenticationOk returns a tuple with the Authentication field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsPrivateLocationSecrets) GetAuthenticationOk() (*SyntheticsPrivateLocationSecretsAuthentication, bool) {
	if o == nil || o.Authentication == nil {
		return nil, false
	}
	return o.Authentication, true
}

// HasAuthentication returns a boolean if a field has been set.
func (o *SyntheticsPrivateLocationSecrets) HasAuthentication() bool {
	if o != nil && o.Authentication != nil {
		return true
	}

	return false
}

// SetAuthentication gets a reference to the given SyntheticsPrivateLocationSecretsAuthentication and assigns it to the Authentication field.
func (o *SyntheticsPrivateLocationSecrets) SetAuthentication(v SyntheticsPrivateLocationSecretsAuthentication) {
	o.Authentication = &v
}

// GetConfigDecryption returns the ConfigDecryption field value if set, zero value otherwise.
func (o *SyntheticsPrivateLocationSecrets) GetConfigDecryption() SyntheticsPrivateLocationSecretsConfigDecryption {
	if o == nil || o.ConfigDecryption == nil {
		var ret SyntheticsPrivateLocationSecretsConfigDecryption
		return ret
	}
	return *o.ConfigDecryption
}

// GetConfigDecryptionOk returns a tuple with the ConfigDecryption field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsPrivateLocationSecrets) GetConfigDecryptionOk() (*SyntheticsPrivateLocationSecretsConfigDecryption, bool) {
	if o == nil || o.ConfigDecryption == nil {
		return nil, false
	}
	return o.ConfigDecryption, true
}

// HasConfigDecryption returns a boolean if a field has been set.
func (o *SyntheticsPrivateLocationSecrets) HasConfigDecryption() bool {
	if o != nil && o.ConfigDecryption != nil {
		return true
	}

	return false
}

// SetConfigDecryption gets a reference to the given SyntheticsPrivateLocationSecretsConfigDecryption and assigns it to the ConfigDecryption field.
func (o *SyntheticsPrivateLocationSecrets) SetConfigDecryption(v SyntheticsPrivateLocationSecretsConfigDecryption) {
	o.ConfigDecryption = &v
}

func (o SyntheticsPrivateLocationSecrets) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Authentication != nil {
		toSerialize["authentication"] = o.Authentication
	}
	if o.ConfigDecryption != nil {
		toSerialize["config_decryption"] = o.ConfigDecryption
	}
	return json.Marshal(toSerialize)
}

type NullableSyntheticsPrivateLocationSecrets struct {
	value *SyntheticsPrivateLocationSecrets
	isSet bool
}

func (v NullableSyntheticsPrivateLocationSecrets) Get() *SyntheticsPrivateLocationSecrets {
	return v.value
}

func (v *NullableSyntheticsPrivateLocationSecrets) Set(val *SyntheticsPrivateLocationSecrets) {
	v.value = val
	v.isSet = true
}

func (v NullableSyntheticsPrivateLocationSecrets) IsSet() bool {
	return v.isSet
}

func (v *NullableSyntheticsPrivateLocationSecrets) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSyntheticsPrivateLocationSecrets(val *SyntheticsPrivateLocationSecrets) *NullableSyntheticsPrivateLocationSecrets {
	return &NullableSyntheticsPrivateLocationSecrets{value: val, isSet: true}
}

func (v NullableSyntheticsPrivateLocationSecrets) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSyntheticsPrivateLocationSecrets) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
