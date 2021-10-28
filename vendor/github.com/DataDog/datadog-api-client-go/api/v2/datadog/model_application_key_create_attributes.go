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

// ApplicationKeyCreateAttributes Attributes used to create an application Key.
type ApplicationKeyCreateAttributes struct {
	// Name of the application key.
	Name string `json:"name"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:-`
}

// NewApplicationKeyCreateAttributes instantiates a new ApplicationKeyCreateAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApplicationKeyCreateAttributes(name string) *ApplicationKeyCreateAttributes {
	this := ApplicationKeyCreateAttributes{}
	this.Name = name
	return &this
}

// NewApplicationKeyCreateAttributesWithDefaults instantiates a new ApplicationKeyCreateAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApplicationKeyCreateAttributesWithDefaults() *ApplicationKeyCreateAttributes {
	this := ApplicationKeyCreateAttributes{}
	return &this
}

// GetName returns the Name field value
func (o *ApplicationKeyCreateAttributes) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ApplicationKeyCreateAttributes) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ApplicationKeyCreateAttributes) SetName(v string) {
	o.Name = v
}

func (o ApplicationKeyCreateAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return json.Marshal(o.UnparsedObject)
	}
	if true {
		toSerialize["name"] = o.Name
	}
	return json.Marshal(toSerialize)
}

func (o *ApplicationKeyCreateAttributes) UnmarshalJSON(bytes []byte) (err error) {
	raw := map[string]interface{}{}
	required := struct {
		Name *string `json:"name"`
	}{}
	all := struct {
		Name string `json:"name"`
	}{}
	err = json.Unmarshal(bytes, &required)
	if err != nil {
		return err
	}
	if required.Name == nil {
		return fmt.Errorf("Required field name missing")
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
	o.Name = all.Name
	return nil
}

type NullableApplicationKeyCreateAttributes struct {
	value *ApplicationKeyCreateAttributes
	isSet bool
}

func (v NullableApplicationKeyCreateAttributes) Get() *ApplicationKeyCreateAttributes {
	return v.value
}

func (v *NullableApplicationKeyCreateAttributes) Set(val *ApplicationKeyCreateAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullableApplicationKeyCreateAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullableApplicationKeyCreateAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApplicationKeyCreateAttributes(val *ApplicationKeyCreateAttributes) *NullableApplicationKeyCreateAttributes {
	return &NullableApplicationKeyCreateAttributes{value: val, isSet: true}
}

func (v NullableApplicationKeyCreateAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApplicationKeyCreateAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
