/*
 * Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
 * This product includes software developed at Datadog (https://www.datadoghq.com/).
 * Copyright 2019-Present Datadog, Inc.
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package datadog

import (
	"encoding/json"
	"fmt"
)

// SyntheticsBasicAuthWeb Object to handle basic authentication when performing the test.
type SyntheticsBasicAuthWeb struct {
	// Password to use for the basic authentication.
	Password string                     `json:"password"`
	Type     SyntheticsBasicAuthWebType `json:"type"`
	// Username to use for the basic authentication.
	Username string `json:"username"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:-`
}

// NewSyntheticsBasicAuthWeb instantiates a new SyntheticsBasicAuthWeb object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSyntheticsBasicAuthWeb(password string, type_ SyntheticsBasicAuthWebType, username string) *SyntheticsBasicAuthWeb {
	this := SyntheticsBasicAuthWeb{}
	this.Password = password
	this.Type = type_
	this.Username = username
	return &this
}

// NewSyntheticsBasicAuthWebWithDefaults instantiates a new SyntheticsBasicAuthWeb object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSyntheticsBasicAuthWebWithDefaults() *SyntheticsBasicAuthWeb {
	this := SyntheticsBasicAuthWeb{}
	var type_ SyntheticsBasicAuthWebType = SYNTHETICSBASICAUTHWEBTYPE_WEB
	this.Type = type_
	return &this
}

// GetPassword returns the Password field value
func (o *SyntheticsBasicAuthWeb) GetPassword() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Password
}

// GetPasswordOk returns a tuple with the Password field value
// and a boolean to check if the value has been set.
func (o *SyntheticsBasicAuthWeb) GetPasswordOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Password, true
}

// SetPassword sets field value
func (o *SyntheticsBasicAuthWeb) SetPassword(v string) {
	o.Password = v
}

// GetType returns the Type field value
func (o *SyntheticsBasicAuthWeb) GetType() SyntheticsBasicAuthWebType {
	if o == nil {
		var ret SyntheticsBasicAuthWebType
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *SyntheticsBasicAuthWeb) GetTypeOk() (*SyntheticsBasicAuthWebType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *SyntheticsBasicAuthWeb) SetType(v SyntheticsBasicAuthWebType) {
	o.Type = v
}

// GetUsername returns the Username field value
func (o *SyntheticsBasicAuthWeb) GetUsername() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Username
}

// GetUsernameOk returns a tuple with the Username field value
// and a boolean to check if the value has been set.
func (o *SyntheticsBasicAuthWeb) GetUsernameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Username, true
}

// SetUsername sets field value
func (o *SyntheticsBasicAuthWeb) SetUsername(v string) {
	o.Username = v
}

func (o SyntheticsBasicAuthWeb) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return json.Marshal(o.UnparsedObject)
	}
	if true {
		toSerialize["password"] = o.Password
	}
	if true {
		toSerialize["type"] = o.Type
	}
	if true {
		toSerialize["username"] = o.Username
	}
	return json.Marshal(toSerialize)
}

func (o *SyntheticsBasicAuthWeb) UnmarshalJSON(bytes []byte) (err error) {
	raw := map[string]interface{}{}
	required := struct {
		Password *string                     `json:"password"`
		Type     *SyntheticsBasicAuthWebType `json:"type"`
		Username *string                     `json:"username"`
	}{}
	all := struct {
		Password string                     `json:"password"`
		Type     SyntheticsBasicAuthWebType `json:"type"`
		Username string                     `json:"username"`
	}{}
	err = json.Unmarshal(bytes, &required)
	if err != nil {
		return err
	}
	if required.Password == nil {
		return fmt.Errorf("Required field password missing")
	}
	if required.Type == nil {
		return fmt.Errorf("Required field type missing")
	}
	if required.Username == nil {
		return fmt.Errorf("Required field username missing")
	}
	err = json.Unmarshal(bytes, &all)
	if err != nil {
		err = json.Unmarshal(bytes, &raw)
		if err != nil {
			return err
		}
		o.UnparsedObject = raw
		return nil
	}
	if v := all.Type; !v.IsValid() {
		err = json.Unmarshal(bytes, &raw)
		if err != nil {
			return err
		}
		o.UnparsedObject = raw
		return nil
	}
	o.Password = all.Password
	o.Type = all.Type
	o.Username = all.Username
	return nil
}

type NullableSyntheticsBasicAuthWeb struct {
	value *SyntheticsBasicAuthWeb
	isSet bool
}

func (v NullableSyntheticsBasicAuthWeb) Get() *SyntheticsBasicAuthWeb {
	return v.value
}

func (v *NullableSyntheticsBasicAuthWeb) Set(val *SyntheticsBasicAuthWeb) {
	v.value = val
	v.isSet = true
}

func (v NullableSyntheticsBasicAuthWeb) IsSet() bool {
	return v.isSet
}

func (v *NullableSyntheticsBasicAuthWeb) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSyntheticsBasicAuthWeb(val *SyntheticsBasicAuthWeb) *NullableSyntheticsBasicAuthWeb {
	return &NullableSyntheticsBasicAuthWeb{value: val, isSet: true}
}

func (v NullableSyntheticsBasicAuthWeb) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSyntheticsBasicAuthWeb) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
